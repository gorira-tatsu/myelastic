package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var datas []QueryData

type QueryData struct {
	Text string `json:"text"`
}

func entry(writer http.ResponseWriter, request *http.Request) {
	var data QueryData

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if data.Text == "" {
		http.Error(writer, "Text field is empty", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received: %+v\n", data)
	datas = append(datas, data)

	writer.Header().Set("Content-Type", "application/json")
}

func search(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Error reading request body", http.StatusInternalServerError)
		return
	}
	search_word := string(body)

	for _, data := range datas {
		if search_word == data.Text {
			fmt.Fprintf(writer, "yes")
		}
	}

	writer.Header().Set("Content-Type", "text/plain")
}

func show(writer http.ResponseWriter, request *http.Request) {
	if datas != nil {
		fmt.Fprintln(writer, datas)
	} else {
		fmt.Fprintf(writer, "none")
	}
}

func main() {
	http.HandleFunc("/entry", entry)
	http.HandleFunc("/show", show)
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
}
