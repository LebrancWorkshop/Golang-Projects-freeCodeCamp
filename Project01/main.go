package main 

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return 
	}

	fmt.Fprintf(w, "POST Success.")

	name := r.FormValue("name")
	fmt.Fprintf(w, "Name: %s\n", name) 
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server is listening on port 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}