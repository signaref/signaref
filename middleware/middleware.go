// middleware project middleware.go
package middleware

import (
	"log"
	"logutil"
	"models"
	"net/http"
	"os"
	"utils"
)

var (
	logger *logutil.LoggerUtil
)

//init logger
func init() {
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file :", err)
	}
	logger = logutil.NewLogger(file, "log: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetOutput(file)
}

//retunes a http.handler based on middleware conditions..
//if all are passed the argument next is returned. else other handlers are retunred
//for example unathenticated handler etc..
func MiddlewareHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.WriteLog("Call On URI:", r.Host, r.URL, " from the following IP address:", utils.GetIpAddr(r))
		if IsAuthenticated(r) {
			next.ServeHTTP(w, r)
		} else {
			logger.WriteLog("Call On ", r.Host, r.URL, " from the following IP address:", utils.GetIpAddr(r), " is not authenticated")
			Unauthenticated().ServeHTTP(w, r)
		}
	})
}

//returns true when athenticated false when not
//athentication headers are obtained from r
func IsAuthenticated(r *http.Request) bool {
	//todo write logic here
	return true
}

//when client is unathenticated this handler automatically called
//this is called from the moddlewarehandle.
func Unauthenticated() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteResponseMessage(w, "Unathenticated request", "", 200, true)
	})
}

//write response to the responsewrite in json fomat.
func WriteResponseMessage(w http.ResponseWriter, message string, trace string, status int, success bool) {
	msg := models.Message{MSG: message, Success: success, Trace: trace, Status: status}
	go logger.LogString(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(msg.Bytes())
}
