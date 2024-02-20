package main

import (
	"CKGame/src/model"
	"net/http"
)

func StartPprofLiten() {

	err := http.ListenAndServe(model.Config.Port, nil)
	if err != nil {
		return
	}
}

func main() {
	model.LoadConfig()
	StartPprofLiten()
}
