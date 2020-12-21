package main

import (
	"errors"
	"log"
	"strings"
)

type Resolver struct {
	conf *Config
}

var ErrHostNotFound = errors.New("Host not found")

func NewResolver(conf *Config) (*Resolver) {
	re := Resolver {
		conf: conf,
	}
	log.Printf("Resolving the following:\n")
	for _, r := range conf.Resolve {
		log.Printf("%s -> %s\n", r.Name, r.IpV4)
	}
	return &re
}

func (r *Resolver) Resolve(domain string) (string, error) {
	for _,r := range r.conf.Resolve {
		if WildcardPatternMatch(strings.ToLower(domain), strings.ToLower(r.Name)) {
			return r.IpV4, nil
		}
	}
	return "", ErrHostNotFound
}