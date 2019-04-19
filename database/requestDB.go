package database

import (
	"encoding/json"
	"fmt"

	"OlympApi/models"
)

func WriteRequestToDB(request *models.Request) {
	requestBytes, _ := json.Marshal(*request)

	fmt.Println(string(requestBytes))
}
