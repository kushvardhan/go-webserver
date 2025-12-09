package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/form" {

	}
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not Found", http.StatusNotFound);
		return;
	}
	if r.Method != "GET" {
		http.Error(w, "Method not found", http.StatusNotFound);
		return;
	}
	fmt.Fprintf(w, "Hello!", http.StatusOK);

}

func main(){
	fmt.Println("Hello from Server");
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer);
	http.HandleFunc("/form", formHandler);
	http.HandleFunc("/hello", helloHandler);

	fmt.Println("Starting server at PORT: 8080.");

	if err:= http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err);
	}
}

