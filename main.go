package main

import (
	"./models"
	"html/template"
	"net/http"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

var posts map[string]*models.Post
var counter int

func indexHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	fmt.Println(counter)

	t.ExecuteTemplate(w, "index", posts)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	id := r.FormValue("id")
	post, found := posts[id]

	if !found {
		http.NotFound(w, r)
	}

	t.ExecuteTemplate(w, "write", post)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
	}
	delete(posts, id)
	http.Redirect(w, r, "/", 302)
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
	title := r.FormValue("title")
	content := r.FormValue("content")
	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		id = models.GenerateId()
		post := models.NewPost(id, title, content)
		posts[post.Id] = post
	}
	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("We listen : 3000! ")
	posts = make(map[string]*models.Post, 0)

	m := martini.Classic()
	counter = 0

	m.Use(func(r *http.Request) {
		if r.URL.Path == "/write" {
			counter++
		}
	})
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs: []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Delims:     render.Delims{"{[{", "}]}"}, // Sets delimiters to the specified strings.
		Charset:    "UTF-8",                     // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                        // Output human readable JSON
	}))
	StaticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", StaticOptions))

	m.Get("/", indexHandler)
	m.Get("/write", writeHandler)
	m.Get("/edit", editHandler)
	m.Get("/delete", deleteHandler)
	m.Post("/SavePost", savePostHandler)

	m.Run()
}
