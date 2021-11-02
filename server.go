package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
)

func runHttpServer(addr string) {
	http.HandleFunc("/signature", signHandler)
	log.Printf("Serving on %s\n", addr)
	go func() {
		log.Fatal(http.ListenAndServe(addr, nil))
	}()
	runtime.Goexit()
}

func signHandler(w http.ResponseWriter, r *http.Request) {
	//	log.Println(r)
	if r.URL.Path != "/signature" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		log.Printf("Error:\n Usupported request method: %s\n. Only 'POST' requests are supported\n", r.Method)
		return
	}
	inPut := &InPut{}
	err := json.NewDecoder(r.Body).Decode(inPut)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//	log.Printf("InPut: %v\n",inPut)
	outPut, err := sign(inPut)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error signing a request data: %s\n", err)
		return
	}
	//	log.Printf("OutPut: %v\n",outPut)
	res, err := json.Marshal(outPut)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshailling output data: %s\n", err)
		return
	}
	//	log.Println(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Printf("Error writing response data: %s\n", err)
	}
	return
}
