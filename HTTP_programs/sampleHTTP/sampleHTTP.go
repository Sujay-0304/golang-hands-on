package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "inside home %s ", r.URL)
	
}

func aboutus(w http.ResponseWriter, r *http.Request ){
	fmt.Fprint(w, "inside about us")
	fmt.Fprintf(w, "%s", r.URL)
}

func projects(res http.ResponseWriter,req *http.Request){
	fmt.Fprint(res, "inside projects" )
}
func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/aboutus", aboutus)
	http.HandleFunc("/projects", projects) 
	fmt.Println("Server running at :8080")
    http.ListenAndServe(":8080", nil) 
}