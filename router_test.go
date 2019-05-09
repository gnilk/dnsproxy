package main

import (
	"log"
	"testing"
)

var glbRouter RouterClient

func getRouter() RouterClient {
	if glbRouter != nil {
		return glbRouter
	}
	glbRouter = NewUnifiRouterClient()
	return glbRouter
}

func TestRouterLogin(t *testing.T) {
	router := getRouter()
	err := router.Login("192.168.1.46", "8443", "Fredrik", "neger6slakt")
	if err != nil {
		t.Error(err)
	}
}

func TestRouterGetAttachedDevices(t *testing.T) {
	router := getRouter()
	if router == nil {
		t.Error("No router")
	}
	devices, err := router.GetAttachedDeviceList()
	if err != nil {
		t.Error(err)
	}
	log.Println(devices)
}
