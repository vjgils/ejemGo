//web Server

package main

import (
"log"
"net/http"
)

func main() {
    http.HandleFunc("/", handler) //cada petición manda llamar a un hadler.
    log.Fatal(http.ListenAndServe("localhost:8002",nil)) //especifica que escuche en el puerto 8000.
}

func handler(w http.ResponseWriter, r *http.Request){
    //2 maneras de hacer Redirect
    http.Redirect(w, r, "http://www.google.com", 301)
    // http.Redirect(w,r,"http://www.google.com",http.StatusMovedPermanently)
}