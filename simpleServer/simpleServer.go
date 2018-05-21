package simpleServer

import (
	"net/http"
	"strconv"
	"fmt"
	"os"
)

type RouteData struct {

	Port int
	Route string
}


func Server(routeData RouteData) {

	http.HandleFunc(routeData.Route , rootHandler)
	err := http.ListenAndServe(":" + strconv.Itoa(routeData.Port), nil)

	if err != nil {
		fmt.Println("Server error :" , err)
		os.Exit(10)
	}

}


func rootHandler(w http.ResponseWriter , r * http.Request) {

	fmt.Fprintf(w , "Serving : %s\n" , r.URL.Path)
	fmt.Printf("Served : %s\n" , r.Host)
}
