package server

import (
	"log"
	"net/http"
)

func Run(configPath string) error {
	config, err := LoadConfig(configPath)
	if err != nil {
		return err
	}

	logger := log.Default()

	router := Router{config: config}
	router.logger = logger

	http.HandleFunc("/", router.Handle)

	logger.Println("server listen :8080 ...")
	return http.ListenAndServe(":8080", nil)
}
