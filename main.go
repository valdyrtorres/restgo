package main

import (
	"log"
	"net/http"

	"./common"
	"./routers"
	"github.com/urfave/negroni" // pequeno middleware web em Go
)

func main() {
	common.StartUp()
	r := routers.InitRouters()

	m := negroni.Classic()
	m.UseHandler(r)

	log.Println("Listening 8721...")
	http.ListenAndServe(":8721", m)
}
