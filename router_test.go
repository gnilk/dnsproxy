package main

import (
	"testing"
)

var glbRouter RouterClient

func getRouter() RouterClient {
	if glbRouter != nil {
		return glbRouter
	}
	glbRouter = NewNetGearRouterClient()
	return glbRouter
}

func TestRouterLogin(t *testing.T) {
	router := getRouter()
	err := router.Login("192.168.1.1", "80", "admin", "neger6slakt")
	if err != nil {
		t.Error(err)
	}
}

func TestRouterGetAttachedDevices(t *testing.T) {
	router := getRouter()
	if router == nil {
		t.Error("No router")
	}
	_, err := router.GetAttachedDeviceList()
	if err != nil {
		t.Error(err)
	}
}
