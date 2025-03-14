package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed static/index.html
var Index string

//go:embed static/favicon.ico
var FavIcon string

//go:embed static/rickroll.gif
var RickRoll []byte

var Seed = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", index)
	router.Get("/index", index)
	router.Get("/fact", fact)
	router.Get("/favicon.ico", favicon)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(Index))
	return
}

func fact(w http.ResponseWriter, r *http.Request) {
	blpString := r.URL.Query().Get("blp")
	blp, err := strconv.Atoi(blpString)
	if err != nil {
		rick(w)
		return
	}

	randomIndex := Seed.Intn(len(Facts)+5) + blp
	if randomIndex < len(Facts) && randomIndex >= 0 {
		fact := Facts[randomIndex]
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fact))
	} else {
		rick(w)
	}

	return
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(FavIcon))
	return
}

func rick(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "image/gif")
	w.Write(RickRoll)
	return
}
