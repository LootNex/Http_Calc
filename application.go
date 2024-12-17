package application

import "net/http"

func RunServer() error {

	http.HandleFunc("/calc", CalculatorHandler())

	return http.ListenAndServe(":8080", nil)
}