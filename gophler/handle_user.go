package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type User struct {
	ID             string
	Username       string
	Email          string
	HashedPassword string
}

func HandleUserNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	RenderTemplate(w, r, "users/new", nil)
}

func HandleUserCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Process of creating users
	user, err := NewUser(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"))

	if err != nil {
		if IsValidationError(err) {
			RenderTemplate(w, r, "users/new", map[string]interface{}{
				"Error": err.Error(),
				"User":  user,
			})
			return
		}
		panic(err)
	}

	err = globalUserStore.Save(user)
	if err != nil {
		panic(err)
	}

	session := NewSession(w)
	session.UserID = user.ID

	http.Redirect(w, r, "/?flash=User+created", http.StatusFound)
}
