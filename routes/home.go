package routes

import (
	"github.com/martini-contrib/render"
	"fmt"
	"go-blog/db/documents"
	"gopkg.in/mgo.v2"
	"go-blog/models"
	"go-blog/sessions"
)

func IndexHandler(rnd render.Render, s *sessions.Session, db *mgo.Database) {
	fmt.Println(s.Username)

	postDocuments := []documents.PostDocument{}
	postsCollection := db.C("posts")
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown}
		posts = append(posts, post)
	}

	rnd.HTML(200, "index", posts)
}