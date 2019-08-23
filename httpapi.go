package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

const authHeaderToken = "X-Auth-Token"

type appHandler struct {
	Sys     *System
	Handler func(*System, http.ResponseWriter, *http.Request) (int, error)
}

// NewHTTPApi returns a gorilla mux router configured with the API
func NewHTTPApi(sys *System) (*mux.Router, error) {
	r := mux.NewRouter()

	mapDomain(r, sys)

	log.Printf("Adding path prefix to router!\n")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend/public")))
	return r, nil
}

func mapDomain(r *mux.Router, sys *System) {
	// Using new domain/persistence/db
	r.Handle("/config", appHandler{sys, getConfig}).Methods("GET")
	r.Handle("/config", appHandler{sys, postConfig}).Methods("POST")
	r.Handle("/devices", appHandler{sys, getDevices}).Methods("GET")
	r.Handle("/devices/{device}/rules", appHandler{sys, getDeviceRules}).Methods("GET")
	r.Handle("/devices/{device}/block", appHandler{sys, getDeviceBlock}).Methods("GET")
	r.Handle("/devices/{device}/release", appHandler{sys, getDeviceRelease}).Methods("GET")
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	tStart := time.Now()

	handlerName := runtime.FuncForPC(reflect.ValueOf(ah.Handler).Pointer()).Name()
	log.Printf("AppHandler::ServeHTTP, serving API request for: %s\n", handlerName)

	// usr, err := checkAuthTokenAndGetUser(r)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusForbidden)
	// 	return
	// }
	//log.Printf("AppHandler::ServeHTTP, token ok, user: %s\n", usr.Name)

	// Call handler and measure roundtrip
	status, err := ah.Handler(ah.Sys, w, r) // Call handler -> this calls the actuall API function

	// logItem := logging.PerfLogItem{
	// 	UserAgent:               r.UserAgent(),
	// 	RequestURI:              r.RequestURI,
	// 	Host:                    r.Host,
	// 	HTTPMethod:              r.Method,
	// 	HTTPReturnStatus:        status,
	// 	HandlerMethod:           handlerName,
	// 	HandlerError:            errString,
	// 	HandlerExecutionTimeSec: duration.Seconds(),
	// }

	//uriParts := strings.Split(r.RequestURI, "/")

	// URL's can be of two forms
	// 1) /api/<version>/<function>/params
	// 2) /<function>/params
	//
	// the '2' item is for generic items and test stuff - not bound to an API Versioning
	// Params are optional, hence '2' can have 2 parts as minimum length
	//
	// if (len(uriParts) > 2) && (uriParts[1] == "api") {
	// 	// Trust the compiler to optimize this - I want the ability to debug simply
	// 	apiVersion := uriParts[2]
	// 	apiDomain := uriParts[3]
	// 	apiFunction := uriParts[4]

	// 	logItem.APIVersion = apiVersion
	// 	logItem.APIDomain = apiDomain
	// 	logItem.APIFunction = apiFunction
	// } else {
	// 	// this is a generic or temporary call
	// 	logItem.APIVersion = "none"
	// 	logItem.APIDomain = "generic"
	// 	logItem.APIFunction = uriParts[1]
	// }

	// _ = ah.perflog.SendJSON(logItem)
	duration := time.Since(tStart)
	log.Printf("AppHandler::ServeHTTP, done - duration: %f\n", duration.Seconds())

	//	fmt.Printf("After\n")
	if err != nil {
		// TODO: Log this error to a nagini_se_systemlog  (or use something else)
		log.Println(err)

		switch status {
		case http.StatusNotFound:
			http.Error(w, http.StatusText(status), http.StatusInternalServerError)
			break
		case http.StatusInternalServerError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			break
		default:
			//http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			http.Error(w, err.Error(), status)
		}
	}
}

func getConfig(sys *System, w http.ResponseWriter, r *http.Request) (int, error) {
	// user, err := checkAuthTokenAndGetUser(r)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }
	// return sendJSONResponse(w, user)
	log.Printf("getConfig\n")
	return http.StatusOK, nil
}

func postConfig(sys *System, w http.ResponseWriter, r *http.Request) (int, error) {
	log.Printf("postConfig\n")
	return http.StatusOK, nil
}

func getDevices(sys *System, w http.ResponseWriter, r *http.Request) (int, error) {
	// user, err := checkAuthTokenAndGetUser(r)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }
	// return sendJSONResponse(w, user)

	log.Printf("getDevices\n")
	//devices := sys.DeviceCache().Devices()
	devices := getDeviceData(sys)
	return sendJSONResponse(w, devices)
}

type DeviceDTO struct {
	Device RouterDevice
	Host   Host
}

func getDeviceData(sys *System) []DeviceDTO {
	dto := make([]DeviceDTO, 0)
	devices := sys.DeviceCache().Devices()
	for _, d := range devices {
		h, _ := sys.RulesEngine().HostFromName(d.Name)
		if h != nil {
			data := DeviceDTO{
				Device: d,
				Host:   *h,
			}
			dto = append(dto, data)
		}
	}
	return dto
}

func getDeviceRules(sys *System, w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	deviceName := vars["device"]
	d, err := sys.DeviceCache().DeviceFromName(deviceName)
	if err != nil {
		log.Printf("Unable to find '%s' in device cache\n", deviceName)
		return http.StatusInternalServerError, err
	}

	h, err := sys.RulesEngine().HostFromName(d.Name)
	if err != nil {
		// If no rule is found in the host section of the config - default rule applies
		log.Printf("No specific host rules defined for '%s', using default!\n", d.Name)
		return http.StatusInternalServerError, err
	}
	log.Printf("%+v\n", h)
	return sendJSONResponse(w, h)
}

func getDeviceBlock(sys *System, w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	deviceName := vars["device"]
	d, err := sys.DeviceCache().DeviceFromName(deviceName)
	if err != nil {
		log.Printf("Unable to find '%s' in device cache\n", deviceName)
		return http.StatusInternalServerError, err
	}

	h, err := sys.RulesEngine().HostFromName(d.Name)
	if err != nil {
		// If not specific host rule is found - we can't block!!
		// TODO: Fix this - insert specific host rule!
		log.Printf("No specific host rules defined for '%s', using can't block device!\n", d.Name)
		return http.StatusInternalServerError, err
	}

	log.Printf("Blocking device '%s'\n", d.Name)
	h.Block()

	return sendJSONResponse(w, h)

}

func getDeviceRelease(sys *System, w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	deviceName := vars["device"]
	d, err := sys.DeviceCache().DeviceFromName(deviceName)
	if err != nil {
		log.Printf("Unable to find '%s' in device cache\n", deviceName)
		return http.StatusInternalServerError, err
	}

	h, err := sys.RulesEngine().HostFromName(d.Name)
	if err != nil {
		// If not specific host rule is found - we can't block!!
		// TODO: Fix this - insert specific host rule!
		log.Printf("No specific host rules defined for '%s', using can't block device!\n", d.Name)
		return http.StatusInternalServerError, err
	}

	log.Printf("Unblocking device '%s'\n", d.Name)
	h.Unblock()

	return sendJSONResponse(w, h)
}

//
// Generic functions, marshals something to JSON and sends it back to the web client
//
func sendJSONResponse(w http.ResponseWriter, data interface{}) (int, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return sendRawAsJSONResponse(w, payload)
}

func sendRawAsJSONResponse(w http.ResponseWriter, payload []byte) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))

	return http.StatusOK, nil
}
