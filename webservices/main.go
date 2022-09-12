package main

import (
	"net/http"
	"webservices/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8001", nil)
}
