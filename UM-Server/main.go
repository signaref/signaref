// PT-Server project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"logutil"
	"mgoSession"
	"models"
	"net/http"
	"os"
	"strings"
	"time"
	"utils"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var (
	DBConnection string
)

var (
	Addr      string
	err       error
	logger    *logutil.LoggerUtil
	dbSession *mgoSession.Session
)

//init all propertie values from config/app.toml file
func init() {

	viper.SetConfigName("app")    // no need to include file extension
	viper.AddConfigPath("config") // set the path of your config file

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		Addr = viper.GetString("connections.Address")
		DBConnection = viper.GetString("connections.DBConnection")

	}
}

//init logger
func init() {
	file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file :", err)
	}
	logger = logutil.NewLogger(file, "log: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetOutput(file)
}

func main() {
	dbSession, err = mgoSession.New(DBConnection, "local")

	if err != nil {
		log.Fatal("mongodb database is not connected")
	} else {
		log.Println("mongodb session has been created")

		dataServer := http.NewServeMux()

		dataServer.Handle("/Trainers", MiddlewareHandle(GetAll("trainers")))
		dataServer.Handle("/AddTrainer", MiddlewareHandle(AddTrainer()))
		dataServer.Handle("/UpdateTrainer", MiddlewareHandle(UpdateTrainer()))
		dataServer.Handle("/TrainerbyKeyword", MiddlewareHandle(GetTrainerByKeyword()))
		dataServer.Handle("/TrainerbySkills", MiddlewareHandle(GetTrainerBySkills()))
		dataServer.Handle("/Trainers/RemoveById", MiddlewareHandle(removeByID("trainers")))
		dataServer.Handle("/Trainers/RemoveAll", MiddlewareHandle(removeAll("trainers")))

		fs := http.FileServer(http.Dir(""))

		dataServer.Handle("/", MiddlewareHandle(fs))

		http.ListenAndServe(Addr, dataServer)

	}
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

		WriteResponseMessage(w, "Unathenticated request", "", 400, true)
	})
}

func AddCourse() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func AddTrainer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			WriteResponseMessage(w, "this action works only for Post method", "action works only for Post method", 400, false)
			return
		}
		var buf []byte
		var T models.Trainer

		if r.Header.Get("Content-Type") != "application/json" {
			WriteResponseMessage(w, "content type should be a in json", "content type should be a in json", 400, false)
			return
		}

		buf, err = ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponseMessage(w, "error in reading data from request body", err.Error(), 400, false)
			return
		}

		err = json.Unmarshal(buf, &T)
		if err != nil {
			WriteResponseMessage(w, "error in converting the data to object", err.Error(), 400, false)
			return
		}

		t := make(map[string]interface{}, 1)
		t["name"] = T.Name
		result, _ := dbSession.FindByQuery("trainers", t)

		if len(result) > 0 {
			WriteResponseMessage(w, "trainer name are alredy existed in the system", "trainer name alredy existed in the system", 400, false)
			return
		}

		/*if models.ValidateMovie(M) != "" {
			WriteResponseMessage(w, models.ValidateMovie(M), models.ValidateMovie(M), 400, false)
			return
		}*/

		//M.MovieId = utils.GUID()
		err = dbSession.Insert("trainers", T)

		if err != nil {
			WriteResponseMessage(w, "error in insering the data", err.Error(), 400, false)
			return
		}

		WriteResponseMessage(w, "Trainer successfully added", "1", 200, true)
	})
}

func UpdateTrainer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			WriteResponseMessage(w, "this action works only for Post method", "action works only for Post method", 400, false)
			return
		}
		var buf []byte
		var T models.Trainer

		if r.Header.Get("Content-Type") != "application/json" {
			WriteResponseMessage(w, "content type should be a in json", "content type should be a in json", 400, false)
			return
		}

		buf, err = ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponseMessage(w, "error in reading data from request body", err.Error(), 400, false)
			return
		}

		err = json.Unmarshal(buf, &T)
		if err != nil {
			WriteResponseMessage(w, "error in converting the data to object", err.Error(), 400, false)
			return
		}

		/*t := make(map[string]interface{}, 1)
		t["name"] = T.Name
		result, _ := dbSession.FindByQuery("trainers", t)

		if len(result) > 0 {
			WriteResponseMessage(w, "trainer name are alredy existed in the system", "trainer name alredy existed in the system", 400, false)
			return
		}*/

		/*if models.ValidateMovie(M) != "" {
			WriteResponseMessage(w, models.ValidateMovie(M), models.ValidateMovie(M), 400, false)
			return
		}*/

		if T.ID == "" {
			WriteResponseMessage(w, "Id must be supplied to update", "Id must be supplied to update", 400, false)
			return
		}

		//M.MovieId = utils.GUID()
		err = dbSession.UpdateByID("trainers", T.ID, T)

		if err != nil {
			WriteResponseMessage(w, "error in updating the data", err.Error(), 400, false)
			return
		}

		WriteResponseMessage(w, "Trainer successfully updated", "1", 200, true)
	})
}

//Get a trainers by any field that is existed in the system.
//Query string contains key=field-name and value=value of the search..[It works as like operator]
func GetTrainerByKeyword() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			WriteResponseMessage(w, "this action works only with get method", "action works only for Post method", 400, false)
			return
		}
		vars := r.URL.Query()
		key := vars["key"]
		value := vars["value"]
		if len(key) < 1 || len(value) < 1 {
			WriteResponseMessage(w, "key and(or) value query string parameters are not passed", "key and(or) value query string parameters are not passed", 400, false)
			return
		}
		result, err := dbSession.FindByRegEx("trainers", key[0], value[0])
		if err != nil {
			WriteResponseMessage(w, "error in db connection", err.Error(), 400, false)
			return
		}
		if len(result) < 1 {
			WriteResponseMessage(w, "no data available", "no data available", 400, false)
			return
		}
		jData, err := json.Marshal(result)
		if err != nil {
			WriteResponseMessage(w, "error in reading the result", err.Error(), 400, false)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	})
}

func GetTrainerBySkills() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			WriteResponseMessage(w, "this action works only with get method", "action works only for Post method", 400, false)
			return
		}
		vars := r.URL.Query()
		value := vars["skills"]

		if len(value) < 1 {
			WriteResponseMessage(w, "key and(or) value query string parameters are not passed", "key and(or) value query string parameters are not passed", 400, false)
			return
		}

		vals := strings.Split(value[0], ",")
		result, err := dbSession.FindByArrayAll("trainers", "skills", vals)
		if err != nil {
			WriteResponseMessage(w, "error in db connection", err.Error(), 400, false)
			return
		}
		if len(result) < 1 {
			WriteResponseMessage(w, "no data available", "no data available", 400, false)
			return
		}
		jData, err := json.Marshal(result)
		if err != nil {
			WriteResponseMessage(w, "error in reading the result", err.Error(), 400, false)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	})
}

//returns all records from the collection provided by the collection name as the parameter
func GetAll(collection string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			WriteResponseMessage(w, "this action works only with get method", "action works only for Post method", 400, false)
			return
		}
		result, err := dbSession.FindAll(collection)

		if err != nil {
			WriteResponseMessage(w, "error in db connection", err.Error(), 400, false)
			return
		}

		if len(result) < 1 {
			WriteResponseMessage(w, "no data available", "no data available", 400, false)
			return
		}
		jData, err := json.Marshal(result)
		if err != nil {
			WriteResponseMessage(w, "error in reading the result", err.Error(), 400, false)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	})
}

//fetches document from the collection by _id (bson id).
//collection name as the parameter
func GetByID(collection string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			vars := mux.Vars(r)
			_id := vars["id"]

			if _id == "" {
				WriteResponseMessage(w, "id has not passed as a parameter", "id has not passed as a parameter", 400, false)
				return
			}

			result, err := dbSession.ListByID(collection, _id)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, err.Error())
			} else {

				jData, err := json.Marshal(result)
				if err != nil {
					w.WriteHeader(400)
					fmt.Fprintln(w, err.Error())
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(jData)
			}
		}
	})
}

func removeByID(collection string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			vals := r.URL.Query()
			_id := vals.Get("id")
			if _id == "" {
				WriteResponseMessage(w, "id has not passed as a parameter", "id has not passed as a parameter", 400, false)
				return
			}
			err := dbSession.DeleteByID(collection, _id)

			if err != nil {
				WriteResponseMessage(w, "error in deleting the data", err.Error(), 400, false)
				return
			}
			WriteResponseMessage(w, "Trainer successfully removed", "1", 200, true)
		}
	})
}

func removeAll(collection string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			err := dbSession.DeleteAll(collection)
			if err != nil {
				WriteResponseMessage(w, "error in deleting the data", err.Error(), 400, false)
			} else {
				WriteResponseMessage(w, "All Trainers are successfully removed", "1", 200, true)
			}

		}
	})
}

//write response to the responsewrite in json fomat.
func WriteResponseMessage(w http.ResponseWriter, message string, trace string, status int, success bool) {
	msg := models.Message{MSG: message, Success: success, Trace: trace, Status: status}
	go logger.LogString(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(msg.Bytes())
}

func WriteLog(status, user, source, message string) {
	l := models.Log{TimeStamp: time.Now().String(), Status: status, Source: source, Message: message}
	go logger.LogString(l)
}
