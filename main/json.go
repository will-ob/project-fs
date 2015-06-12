package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ProjectStore struct {
}

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

func getHttp(addr string) (r *http.Response) {
	var err error
	var resp *http.Response
	if os.Getenv("UNSAFE_TLS") == "true" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err = client.Get(addr)
	} else {
		resp, err = http.Get(addr)
	}
	check(err)
	return resp
}

func (me *ProjectStore) GetJsonIndex() *ProjectJsonCollection {
	pj := ProjectJsonCollection{Json: []ProjectJson{}}
	resp := getHttp(os.Getenv("PROJECT_API_URL") + "/v1/projects/")
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&pj.Json)
	check(err)
	return &pj
}

func (me *ProjectStore) GetMarkdown(id string) ([]byte, error) {
	resp := getHttp(os.Getenv("PROJECT_API_URL") + "/v1/projects/" + id)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read markdown body")
	}
	return []byte(body), nil

}
