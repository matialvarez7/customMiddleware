package main

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Execute middleware before request phase!")
		// Pass control back to the handler
		handler.ServeHTTP(w, r)
		fmt.Println("Execute middleware after response phase!")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	fmt.Println("Executing mainhandler...")
	w.Write([]byte("OK"))
}

func main() {
	// HandlerFunc retturns a HTTP Handler
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":3000", nil)
}
