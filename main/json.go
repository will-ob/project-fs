package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

type ProjectStore struct {
	Transport http.RoundTripper
}

type ProjectJson struct {
	Id string
}

type ProjectJsonCollection struct {
	Json []ProjectJson
}

func check(err error) {
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		log.Fatal(err)
	}
}

func (me *ProjectStore) getHttp(addr string, header http.Header) (r *http.Response) {
	var err error
	var resp *http.Response
	client := &http.Client{Transport: me.Transport}
	req, err := http.NewRequest("GET", addr, nil)
	check(err)
	req.Header = header
	req.Header.Add("X-API-Key", os.Getenv("PROJECT_API_TOKEN"))
	resp, errr := client.Do(req)
	check(errr)
	return resp
}

func (me *ProjectStore) GetJsonIndex() *ProjectJsonCollection {
	log.Println("Get " + os.Getenv("PROJECT_API_URL") + "/v0.1/projects")
	pj := ProjectJsonCollection{Json: []ProjectJson{}}
	resp := me.getHttp(os.Getenv("PROJECT_API_URL")+"/v0.1/projects", map[string][]string{})
	// Need to check for error response, if not error, json decode.
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		log.Println("GET /v0.1/projects => 200")
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&pj.Json)
		check(err)
		return &pj
	}
	log.Println("Error: unexpected response")
	log.Println("GET /v0.1/projects => " + resp.Status)
	log.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	log.Println("\n" + string(body))
	check(err)
	emptyCollection := ProjectJsonCollection{Json: []ProjectJson{}}
	return &emptyCollection
}

func (me *ProjectStore) GetMarkdown(id string) ([]byte, error) {
	resp := me.getHttp(os.Getenv("PROJECT_API_URL")+"/v0.1/projects/"+id+"/content", map[string][]string{"Accept": {"text/markdown"}})
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read markdown body")
	}
	return []byte(body), nil
}
