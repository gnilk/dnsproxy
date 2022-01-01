package main

import (
	"errors"
	"log"
	"strings"
)

type Resolver struct {
	conf *Config
	sys  *System
}

var ErrHostNotFound = errors.New("Host not found")

func NewResolver(system *System) *Resolver {
	re := Resolver{
		conf: system.config,
		sys:  system,
	}
	log.Printf("Resolving the following:\n")
	for _, r := range system.config.Resolve {
		log.Printf("%s -> %s\n", r.Name, r.IpV4)
	}
	return &re
}

func (r *Resolver) Resolve(domain string) (string, error) {
	for _, r := range r.conf.Resolve {
		log.Printf("testing: %s", r.Name)
		if WildcardPatternMatch(strings.ToLower(domain), strings.ToLower(r.Name)) {
			return r.IpV4, nil
		}
	}

	log.Printf("Resolve with device cache")
	domain = strings.TrimSuffix(domain, ".")
	ip, err := r.sys.DeviceCache().NameToIP(domain)
	if err != nil {
		log.Printf("Name %s not found\n", domain)
		return "", ErrHostNotFound
	}

	return ip.String(), nil
}
