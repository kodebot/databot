package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	res, _ := ioutil.ReadFile("./home.html")
	fmt.Fprint(w, string(res))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9022", nil))
}
