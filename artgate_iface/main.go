package main

import (
"github.com/julienschmidt/httprouter"
"log"
"net/http"
)

func StubForNotFound(w http.ResponseWriter, r *http.Request) {

}

func main() {
log.Fatal(http.ListenAndServe(":3000", NewApp()))
}

func NewApp() Middleware {

	router := httprouter.New()
	router.Handle("GET", "/", HandleHome)
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))
	router.Handle("GET", "/login", HandleSessionNew)
	router.Handle("POST", "/login", HandleSessionCreate)
	router.Handle("GET", "/user/:userID", HandleUserShow)
	router.NotFound = http.HandlerFunc(StubForNotFound)

	secureRouter := httprouter.New()
	secureRouter.Handle("GET",  "/sign-out", HandleSessionDestroy)
	secureRouter.Handle("GET",  "/account",  HandleUserEdit)
	secureRouter.Handle("POST", "/account",  HandleUserUpdate)

	secureRouter.Handle("GET",  "/setup-ip", HandleSetupEthEdit)
	secureRouter.Handle("POST", "/setup-ip", HandleSetupEthUpdate)

	secureRouter.Handle("GET",  "/setup-artnet", HandleSetupArtnetEdit)
	secureRouter.Handle("POST", "/setup-artnet", HandleSetupArtnetUpdate)

	secureRouter.Handle("GET",  "/setup-artnet-out", HandleSetupArtnetOutEdit)
	secureRouter.Handle("POST", "/setup-artnet-out", HandleSetupArtnetOutUpdate)

	middleware := Middleware{}
	middleware.Add(router)
	middleware.Add(http.HandlerFunc(RequireLogin))
	middleware.Add(secureRouter)

	return middleware
}