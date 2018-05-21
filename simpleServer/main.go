package main

/*
This example is build on server example given in "packt go system programing"
*/


import (
	"fmt"
	"time"
	"net/http"
	"strconv"
)


type RouteData struct {

	Port int
	RouterMap map[string]func(w http.ResponseWriter , r *http.Request)
}


func Server(routeData RouteData) {

	serverMux := http.NewServeMux()

	for route , routeHandler := range routeData.RouterMap {
		serverMux.HandleFunc(route , routeHandler)
	}

	fmt.Println("Servering on localhost:" , routeData.Port)
	http.ListenAndServe(":" + strconv.Itoa(routeData.Port) , serverMux)

}


func about(w http.ResponseWriter , r * http.Request) {
	fmt.Fprintf(w , "This is the about page : %s\n" , r.URL.Path)
	fmt.Printf("Served : %s\n" , r.Host)
}



func cv(w http.ResponseWriter , r * http.Request) {

	fmt.Fprintf(w , "This is the /CV page : %s\n" , r.URL.Path)
	fmt.Printf("Served : %s\n" , r.Host)
}



func timeHandler(w http.ResponseWriter , r * http.Request) {

	currentTime := time.Now().Format(time.RFC1123)
	title := currentTime
	Body := "The current time is :"
	fmt.Fprintf(w , "<h1 align=\"center\"%s</h1><h2 align = \"center\">%s</h2>" , Body , title)

	fmt.Printf("Served : %s for %s\n" , r.URL.Path , r.Host)
}



func home(w http.ResponseWriter , r * http.Request) {
	
	if r.URL.Path == "/" {
		fmt.Fprintf(w , "Welcome to my home page!\n")
	} else {
		fmt.Fprintf(w , "Unknown page : %s from %s\n" , r.URL.Path ,  r.Host)
	}

	fmt.Printf("Served : %s for %s\n" , r.URL.Path , r.Host)
}


func main() {

	// route order matters , last route catches everything.

	routeMap := map[string]func(w http.ResponseWriter , r *http.Request) {
		"/about" : about ,
		"/CV"    : cv,
		"/time"  : timeHandler,
		"/"      : home ,
	}

	routeData := RouteData { Port : 8001 , RouterMap : routeMap}
	Server(routeData)

}
