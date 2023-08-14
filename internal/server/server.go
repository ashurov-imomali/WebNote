package server

import (
	"github.com/gorilla/mux"
	"main/configs"
	"main/internal/handlers"
	"net/http"
)

func Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/Create", handlers.CreateNote).Methods(http.MethodPost)
	router.HandleFunc("/Read", handlers.ReadNote).Methods(http.MethodGet)
	router.HandleFunc("/Update", handlers.UpdateNote).Methods(http.MethodPatch)
	router.HandleFunc("/Delete", handlers.DeleteNote).Methods(http.MethodDelete)

	NewConfigs, err := configs.InitConfig()
	if err != nil {
		return err
	}
	address := NewConfigs.Host + NewConfigs.Port

	err = http.ListenAndServe(address, router)
	if err != nil {
		return err
	}
	return nil
}
