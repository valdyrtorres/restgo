package main

import (
	"log"
	"net/http"
	"./common"
	"./routers"
	"github.com/urfave/negroni"
)

func main() {
	common.StartUp()
	r := routers.InitRouters()

	m := negroni.Classic()
	m.UseHandler(r)

	log.Println("Listening 8701...")
	http.ListenAndServe(":8701", m)
}