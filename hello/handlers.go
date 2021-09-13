package hello

import (
	"encoding/json"
	"net/http"
)

type Person struct{
	Name string `json:"PersonSuperName"`
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		name = "unknown"
	}

	person := Person{
		Name: name,
	}

	json.NewEncoder(writer).Encode(person)
}
