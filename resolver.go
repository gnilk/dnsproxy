package main

import (
	"errors"
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
	return &re
}

func (r *Resolver) Resolve(domain string) (string, error) {
	for _,r := range r.conf.Resolve {
		if WildcardPatternMatch(domain, strings.ToUpper(r.Name)) {
			return r.IpV4, nil
		}
	}
	return "", ErrHostNotFound
}