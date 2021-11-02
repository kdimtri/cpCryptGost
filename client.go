package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func request(data string, addr string) {
	//	log.Println(data)
	url := fmt.Sprintf("http://%s/signature", addr)
	inPut := map[string]string{"data": data}
	reqData, err := json.Marshal(inPut)
	if err != nil {
		log.Fatalf("Failed marshalling input data: %s", err)
	}
	//	log.Println(string(reqData))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqData))
	if err != nil {
		log.Printf("Error making new request: %s", err)
		return
	}
	req.Header.Set("X-Custom-Header", "data")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return
	}
	outPut := &OutPut{}
	//	log.Printf("outPut: %v\n",outPut)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if err := json.NewDecoder(bytes.NewBuffer(body)).Decode(outPut); err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	//	log.Printf("%s\n", body)
	ok, err := verify(data, outPut)
	if err != nil {
		log.Fatalf("Failed verifying response: %s", err)
	}
	fmt.Printf("Output: \n%s\nVerified:\n%v\n", body, ok)
	return
}
