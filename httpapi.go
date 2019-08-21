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

func getDevices(sys *System, w http.ResponseWriter, r *http.Request) (int, error) {
	// user, err := checkAuthTokenAndGetUser(r)
	// if err != nil {
	// 	return http.StatusInternalServerError, err
	// }
	// return sendJSONResponse(w, user)

	log.Printf("getDevices\n")
	devices := sys.DeviceCache().Devices()
	return sendJSONResponse(w, devices)
}

func postConfig(sys *System, w http.ResponseWriter, r *http.Request) (int, error) {
	log.Printf("postConfig\n")
	return http.StatusOK, nil
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
