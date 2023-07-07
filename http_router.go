package main

import (
	"encoding/json"
	"net/http"
)

type router struct {
	*http.ServeMux
}

func NewHttpRouter() *router {
	return &router{http.NewServeMux()}
}

type ResponseGetArticles struct {
	Articles []Article `json:"articles"`
}

func (r *router) InitXakepRoute(parser IXakepParse) {
	r.HandleFunc("/getArticles", func(w http.ResponseWriter, r *http.Request) {
		firstPage, err := parser.GetFirstPage()
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		response := ResponseGetArticles{firstPage}

		json, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Write([]byte(json))
	})
}
