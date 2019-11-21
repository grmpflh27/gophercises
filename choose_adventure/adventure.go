package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

func loadStory(storyPath string) map[string]interface{} {
	fp, err := os.Open(storyPath)
	if err != nil {
		log.Fatal("Could not open story json")
	}

	defer fp.Close()
	byteValue, _ := ioutil.ReadAll(fp)

	// TODO: proper parsing

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	return result
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CHOOSE YOUR OWN ADVENTURE!")
	fmt.Println("Endpoint Hit: homePage")
}

func setupRouter(chapters map[string]interface{}) *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)

	tmpl := template.Must(template.ParseFiles("chapter_template.html"))
	for k, v := range chapters {
		curChapterName := k
		curChapter := v
		myRouter.HandleFunc("/"+curChapterName, func(w http.ResponseWriter, r *http.Request) {
			tmpl.Execute(w, curChapter)
		})
	}

	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, chapters["intro"])
	})

	return myRouter
}

func main() {
	fmt.Println("Starting this adventure")
	storyJSONPath := "gopher.json"

	// loading story
	chapters := loadStory(storyJSONPath)

	// define routes and start webserver
	router := setupRouter(chapters)
	http.ListenAndServe(":10000", router)
}

type Chapter struct {
	Title   string           `json:"title"`
	Story   []string         `json:"story"`
	Options []ChapterOptions `json:"options"`
}

type ChapterOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
