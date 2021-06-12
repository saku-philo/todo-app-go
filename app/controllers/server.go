package controllers

import (
	"net/http"
	"todo_app/config"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	// StripPrefix で /staticを取り除く
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
