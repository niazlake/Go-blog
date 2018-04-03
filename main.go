package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	t.ExecuteTemplate(w, "write", nil)
}
func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue(" title")
	content := r.FormValue("content")
}

func main() {
	fmt.Println("We listen : 3000! ")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/savePost", savePostHandler)

	http.ListenAndServe(":3000", nil)
}
