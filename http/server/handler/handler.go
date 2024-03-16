package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"go.uber.org/zap"
)
// format json data
type Posts struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

const (
	dbPath = "" // path to json file
)

// send
func Send(w http.ResponseWriter, r *http.Request) {
	var posts []Posts
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	if r.Method == "POST" {
		title := r.URL.Query().Get("title")
		text := r.URL.Query().Get("text")

		if title == "" || text == "" {
			logger.Warn("Text or title is null")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		db, err := os.Open(dbPath)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		data, err := io.ReadAll(db)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(data, &posts)
		if err != nil {
			panic(err)
		}
		posts = append(posts, Posts{Title: title, Text: text})
		data, err = json.Marshal(posts)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(dbPath, data, 0666)
		if err != nil {
			panic(err)
		}
		logger.Info("Note is added")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

// index page
func Index(w http.ResponseWriter, r *http.Request) {
	var posts []Posts

	fmt.Fprint(w, "Index page\n\n")
	db, err := os.Open(dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	data, err := io.ReadAll(db)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &posts)
	if err != nil {
		panic(err)
	}

	for i := range posts {
		fmt.Fprintf(w, "%s\n%s\n\n", posts[i].Title, posts[i].Text)
	}
}
