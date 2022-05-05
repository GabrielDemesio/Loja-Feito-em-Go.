package main

import (
	"net/http"

	"gabriel/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
