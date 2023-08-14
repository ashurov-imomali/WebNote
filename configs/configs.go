package configs

import (
	"encoding/json"
	"log"
	"main/pkg/models"
	"os"
)

func InitConfig() (*models.Config, error) {
	//Чтение из файла
	bytes, err := os.ReadFile("./configs/configs.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//Записиваем в Структуру
	var NewConfigs models.Config
	err = json.Unmarshal(bytes, &NewConfigs)
	if err != nil {
		return nil, err
	}

	//Возврашаем :)
	return &NewConfigs, nil
}
