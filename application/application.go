//Application package includes loadJSON saveJSON and StartApplication fucntions
package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/karagozber/DictionaryApi/handlers"
)

//This is the dictionary file that includes all key-value pairs
const dictionary string = "storage/dictionary.json"

//Starts application
func StartApplication() {

	dictionaryHandler := handlers.NewDictionaryHandlers()
	loadJSON(dictionaryHandler)
	go saveJSON(1, dictionaryHandler)
	http.HandleFunc("/api/get", dictionaryHandler.GetValue)
	http.HandleFunc("/api/set", dictionaryHandler.SetValue)
	http.HandleFunc("/api/flush", dictionaryHandler.FlushDictionary)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}

//loadJson loads data from dictionary.json file if exists
func loadJSON(h *handlers.DictionaryHandlers) {

	fileInfo, err := os.Stat(dictionary)
	if !os.IsNotExist(err) {
		if fileInfo.Size() != 0 {

			file, err := os.OpenFile(dictionary, os.O_CREATE|os.O_RDONLY, 0666)
			defer file.Close()
			if err != nil {

				fmt.Println(err)

			}

			tempMap := map[string]string{}
			byteValue, _ := ioutil.ReadFile(dictionary)
			json.Unmarshal(byteValue, &tempMap)

			for k, v := range tempMap {
				h.Data[k] = v
			}

			fmt.Println("File loaded to map")

		}
	}

}

//saveJSON saves data to dictionary.json file in every n minute
func saveJSON(n time.Duration, h *handlers.DictionaryHandlers) {

	for range time.Tick(n * time.Minute) {

		data, _ := json.Marshal(h.Data)

		file, err := os.OpenFile(dictionary, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		file.Truncate(0)

		_, writeError := file.Write(data)
		if writeError != nil {
			log.Fatal(writeError)
		} else {
			fmt.Println("Dictionary map saved to dictionary.json file", string(data))
		}
	}

}
