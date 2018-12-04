package main

/*
	Rules engine for the DNS proxy
*/
import (
	"errors"
	"log"
	"strings"
	"time"
)

type RulesEngine struct {
	conf        *Config
	debug       bool
	deviceCache *DeviceCache
}

func NewRulesEngine(conf *Config) (*RulesEngine, error) {
	re := RulesEngine{
		conf:        conf,
		debug:       true,
		deviceCache: nil,
	}
	return &re, nil
}

//
// SetDeviceCache assigns the cache object for device name to/from ip lookup
//
func (r *RulesEngine) SetDeviceCache(dc *DeviceCache) {
	r.deviceCache = dc
}

//
// Debugging set to true to enable debug logging specifically from this component
//
func (r *RulesEngine) Debugging(enable bool) {
	log.Printf("[INFO] RulesEngine Debugging: %v", enable)
	r.debug = enable
}

//
// Evaluate takes a domain and a host and evaluates if any rules apply
//
// For the DNS proxy the "domain" is the question and the host is the originator
//
func (r *RulesEngine) Evaluate(domain, host string) (ActionType, error) {
	action := r.conf.DefaultRule

	if r.debug {
		log.Printf("Evaluate, domain: %s, host: %s", domain, host)
	}

	// First check 'host' portion of the config - this is generally used for white/black-listing
	// of specific hosts but can also be used for generic rules like disabling DNS within a certain
	// time-span
	for _, h := range r.conf.Hosts {
		if r.HostMatch(host, h.Name) {
			if r.debug {
				log.Printf("Hostmatch, %s - %s, evaluting rules...\n", host, h.Name)
			}
			action, err := r.EvaluateRules(h.Rules)
			if err != nil {
				if r.debug {
					log.Printf("[WARN] Rules evaluation error: %s\n", err.Error())
				}
				return r.conf.OnErrorRule, err
			}
			if r.debug {
				log.Printf("  Result: %s\n", action.String())
			}

			// If we got blocked - let's leave - no need to evaluate further
			if action != ActionTypeNone {
				return action, nil
			}
		}
	}

	// Now let's check domain configuration
	// This is used to block specific domains either fully or within a certain time block
	for _, d := range r.conf.Domains {
		if r.DomainMatch(domain, d.Name) {
			if r.debug {
				log.Printf("Domainmatch, %s - %s, evaluating rules...\n", domain, d.Name)
			}
			for _, h := range d.Hosts {
				if r.HostMatch(host, h.Name) {
					if r.debug {
						log.Printf("Hostmatch, %s - %s, evaluating rules...\n", host, h.Name)
					}
					// Evaluate rules!!!!
					return r.EvaluateRules(h.Rules)
				}
			}
		}
	}

	return action, nil
}

func (r *RulesEngine) DomainMatch(domain, pattern string) bool {
	return WildcardPatternMatch(domain, pattern)
}

func (r *RulesEngine) HostMatch(host, pattern string) bool {

	// host is IP and pattern comes from config
	// check if they are equal - we have a match
	if host == pattern {
		return true
	}

	// Now match against device cache
	if r.deviceCache != nil {
		// Translate host IP to device name
		name, err := r.deviceCache.IPToName(host)
		if err == nil {
			// No error - let's match with it
			return WildcardPatternMatch(name, pattern)
		}
	}
	return WildcardPatternMatch(host, pattern)
}

func (re *RulesEngine) EvaluateRules(rules []Rule) (ActionType, error) {
	for _, r := range rules {
		action, err := r.EvaluateRule()
		if err != nil {
			if re.debug {
				log.Printf("[ERROR] EvaluateRules, failed with error: %s\n", err.Error())
			}
			return re.conf.OnErrorRule, err
		}

		if re.debug {
			log.Printf("EvaluateRules: %s -> %s (result)\n", r.Type.String(), action.String())
		}

		// None means is used to signal the 'false' from evaluation
		// Any other evaluation means a rule was it - and we break further evaluation
		if action != ActionTypeNone {
			return action, nil
		}
	}
	return ActionTypeNone, nil
}

func (r *Rule) EvaluateRule() (ActionType, error) {
	switch r.Type {
	case ActionTypeBlockedDevice:
		return ActionTypeBlockedDevice, nil
	case ActionTypeBlockedSiteBan:
		return ActionTypeBlockedSiteBan, nil
	case ActionTypeBlockedTimeSpan:
		return r.EvaluateTimeSpanBlock(r.TimeSpan)
	case ActionTypePass:
		return ActionTypePass, nil
	case ActionTypeNone:
		return ActionTypeNone, nil
	}
	log.Printf("[Warninig] Rule::EvaluateRule, invalid rule type: %v\n", r)
	return ActionTypeNone, nil
}

// EvaluateTimeSpanBlock evaluates the current time with respect to definition
func (r *Rule) EvaluateTimeSpanBlock(strTimeSpan string) (ActionType, error) {

	// TODO: Make it possible to store this preparsed data -> need possibility in modelconfig (would be good anyway)
	spans := strings.Split(strTimeSpan, "-")
	if len(spans) != 2 {
		return ActionTypeNone, errors.New("Timespan definition error, use: \"HH:mm-HH:mm\"")
	}

	tStart, err := time.Parse("15:04", spans[0])
	if err != nil {
		return ActionTypeNone, err
	}
	tEnd, err := time.Parse("15:04", spans[1])
	if err != nil {
		return ActionTypeNone, err
	}

	tNow, _ := time.Parse("15:04", time.Now().Format("15:04"))

	log.Printf("EvaluateTimeSpanBlock: Now: %s, Start: %s, End: %s\n", tNow.Format("15:04"), tStart.Format("15:04"), tEnd.Format("15:04"))

	if tNow.After(tStart) && tEnd.After(tNow) {
		return r.Type, nil
	}

	//	log.Printf("EvaluateTimeSpanBlock, %s -> pass\n", strTimeSpan)
	return ActionTypeNone, nil
}
