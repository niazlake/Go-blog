package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"fmt"
	"go-blog/sessions"
)

func GetLoginHandler(rnd render.Render) {
	rnd.HTML(200, "login", nil)
}

func PostLoginHandler(rnd render.Render, r *http.Request, s *sessions.Session) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Print(username)
	fmt.Print(password)
	s.Username = username

	rnd.Redirect("/")
}
