package main

import (
	"errors"
	"log"
	"net"
)

var ErrDeviceNameNotFound = errors.New("Device name not found")
var ErrDeviceIPNotFound = errors.New("Device IP not found")

type DeviceCache struct {
	routerClient RouterClient
	devFromName  map[string]RouterDevice
	devFromIP    map[string]RouterDevice
}

func NewDeviceCache(routerClient RouterClient) *DeviceCache {
	dc := DeviceCache{
		routerClient: routerClient,
		devFromName:  make(map[string]RouterDevice),
		devFromIP:    make(map[string]RouterDevice),
	}
	return &dc
}

func (dc *DeviceCache) Initialize() error {
	devices, err := dc.routerClient.GetAttachedDeviceList()
	if err != nil {
		log.Printf("[ERROR] DeviceCache::Initialize, failed to retrieve list of attached devices: %s\n", err.Error())
		return err
	}
	// set to table
	for _, d := range devices {
		dc.devFromName[d.Name] = d
		dc.devFromIP[d.IP.String()] = d
	}
	return nil
}

func (dc *DeviceCache) NameToIP(name string) (net.IP, error) {
	if d, ok := dc.devFromName[name]; ok {
		return d.IP, nil
	}
	return nil, ErrDeviceNameNotFound
}

func (dc *DeviceCache) IPToName(ipaddr string) (string, error) {
	if d, ok := dc.devFromIP[ipaddr]; ok {
		return d.Name, nil
	}
	return "", ErrDeviceIPNotFound
}

func (dc *DeviceCache) Dump() {
	for _, d := range dc.devFromName {
		log.Printf("%s : %s\n", d.Name, d.IP.String())
	}
}
