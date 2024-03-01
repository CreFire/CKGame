package main

import (
	"CKGame/src/model"
	log "github.com/sirupsen/logrus"
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
	log.Info("start liten:8080")
	StartPprofLiten()

}
