package main

import (
	"./models"
	"net/http"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/russross/blackfriday"
)

var posts map[string]*models.Post
var counter int

func indexHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	rnd.HTML(200, "index", posts)
}
func writeHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	rnd.HTML(200, "write", nil)
}
func editHandler(rnd render.Render, r *http.Request, params martini.Params) {
	id := params["id"]
	post, found := posts[id]

	if !found {
		rnd.Redirect("/")
		return
	}
	rnd.HTML(200, "write", post)

}

func deleteHandler(rnd render.Render, r *http.Request, params martini.Params) {
	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}
	delete(posts, id)
	rnd.Redirect("/")
}

func savePostHandler(rnd render.Render, r *http.Request) {
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
	rnd.Redirect("/")
}
func getHtmlHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	htmlBytes := blackfriday.MarkdownBasic([]byte(md))

	rnd.JSON(200, map[string]interface{}{"html": string(htmlBytes)})
}

func main() {
	fmt.Println("We listen : 3000! ")
	posts = make(map[string]*models.Post, 0)

	m := martini.Classic()
	counter = 0

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Charset:    "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                       // Output human readable JSON
	}))
	StaticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", StaticOptions))

	m.Get("/", indexHandler)
	m.Get("/write", writeHandler)
	m.Get("/edit/:id", editHandler)
	m.Get("/delete/:id", deleteHandler)
	m.Post("/SavePost", savePostHandler)
	m.Post("/getHtml", getHtmlHandler)
	m.Run()
}
