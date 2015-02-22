package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type ProjectJson struct {
	Id string
}

type ProjectJsonCollection struct {
	Json []ProjectJson
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetProjectJsonIndex() *ProjectJsonCollection {
	pj := ProjectJsonCollection{Json: []ProjectJson{}}
	resp, err := http.Get("http://localhost:3333/projects/")
	check(err)
	decoder := json.NewDecoder(resp.Body)
	errr := decoder.Decode(&pj.Json)
	check(errr)
	return &pj
}

func GetProjectMarkdown(id string) ([]byte, error) {
	resp, err := http.Get("http://localhost:3333/projects/" + id)
	check(err)
	if err != nil {
		return nil, errors.New("Could not retrieve markdown body")
	}

	defer resp.Body.Close()
	body, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		return nil, errors.New("Could not read markdown body")
	}
	return []byte(body), nil

}
