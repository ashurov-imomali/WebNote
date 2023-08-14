package handlers

import (
	"encoding/json"
	"io"
	"log"
	"main/pkg/functions"
	"main/pkg/models"
	"net/http"
	"time"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	Notes := functions.SliceNotes()

	var NewNote models.Note
	NewNote.Id = len(Notes) + 1

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	NewNote.Content = string(bytes)
	NewNote.Date = time.Now()

	Notes = append(Notes, NewNote)

	bytes, err = json.MarshalIndent(Notes, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = functions.WriteToFile(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ReadNote(w http.ResponseWriter, r *http.Request) {
	Notes := functions.SliceNotes()

	strId := r.URL.Query().Get("Id")

	bytes, err := functions.ReadFromSlice(strId, Notes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Write(bytes)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	Notes := functions.SliceNotes()

	strId := r.URL.Query().Get("Id")

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = functions.UpdateSlice(&Notes, strId, bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	bytes, err = json.MarshalIndent(Notes, "", " ")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = functions.WriteToFile(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	Notes := functions.SliceNotes()
	strId := r.URL.Query().Get("Id")
	err := functions.DeleteFromSlice(&Notes, strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	bytes, err := json.MarshalIndent(Notes, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = functions.WriteToFile(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
