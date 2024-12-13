package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jp-vergine/go-tutorial/user"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/signin", signinHandler).Methods("POST")
	router.HandleFunc("/auth", authHandler).Methods("GET")

	// Wrap the router with middleware
	wrappedRouter := middlewareHandler(router)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", wrappedRouter)
}

func home(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello my friend %s!", name)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	service := user.Service{}
	response := service.SignupService("john_doe", "john@example.com", "securepassword")
	fmt.Fprint(w, response)
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	service := user.Service{}
	response := service.SigninService("john_doe", "securepassword")
	fmt.Fprint(w, response)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Auth endpoint reached!")
}

func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New Request !", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		// apply some logic after call ...
	})
}
