package main

import (
	"example/Store-Api/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
