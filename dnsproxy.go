package main

/*
	"Parental" DNS Proxy with IP/Domain block rules
	Will forward DNS requests transparently to the configured DNS service

	Howto
	1) Install this on a server-type machine (24/7) on your network (Raspberry PI is fine)
	2) Configure the DHCP of your router to supply the server as your DNS
	3) Configure rules for your home devices

	What you can do:
	1) White/Black list devices either with respect to DNS access
	2) Block per domain
	3) Block a device or domain within specific time ranges

	TODO:
	- Fetch DNS settings from router (doesn't require you to key them in)
	- Alias for IP names when Router manufacture is not supported
	- Implement a better way to instansiate router clients and implement more router clients...
	- Ability to serve a 'redirect' response to some kind of info site...

	Advanced usage:
	- Install InfluxDB, NodeRED and Grafana
	- Configure the 'tail' command in NodeRED to read the performance log and push to Influx
	- Setup a nice dashboard showing most frequently used sites per device/hour-of-day/etc..
*/

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/miekg/dns"
)

//
// Just a wrapper for global variables used a bit throughout the system
//
var sys *System

func main() {
	// Suck in the system configuration
	// NOTE: This will panic and fail if basics are wrong
	sys = NewSystem("config.json")

	// Start proxy
	log.Printf("[INFO] Starting proxy at: %s\n", sys.Config().ListenAddress)

	go func() {
		//srv := &dns.Server{Addr: ":53", Net: "udp", Handler: dns.HandlerFunc(dnsUdpHandler)}
		srv := &dns.Server{Addr: sys.Config().ListenAddress, Net: "udp", Handler: dns.HandlerFunc(dnsUdpHandler)}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("[ERROR] Failed to set udp listener %s\n", err.Error())
		}
	}()
	go func() {
		srv := &dns.Server{Addr: sys.Config().ListenAddress, Net: "tcp", Handler: dns.HandlerFunc(dnsTcpHandler)}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("[ERROR] Failed to set tcp listener %s\n", err.Error())
		}
	}()
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	for {
		s := <-sig

		switch s {
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGTERM:
			log.Println("sigterm, terminating...")
			os.Exit(1)
		case syscall.SIGHUP:
			log.Println("sighup, reloading configuration")
		}
	}
}

//
// This is the core of the proxy
// Takes any DNS query push it through the rules engine if result is PASS the exchange is done
// and we serve the response back to the client
//
// If we block we serve an "error"
//

const (
	notIPQuery = 0
	_IP4Query  = 4
	_IP6Query  = 6
)

func isIPQuery(q dns.Question) int {
	if q.Qclass != dns.ClassINET {
		return notIPQuery
	}

	switch q.Qtype {
	case dns.TypeA:
		return _IP4Query
	case dns.TypeAAAA:
		return _IP6Query
	default:
		return notIPQuery
	}
}

func writeFailure(w dns.ResponseWriter, message *dns.Msg) {
	m := new(dns.Msg)
	m.SetRcode(message, dns.RcodeServerFailure)
	w.WriteMsg(m)
}

func writeBlockedRoute(w dns.ResponseWriter, message *dns.Msg, IPQuery int) {
	// Allow this to be configured
	nullroute := net.ParseIP(sys.Config().IPv4BlockResolve)
	nullroutev6 := net.ParseIP(sys.Config().IPv6BlockResolve)

	q := message.Question[0]

	m := new(dns.Msg)
	m.SetReply(message)

	switch IPQuery {
	case _IP4Query:
		rrHeader := dns.RR_Header{
			Name:   q.Name,
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    10,
		}
		a := &dns.A{Hdr: rrHeader, A: nullroute}
		m.Answer = append(m.Answer, a)
	case _IP6Query:
		rrHeader := dns.RR_Header{
			Name:   q.Name,
			Rrtype: dns.TypeAAAA,
			Class:  dns.ClassINET,
			Ttl:    10,
		}
		a := &dns.AAAA{Hdr: rrHeader, AAAA: nullroutev6}
		m.Answer = append(m.Answer, a)
	}

	w.WriteMsg(m)
}

func isBlockingAction(action ActionType) bool {
	if (action == ActionTypeBlockedDevice) ||
		(action == ActionTypeBlockedSiteBan) ||
		(action == ActionTypeBlockedTimeSpan) {
		return true
	}
	return false
}

func doDnsExchange(w dns.ResponseWriter, m *dns.Msg, proto string) {

	m.Question[0].Name = strings.ToUpper(m.Question[0].Name)

	// No need to do this everytime
	c := new(dns.Client)
	c.Net = proto
	// TODO: Dig this out from Nameserver array
	r, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		fmt.Printf("[ERROR] Resolving '%s' while doing c.Exchange: %s\n", m.Question[0].Name, err.Error())
		return
	}

	// {
	// 	for i := 0; i < len(m.Question); i++ {
	// 		log.Printf("Question: %d\n %+v\n", i, m)
	// 	}

	// 	for i := 0; i < len(r.Answer); i++ {
	// 		log.Printf("Answer: %d\n %+v\n", i, r)
	// 	}
	// }

	r.Question[0].Name = strings.ToLower(r.Question[0].Name)
	for i := 0; i < len(r.Answer); i++ {
		r.Answer[i].Header().Name = strings.ToLower(r.Answer[i].Header().Name)
	}
	w.WriteMsg(r)

}

// TODO: Clean this up!!!
func dnsHandler(w dns.ResponseWriter, m *dns.Msg, proto string) {

	tStart := time.Now()

	// Evaluate this DNS request
	domain := strings.ToLower(m.Question[0].Name)
	clientAddr := StripPortFromAddr(strings.ToLower(w.RemoteAddr().String()))
	action, err := sys.RulesEngine().Evaluate(domain, clientAddr)

	if err != nil {
		log.Printf("Error while evaluating rules: %s\n", err.Error())
	}

	IPQuery := isIPQuery(m.Question[0])
	if isBlockingAction(action) && IPQuery > 0 {

		// perhaps call 'dns.HandleFailed(w,m)' instead
		writeBlockedRoute(w, m, IPQuery)
	} else {
		doDnsExchange(w, m, proto)
	}

	// Log this request
	duration := time.Since(tStart)
	perfItem := LogItem{
		HostToResolve: domain,
		RequestedBy:   clientAddr,
		Action:        action.String(),
		Duration:      duration.Seconds(),
	}

	//
	// Logging:
	//   <timestamp>
	//   <Host To Resolve>
	//   <Requested By>
	//   <Action>
	//   <Duration>
	//
	sys.PerfLog().WriteItem(&perfItem)

	//fmt.Printf("%s,%s,%s,%f\n", action.String(), domain, clientAddr, duration.Seconds())
}

func dnsUdpHandler(w dns.ResponseWriter, m *dns.Msg) {
	dnsHandler(w, m, "udp")
}

func dnsTcpHandler(w dns.ResponseWriter, m *dns.Msg) {
	dnsHandler(w, m, "tcp")
}
