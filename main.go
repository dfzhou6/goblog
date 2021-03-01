package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to goblog!</h1>")
}

func AboutFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>This is About!</h1>")
}

func NotFoundFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page not found!</h1>")
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>CreateArticle succeed!</h1>")
}

func ShowArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Show Article "+params["id"]+"!</h1>")
}

func ShowArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Show all articles!</h1>")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeFunc).Methods("Get").Name("home")
	router.HandleFunc("/about", AboutFunc).Methods("GET").Name("about")
	router.HandleFunc("/404", NotFoundFunc).Methods("GET").Name("404")
	router.HandleFunc("/article", CreateArticle).Methods("POST").Name("article.store")
	router.HandleFunc("/articles", ShowArticles).Methods("GET").Name("article.list")
	router.HandleFunc("/article/{id:[0-9]+}", ShowArticle).Methods("GET").Name("article.show")

	fmt.Println(router.Get("home").URL())
	fmt.Println(router.Get("about").URL())
	fmt.Println(router.Get("404").URL())
	fmt.Println(router.Get("article.show").URL("id", "123"))
	http.ListenAndServe(":3000", router)
}
