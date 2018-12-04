package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type System struct {
	performanceLog LogClient
	config         *Config
	rulesEngine    *RulesEngine
	routerClient   RouterClient
	deviceCache    *DeviceCache
	mutex          sync.Mutex
}

//
// NewSystem creates a new system object and initializes the sub systems
//
func NewSystem(cfgFileName string) *System {

	sys := System{}
	cfg, err := sys.loadConfig(cfgFileName)
	if err != nil {
		log.Panic(err)
	}
	sys.config = cfg
	err = sys.validateConfig(cfg)
	if err != nil {
		log.Panic(err)
	}

	re, err := NewRulesEngine(sys.config)
	if err != nil {
		log.Printf("[ERROR] failed create rules engine: ", err.Error())
		os.Exit(1)
	}

	logger, err := NewLogFileClient(sys.config.Logfile)
	if err != nil {
		log.Panic(err.Error())
	}
	sys.performanceLog = logger

	// Setup router and device cache - this will download local lan device names from the router
	// Add's support for 'names' instead of IP for the proxy host rules
	if sys.config.Router.Engine != RouterTypeNone {
		log.Printf("[INFO] Router configuration found - trying...")
		err = sys.initializeRouter(sys.config.Router)
		if err != nil {
			log.Printf("[ERROR] Router initialization failed: %s\n", err.Error())
			log.Printf("[WARN] Device Name lookup disabled - requires working router connection\n")
		} else {
			dc := NewDeviceCache(sys.routerClient)
			err = dc.Initialize()
			if err != nil {
				log.Printf("[ERROR] Device Cache initialization failed: %s\n", err.Error())
			} else {
				log.Printf("[INFO] Ok, device list downloaded")
				sys.deviceCache = dc
				dc.Dump()
				if sys.Config().Router.PollChanges {
					log.Printf("[INFO] Starting router auto refresh, interval: %d sec", sys.Config().Router.PollInterval)
					dc.StartAutoRefresh(sys.Config().Router.PollInterval)
				}
			}
		}
	}

	// Attach device cache to the rules engine
	re.SetDeviceCache(sys.deviceCache)
	sys.rulesEngine = re

	return &sys
}

func (sys *System) ReloadConfig() error {

	cfg, err := sys.loadConfig("config.json")
	if err != nil {
		log.Printf("[Error] Unable to PARSE configuration\n")
		return err
	}

	// Take full mutex!!
	sys.mutex.Lock()
	defer sys.mutex.Unlock()

	err = sys.validateConfig(cfg)
	if err != nil {
		log.Printf("[Error] Invalid configuration\n")
		return err
	}

	re, err := NewRulesEngine(cfg)
	if err != nil {
		log.Printf("[ERROR] failed create rules engine: ", err.Error())
		return err
	}

	// Create new log file if it changed, otherwise keep old
	if cfg.Logfile != sys.config.Logfile {
		err = sys.performanceLog.Close()
		if err != nil {
			log.Printf("[ERROR] Unable to close old log file: ", err.Error())
			return err
		}

		logger, err := NewLogFileClient(cfg.Logfile)
		if err != nil {
			log.Panic(err.Error())
		}
		// Set it..
		sys.performanceLog = logger

	}

	re.SetDeviceCache(sys.deviceCache)
	sys.rulesEngine = re
	sys.config = cfg

	return nil
}

//
// Config return internal config object
//
func (sys *System) Config() *Config {
	sys.mutex.Lock()
	defer sys.mutex.Unlock()
	return sys.config
}

//
// RulesEngine return internal RulesEngine object
//
func (sys *System) RulesEngine() *RulesEngine {
	sys.mutex.Lock()
	defer sys.mutex.Unlock()
	return sys.rulesEngine
}

//
// RouterClient returns the internal/global RouterClient object
//
func (sys *System) RouterClient() RouterClient {
	sys.mutex.Lock()
	defer sys.mutex.Unlock()
	return sys.routerClient
}

//
// DeviceCache returns the internal/global DeviceCache object
//
func (sys *System) DeviceCache() *DeviceCache {
	sys.mutex.Lock()
	defer sys.mutex.Unlock()
	return sys.deviceCache
}

func (sys *System) PerfLog() LogClient {
	sys.mutex.Lock()
	defer sys.mutex.Unlock()
	return sys.performanceLog
}

//
// Support functions to get all subsystems up and running
//
func (sys *System) initializeRouter(router Router) error {

	if router.Engine != RouterTypeNetGear {
		return fmt.Errorf("Unknown router type '%s', check configuration", router.Engine.String())
	}

	routerClient := NewNetGearRouterClient()
	err := routerClient.Login(router.Host, router.Port, router.User, router.Password)
	if err != nil {
		return err
	}
	sys.routerClient = routerClient

	return nil
}

func (sys *System) validateConfig(config *Config) error {
	// No listening address - set a default
	if config.ListenAddress == "" {
		config.ListenAddress = ":53"
	}

	// No forwarding name server - set a default (should perhaps resolve system ns here)
	if len(config.NameServers) == 0 {
		nsdefault := NameServer{
			IP: "8.8.8.8:53",
		}
		config.NameServers = append(config.NameServers, nsdefault)
	}

	if config.Router.PollChanges == true {
		interval := config.Router.PollInterval
		if interval < 10 {
			log.Printf("[WARN] Router poll interval too low (%d) resetting to min (10 sec)\n", interval)
			config.Router.PollInterval = 10
		}
	}
	return nil
}

func (sys *System) loadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var conf = new(Config)
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil

}
