package main

import (
	"errors"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var ErrDeviceNameNotFound = errors.New("Device name not found")
var ErrDeviceIPNotFound = errors.New("Device IP not found")

type DeviceCache struct {
	routerConfig Router
	routerClient RouterClient
	devFromName  map[string]RouterDevice
	devFromIP    map[string]RouterDevice
	lock         sync.Mutex
}

func NewDeviceCache(routerClient RouterClient, router Router) *DeviceCache {
	dc := DeviceCache{
		routerClient: routerClient,
		devFromName:  make(map[string]RouterDevice),
		devFromIP:    make(map[string]RouterDevice),
	}
	return &dc
}

func (dc *DeviceCache) Initialize() error {
	return dc.Refresh()
}

func (dc *DeviceCache) Refresh() error {
	err := dc.doRefresh()
	if err != nil {
		err = dc.reInitialize()
		if err != nil {
			log.Printf("[ERROR] DeviceCache::Refresh, Unable to reinitialize router, err: %s\n", err.Error())
		}
	}
	return err
}

func (dc *DeviceCache) reInitialize() error {
	return dc.routerClient.Login(dc.routerConfig.Host, dc.routerConfig.Port, dc.routerConfig.User, dc.routerConfig.Password)
}

func (dc *DeviceCache) doRefresh() error {

	devices, err := dc.routerClient.GetAttachedDeviceList()
	if err != nil {
		log.Printf("[ERROR] DeviceCache::Refresh, failed to retrieve list of attached devices: %s\n", err.Error())
		return err
	}

	dc.lock.Lock()
	defer dc.lock.Unlock()

	// set to table
	for _, d := range devices {
		// TODO: A device can have multiple IP's this will just go into a toggle!!!!
		dev, ok := dc.devFromName[d.Name]
		// Check for device changes and/or new devices
		if ok {
			// Device already in map..
			if !dev.IP.Equal(d.IP) {
				log.Printf("  %s changed IP, from: %s -> %s", dev.Name, dev.IP.String(), d.IP.String())
			}
		} else {
			log.Printf("New device: %s: %s", d.Name, d.IP.String())
		}
		dc.devFromName[d.Name] = d
		dc.devFromIP[d.IP.String()] = d
	}
	return nil
}

func (dc *DeviceCache) StartAutoRefresh(pollintervalsec int) {
	go func() {
		for {
			time.Sleep(time.Duration(pollintervalsec) * time.Second)
			dc.Refresh()
		}
	}()
}

func (dc *DeviceCache) NameToIP(name string) (net.IP, error) {
	dc.lock.Lock()
	defer dc.lock.Unlock()

	if d, ok := dc.devFromName[strings.ToLower(name)]; ok {
		return d.IP, nil
	}
	return nil, ErrDeviceNameNotFound
}

func (dc *DeviceCache) IPToName(ipaddr string) (string, error) {
	dc.lock.Lock()
	defer dc.lock.Unlock()

	if d, ok := dc.devFromIP[ipaddr]; ok {
		return d.Name, nil
	}
	return "", ErrDeviceIPNotFound
}

func (dc *DeviceCache) Dump() {
	dc.lock.Lock()
	defer dc.lock.Unlock()

	for _, d := range dc.devFromName {
		log.Printf("%s : %s\n", d.Name, d.IP.String())
	}
}
