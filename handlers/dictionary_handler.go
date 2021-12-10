//handlers package includes handlers of /api/get /api/set and /api/flush endpoints
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//This is the log file of http requests
const logFile string = "log.txt"

type DictionaryHandlers struct {
	Data map[string]string `json:"data"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//GetValue returns key and value of key in the request in KeyValue struct || {"key":"${key}":"value":"{value}"}
func (d *DictionaryHandlers) GetValue(w http.ResponseWriter, r *http.Request) {
	logRequests(r, logFile)
	if r.Method == "GET" {
		key := r.URL.Query()["Key"]
		if len(key) == 0 {
			dictionary, _ := json.Marshal(d)
			fmt.Fprint(w, string(dictionary))
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			return
		} else if len(key) == 1 {
			if _, ok := d.Data[key[0]]; ok {
				keyValue := KeyValue{Key: key[0], Value: d.Data[key[0]]}
				value, _ := json.Marshal(keyValue)
				w.Write(value)
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

//SetValue adds given key and value in request to DictionaryHandler.Data map and returns http.StatusOK
func (d *DictionaryHandlers) SetValue(w http.ResponseWriter, r *http.Request) {
	logRequests(r, logFile)
	if r.Method == "POST" {
		key := r.URL.Query()["Key"]
		value := r.URL.Query()["Value"]
		if len(key) == 0 || len(value) == 0 {
			w.WriteHeader(http.StatusBadRequest)
		} else if len(key) == 1 && len(value) == 1 {
			d.Data[key[0]] = value[0]
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//FlushDictionary deletes all data in dictionary.json file DictionaryHandlers.Data map and returns http.StatusOK
func (d *DictionaryHandlers) FlushDictionary(w http.ResponseWriter, r *http.Request) {
	logRequests(r, logFile)
	if r.Method == "DELETE" {
		_, err := os.Stat("storage/dictionary.json")
		if err != nil {
			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "No data found")
			}
		} else {
			file, fileError := os.OpenFile("storage/dictionary.json", os.O_TRUNC, 0666)
			if fileError != nil {
				log.Fatal(fileError)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				d.Data = map[string]string{}
				file.Truncate(0)
				w.WriteHeader(http.StatusOK)
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//NewDictionaryHandlers creates a new struct in DictionaryHandlers type and returns address
func NewDictionaryHandlers() *DictionaryHandlers {
	return &DictionaryHandlers{
		Data: map[string]string{},
	}
}

//logRequest logs http requests to file given as logFile
func logRequests(r *http.Request, logFile string) {
	logMessage := fmt.Sprint(time.Now()) + "\t Request: " + r.Method + fmt.Sprint(r.URL) + "\n"
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, writeError := file.WriteString(logMessage)
	if writeError != nil {
		log.Fatal(writeError)
	}
}
