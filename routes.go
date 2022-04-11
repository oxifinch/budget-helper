package main

import (
	"fmt"
	"net/http"
)

// Routes are named as "routeDOMAIN"

func routeHome(writer http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(writer, "Welcome to Budget Helper!\n")
	if err != nil {
		panic(err)
	}
}

// TODO: This method of registering routes works, but will probably not be nice
// to work with when there are A LOT of routes which then need to be manually
// added/removed from this function. Is there a better way?
func initRouter() http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", routeHome)

	fmt.Printf("Registered routes...\n")
	return *router
}
