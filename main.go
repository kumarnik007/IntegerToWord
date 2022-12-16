package main

import (
	"log"
	"net/http"
	"strconv"
)

const PORT_NUMBER int = 8081

func main() {
	addr := ":" + strconv.Itoa(PORT_NUMBER)

	http.Handle("/identity", ApiHandler(Identity))
	http.Handle("/convert", ApiHandler(Convert))

	log.Println("== Welcome to my Web Server ==")
	log.Println("== Server Is Listening On Port", addr, " ==")
	log.Fatal(http.ListenAndServe(addr, nil))
}
