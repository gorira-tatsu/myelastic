package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var datas []one_data

type one_data struct {
	date time.Time
	text string
}

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
	datas = append(datas, one_data{time.Now(), data.Text})

	writer.Header().Set("Content-Type", "application/json")
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
	//http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
}
