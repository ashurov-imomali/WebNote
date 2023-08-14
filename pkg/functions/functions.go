package functions

import (
	"encoding/json"
	"errors"
	"log"
	"main/pkg/models"
	"os"
	"strconv"
)

func SliceNotes() []models.Note {
	Notes := make([]models.Note, 0)
	if ReadFromFile() != nil {
		err := json.Unmarshal(ReadFromFile(), &Notes)
		if err != nil {
			return nil
		}
		return Notes
	}
	return Notes
}

func ReadFromFile() []byte {
	bytes, err := os.ReadFile("notes.json")
	if err != nil {
		log.Println(err)
		return nil
	}
	return bytes
}

func WriteToFile(bytes []byte) error {
	err := os.WriteFile("notes.json", bytes, 0777)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromSlice(Notes *[]models.Note, strId string) error {
	id, err := strconv.Atoi(strId)
	if err != nil {
		return err
	}

	CheckID, index := CheckId(id, *Notes)

	if CheckID == true {
		*Notes = append((*Notes)[:index], (*Notes)[index+1:]...)
		return nil
	}
	return errors.New("BadRequest")
}

func CheckId(id int, Notes []models.Note) (bool, int) {
	for noteIndex, note := range Notes {
		if note.Id == id {
			return true, noteIndex
		}
	}
	return false, 0
}

func UpdateSlice(Notes *[]models.Note, strId string, bytes []byte) error {
	id, err := strconv.Atoi(strId)
	if err != nil {
		return err
	}
	CheckID, index := CheckId(id, *Notes)

	if CheckID == true {
		(*Notes)[index].Content = string(bytes)
		return nil
	}

	return errors.New("BadRequest")
}

func ReadFromSlice(strId string, Notes []models.Note) ([]byte, error) {
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, err
	}
	CheckID, index := CheckId(id, Notes)
	if CheckID == true {
		return []byte(Notes[index].Content), nil
	}
	return nil, errors.New("BadRequest")
}
