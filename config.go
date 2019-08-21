package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
)

//
// this code is generated by the modelgenerator
// data model source = config_datamodel.xml
//

type ActionType int64

const (
	_                                        = iota
	ActionTypeNone                ActionType = 0
	ActionTypePass                ActionType = 1
	ActionTypeBlockedDevice       ActionType = 2
	ActionTypeBlockedSiteBan      ActionType = 3
	ActionTypeBlockedTimeSpan     ActionType = 4
	ActionTypePassAccumulatedTime ActionType = 5
)

var mapActionTypeToName = map[ActionType]string{
	0: "ActionTypeNone",
	1: "ActionTypePass",
	2: "ActionTypeBlockedDevice",
	3: "ActionTypeBlockedSiteBan",
	4: "ActionTypeBlockedTimeSpan",
	5: "ActionTypePassAccumulatedTime",
}

var mapActionTypeToValue = map[string]ActionType{
	"ActionTypeNone":                0,
	"ActionTypePass":                1,
	"ActionTypeBlockedDevice":       2,
	"ActionTypeBlockedSiteBan":      3,
	"ActionTypeBlockedTimeSpan":     4,
	"ActionTypePassAccumulatedTime": 5,
}

func (this *ActionType) String() string {
	return mapActionTypeToName[*this]
}

func (this ActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.String())
}

func (this *ActionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ActionType should be a string")
	}
	v, ok := mapActionTypeToValue[s]
	if !ok {
		return fmt.Errorf("invalid ActionType")
	}
	*this = v
	return nil
}

type RouterType int64

const (
	_                            = iota
	RouterTypeNone    RouterType = 0
	RouterTypeNetGear RouterType = 1
	RouterTypeUnifi   RouterType = 2
)

var mapRouterTypeToName = map[RouterType]string{
	0: "RouterTypeNone",
	1: "RouterTypeNetGear",
	2: "RouterTypeUnifi",
}

var mapRouterTypeToValue = map[string]RouterType{
	"RouterTypeNone":    0,
	"RouterTypeNetGear": 1,
	"RouterTypeUnifi":   2,
}

func (this *RouterType) String() string {
	return mapRouterTypeToName[*this]
}

func (this RouterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.String())
}

func (this *RouterType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RouterType should be a string")
	}
	v, ok := mapRouterTypeToValue[s]
	if !ok {
		return fmt.Errorf("invalid RouterType")
	}
	*this = v
	return nil
}

//
// Config is generated
//
type Config struct {
	Logfile          string
	Router           Router
	ListenAddress    string
	NameServers      []NameServer
	DefaultRule      ActionType
	OnErrorRule      ActionType
	IPv4BlockResolve string
	IPv6BlockResolve string
	Hosts            []Host
	Domains          []Domain
}

func (this *Config) GetLogfile() string {
	return this.Logfile
}

func (this *Config) SetLogfile(value string) {
	this.Logfile = value
}

func (this *Config) GetRouter() Router {
	return this.Router
}

func (this *Config) SetRouter(value Router) {
	this.Router = value
}

func (this *Config) GetListenAddress() string {
	return this.ListenAddress
}

func (this *Config) SetListenAddress(value string) {
	this.ListenAddress = value
}

func (this *Config) GetNameServersAsRef() []NameServer {
	return this.NameServers[:len(this.NameServers)]
}

func (this *Config) GetNameServersAsCopy() []NameServer {
	newSlice := make([]NameServer, len(this.NameServers))
	copy(newSlice, this.NameServers)
	return newSlice
}

func (this *Config) SetNameServers(value []NameServer) {
	this.NameServers = make([]NameServer, len(value))
	copy(this.NameServers, value)
}

func (this *Config) GetDefaultRule() ActionType {
	return this.DefaultRule
}

func (this *Config) SetDefaultRule(value ActionType) {
	this.DefaultRule = value
}

func (this *Config) GetOnErrorRule() ActionType {
	return this.OnErrorRule
}

func (this *Config) SetOnErrorRule(value ActionType) {
	this.OnErrorRule = value
}

func (this *Config) GetIPv4BlockResolve() string {
	return this.IPv4BlockResolve
}

func (this *Config) SetIPv4BlockResolve(value string) {
	this.IPv4BlockResolve = value
}

func (this *Config) GetIPv6BlockResolve() string {
	return this.IPv6BlockResolve
}

func (this *Config) SetIPv6BlockResolve(value string) {
	this.IPv6BlockResolve = value
}

func (this *Config) GetHostsAsRef() []Host {
	return this.Hosts[:len(this.Hosts)]
}

func (this *Config) GetHostsAsCopy() []Host {
	newSlice := make([]Host, len(this.Hosts))
	copy(newSlice, this.Hosts)
	return newSlice
}

func (this *Config) SetHosts(value []Host) {
	this.Hosts = make([]Host, len(value))
	copy(this.Hosts, value)
}

func (this *Config) GetDomainsAsRef() []Domain {
	return this.Domains[:len(this.Domains)]
}

func (this *Config) GetDomainsAsCopy() []Domain {
	newSlice := make([]Domain, len(this.Domains))
	copy(newSlice, this.Domains)
	return newSlice
}

func (this *Config) SetDomains(value []Domain) {
	this.Domains = make([]Domain, len(value))
	copy(this.Domains, value)
}

// ToJSON creates a JSON representation of the data for the type
func (this *Config) ToJSON() string {
	b, err := json.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// ToXML creates an XML representation of the data for the type
func (this *Config) ToXML() string {
	b, err := xml.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// ConfigFromJSON converts a JSON representation to the data type
func ConfigFromJSON(jsondata string) (*Config, error) {
	var value Config
	err := json.Unmarshal([]byte(jsondata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// ConfigFromXML converts an XML representation to the type
func ConfigFromXML(xmldata string) (*Config, error) {
	var value Config
	err := xml.Unmarshal([]byte(xmldata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

//
// Router is generated
//
type Router struct {
	Host         string
	Port         string
	User         string
	Password     string
	Engine       RouterType
	PollChanges  bool
	PollInterval int
	TimeoutSec   int
}

func (this *Router) GetHost() string {
	return this.Host
}

func (this *Router) SetHost(value string) {
	this.Host = value
}

func (this *Router) GetPort() string {
	return this.Port
}

func (this *Router) SetPort(value string) {
	this.Port = value
}

func (this *Router) GetUser() string {
	return this.User
}

func (this *Router) SetUser(value string) {
	this.User = value
}

func (this *Router) GetPassword() string {
	return this.Password
}

func (this *Router) SetPassword(value string) {
	this.Password = value
}

func (this *Router) GetEngine() RouterType {
	return this.Engine
}

func (this *Router) SetEngine(value RouterType) {
	this.Engine = value
}

func (this *Router) GetPollChanges() bool {
	return this.PollChanges
}

func (this *Router) SetPollChanges(value bool) {
	this.PollChanges = value
}

func (this *Router) GetPollInterval() int {
	return this.PollInterval
}

func (this *Router) SetPollInterval(value int) {
	this.PollInterval = value
}

func (this *Router) GetTimeoutSec() int {
	return this.TimeoutSec
}

func (this *Router) SetTimeoutSec(value int) {
	this.TimeoutSec = value
}

// ToJSON creates a JSON representation of the data for the type
func (this *Router) ToJSON() string {
	b, err := json.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// ToXML creates an XML representation of the data for the type
func (this *Router) ToXML() string {
	b, err := xml.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// RouterFromJSON converts a JSON representation to the data type
func RouterFromJSON(jsondata string) (*Router, error) {
	var value Router
	err := json.Unmarshal([]byte(jsondata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// RouterFromXML converts an XML representation to the type
func RouterFromXML(xmldata string) (*Router, error) {
	var value Router
	err := xml.Unmarshal([]byte(xmldata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

//
// NameServer is generated
//
type NameServer struct {
	IP string
}

func (this *NameServer) GetIP() string {
	return this.IP
}

func (this *NameServer) SetIP(value string) {
	this.IP = value
}

// ToJSON creates a JSON representation of the data for the type
func (this *NameServer) ToJSON() string {
	b, err := json.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// ToXML creates an XML representation of the data for the type
func (this *NameServer) ToXML() string {
	b, err := xml.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// NameServerFromJSON converts a JSON representation to the data type
func NameServerFromJSON(jsondata string) (*NameServer, error) {
	var value NameServer
	err := json.Unmarshal([]byte(jsondata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// NameServerFromXML converts an XML representation to the type
func NameServerFromXML(xmldata string) (*NameServer, error) {
	var value NameServer
	err := xml.Unmarshal([]byte(xmldata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

//
// Domain is generated
//
type Domain struct {
	Name  string
	Hosts []Host
}

func (this *Domain) GetName() string {
	return this.Name
}

func (this *Domain) SetName(value string) {
	this.Name = value
}

func (this *Domain) GetHostsAsRef() []Host {
	return this.Hosts[:len(this.Hosts)]
}

func (this *Domain) GetHostsAsCopy() []Host {
	newSlice := make([]Host, len(this.Hosts))
	copy(newSlice, this.Hosts)
	return newSlice
}

func (this *Domain) SetHosts(value []Host) {
	this.Hosts = make([]Host, len(value))
	copy(this.Hosts, value)
}

// ToJSON creates a JSON representation of the data for the type
func (this *Domain) ToJSON() string {
	b, err := json.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// ToXML creates an XML representation of the data for the type
func (this *Domain) ToXML() string {
	b, err := xml.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// DomainFromJSON converts a JSON representation to the data type
func DomainFromJSON(jsondata string) (*Domain, error) {
	var value Domain
	err := json.Unmarshal([]byte(jsondata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// DomainFromXML converts an XML representation to the type
func DomainFromXML(xmldata string) (*Domain, error) {
	var value Domain
	err := xml.Unmarshal([]byte(xmldata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

//
// Host is generated
//
type Host struct {
	Name  string
	Rules []Rule
}

func (this *Host) GetName() string {
	return this.Name
}

func (this *Host) SetName(value string) {
	this.Name = value
}

func (this *Host) GetRulesAsRef() []Rule {
	return this.Rules[:len(this.Rules)]
}

func (this *Host) GetRulesAsCopy() []Rule {
	newSlice := make([]Rule, len(this.Rules))
	copy(newSlice, this.Rules)
	return newSlice
}

func (this *Host) SetRules(value []Rule) {
	this.Rules = make([]Rule, len(value))
	copy(this.Rules, value)
}

// ToJSON creates a JSON representation of the data for the type
func (this *Host) ToJSON() string {
	b, err := json.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// ToXML creates an XML representation of the data for the type
func (this *Host) ToXML() string {
	b, err := xml.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// HostFromJSON converts a JSON representation to the data type
func HostFromJSON(jsondata string) (*Host, error) {
	var value Host
	err := json.Unmarshal([]byte(jsondata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// HostFromXML converts an XML representation to the type
func HostFromXML(xmldata string) (*Host, error) {
	var value Host
	err := xml.Unmarshal([]byte(xmldata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

//
// Rule is generated
//
type Rule struct {
	Type     ActionType
	TimeSpan string
	MaxTime  string
}

func (this *Rule) GetType() ActionType {
	return this.Type
}

func (this *Rule) SetType(value ActionType) {
	this.Type = value
}

func (this *Rule) GetTimeSpan() string {
	return this.TimeSpan
}

func (this *Rule) SetTimeSpan(value string) {
	this.TimeSpan = value
}

func (this *Rule) GetMaxTime() string {
	return this.MaxTime
}

func (this *Rule) SetMaxTime(value string) {
	this.MaxTime = value
}

// ToJSON creates a JSON representation of the data for the type
func (this *Rule) ToJSON() string {
	b, err := json.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// ToXML creates an XML representation of the data for the type
func (this *Rule) ToXML() string {
	b, err := xml.MarshalIndent(this, "", "    ")
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

// RuleFromJSON converts a JSON representation to the data type
func RuleFromJSON(jsondata string) (*Rule, error) {
	var value Rule
	err := json.Unmarshal([]byte(jsondata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// RuleFromXML converts an XML representation to the type
func RuleFromXML(xmldata string) (*Rule, error) {
	var value Rule
	err := xml.Unmarshal([]byte(xmldata), &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}
