package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// NOTE: "next" is a common name for middleware argument
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		/* NOTE: next is something "moving to the next..."
		It might move to next middleware or another part of the file where we return mux */
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/", // NOTE: "/" applys to ENTIRE SITE for a cookie
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
