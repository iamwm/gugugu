package main

import (
	"fmt"
	"github.com/iamwm/gugugu/blueprints/songs"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("GuGuGu...")
	router := httprouter.New()
	songs.ConstructBlueprint(router)
	log.Fatal(http.ListenAndServe(":8848", router))
}
