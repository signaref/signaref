// PT-Server project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"logutil"
	"mgoSession"
	"middleware"
	"models"
	"net/http"
	"os"
	"time"
	_ "utils"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
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

		router := mux.NewRouter().StrictSlash(true)

		router.Handle("/um/scopes/", middleware.MiddlewareHandle(GDAll("scopes")))

		router.Handle("/um/scopes/{id}", middleware.MiddlewareHandle(ScopeGPPD("scopes")))

		fs := http.FileServer(http.Dir(""))

		router.Handle("/", middleware.MiddlewareHandle(fs))

		http.ListenAndServe(Addr, router)

	}
}

func ScopeGPPD(collection string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("yes")
		if r.Method == "GET" {
			GetBy(id, collection, w, r)
		} else if r.Method == "DELETE" {
			DeleteBy(id, collection, w, r)
		} else if r.Method == "PUT" {
			UpdateScopeBy(id, collection, w, r)
		} else if r.Method == "POST" {
			CreateScope(collection, w, r)
		} else {
			WriteResponseMessage(w, "not implemented", "not implemented", 200, true)
		}
	})
}

func GDAll(collection string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			GetAll(collection, w, r)
		} else if r.Method == "DELETE" {
			DeleteAll(collection, w, r)
		} else {
			WriteResponseMessage(w, "not implemented", "not implemented", 200, true)
		}
	})
}

func CreateScope(collection string, w http.ResponseWriter, r *http.Request) {
	var buf []byte
	var S models.Scope

	if r.Header.Get("Content-Type") != "application/json" {
		WriteResponseMessage(w, "content type should be a in json", "content type should be a in json", 400, false)
		return
	}

	buf, err = ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponseMessage(w, "error in reading data from request body", err.Error(), 400, false)
		return
	}

	err = json.Unmarshal(buf, &S)
	if err != nil {
		WriteResponseMessage(w, "error in converting the data to object", err.Error(), 400, false)
		return
	}

	s := make(map[string]interface{}, 1)
	s["scope"] = S.Scope
	result, _ := dbSession.FindByQuery("scopes", s)

	if len(result) > 0 {
		WriteResponseMessage(w, "scope alredy existed in the system", "trainer name alredy existed in the system", 400, false)
		return
	}

	err = dbSession.Insert("scopes", S)

	if err != nil {
		WriteResponseMessage(w, "error in insering the data", err.Error(), 400, false)
		return
	}

	WriteResponseMessage(w, "Scope successfully added", "1", 200, true)
}

func UpdateScopeBy(id, collection string, w http.ResponseWriter, r *http.Request) {
	var buf []byte
	var T models.Scope

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

	if id == "" {
		WriteResponseMessage(w, "Id must be supplied to update", "Id must be supplied to update", 400, false)
		return
	}
	//t.ID = id
	//M.MovieId = utils.GUID()
	fmt.Println(T)
	err = dbSession.UpdateByID(collection, bson.ObjectIdHex(id), &T)
	if err != nil {
		WriteResponseMessage(w, "error in updating the data", err.Error(), 400, false)
		return
	}

	WriteResponseMessage(w, "Trainer successfully updated", "1", 200, true)
}

//returns all records from the collection provided by the collection name as the parameter
func GetAll(collection string, w http.ResponseWriter, r *http.Request) {
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
}

func DeleteAll(collection string, w http.ResponseWriter, r *http.Request) {
	err := dbSession.DeleteAll(collection)
	if err != nil {
		WriteResponseMessage(w, "error in deleting the data", err.Error(), 400, false)
	} else {
		WriteResponseMessage(w, "entire collection has been are successfully removed", "1", 200, true)
	}
}

func GetBy(id, collection string, w1 http.ResponseWriter, r1 *http.Request) {
	result, err := dbSession.ListByID(collection, id)
	if err != nil {
		w1.WriteHeader(400)
		fmt.Fprintln(w1, err.Error())
	} else {
		jData, err := json.Marshal(result)
		if err != nil {
			w1.WriteHeader(400)
			fmt.Fprintln(w1, err.Error())
		}
		w1.Header().Set("Content-Type", "application/json")
		w1.Write(jData)
	}
}

func DeleteBy(id, collection string, w1 http.ResponseWriter, r1 *http.Request) {
	err := dbSession.DeleteByID(collection, id)
	if err != nil {
		WriteResponseMessage(w1, "error in deleting the data", err.Error(), 400, false)
		return
	}
	WriteResponseMessage(w1, "Trainer successfully removed", "1", 200, true)
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
