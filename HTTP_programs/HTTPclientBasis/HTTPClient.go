package main

import (
    "fmt"
    "io"
    "net/http"
)
func main() {
    localHost := "http://localhost:3000/"
    http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
        fmt.Fprint(w, "inside the home directory at localhost 3000")
    })
    http.HandleFunc("/get", Get)
    http.HandleFunc("/redirect",Redirect) 
    http.HandleFunc("/notfound", NotFound) 
    http.HandleFunc("/servefile", sendFile)
    fmt.Println("server started at:", localHost)
    http.ListenAndServe(":3000", nil) 
}
func Get(w http.ResponseWriter , r *http.Request){
    resp, err := http.Get("https://www.google.com")
    if err != nil {
        fmt.Println("Error fetching data:", err)
        return
    }
    // defer resp.Body.Close() 
    body, _ := io.ReadAll(resp.Body)
    fmt.Fprintln(w, string(body))
}
func sendFile(w http.ResponseWriter , r *http.Request){
    http.ServeFile(w,r,"index.html")
}
func Redirect(w http.ResponseWriter , r *http.Request){
        // fmt.Fprintf(w, "this is the redirect file")
        http.Redirect(w,r,"get", 200)
}
func NotFound(w http.ResponseWriter,  r *http.Request) { 
    http.NotFound(w,r) 
    http.Error(w, "something went wrong", http.StatusExpectationFailed) 
}