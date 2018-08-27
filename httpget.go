package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func hotRestart() {
	s := basicAuth()
	fmt.Printf(s)
}

func basicAuth() string {
	fmt.Printf("calling..\n")
	var username = ""
	var passwd = ""
	client := &http.Client{}
	req, err := http.NewRequest("GET", "", nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	fmt.Printf("Http response code is %d\n", resp.StatusCode)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}
