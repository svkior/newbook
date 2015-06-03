package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func HandleUserImage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//userID := params.ByName("user_id")
	//imageID := params.ByName("image_id")
}

func StubForNotFound(w http.ResponseWriter, r *http.Request) {

}

func main() {
	log.Fatal(http.ListenAndServe(":3000", NewApp()))
}

func NewApp() Middleware {
	router := httprouter.New()
	router.Handle("GET", "/", HandleHome)
	router.Handle("GET", "/register", HandleUserNew)
	router.Handle("POST", "/register", HandleUserCreate)
	router.Handle("GET", "/login", HandleSessionNew)
	router.Handle("POST", "/login", HandleSessionCreate)
	router.Handle("GET", "/image/:imageID", HandleImageShow)
	router.Handle("GET", "/user/:userID", HandleUserShow)
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))
	router.ServeFiles("/im/*filepath", http.Dir("data/images/"))

	router.NotFound = http.HandlerFunc(StubForNotFound)

	secureRouter := httprouter.New()
	secureRouter.Handle("GET", "/sign-out", HandleSessionDestroy)
	secureRouter.Handle("GET", "/account", HandleUserEdit)
	secureRouter.Handle("POST", "/account", HandleUserUpdate)
	secureRouter.Handle("GET", "/images/new", HandleImageNew)
	secureRouter.Handle("POST", "/images/new", HandleImageCreate)

	middleware := Middleware{}
	middleware.Add(router)
	middleware.Add(http.HandlerFunc(RequireLogin))
	middleware.Add(secureRouter)

	return middleware
}
