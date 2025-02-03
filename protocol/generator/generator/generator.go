package generator

import (
	"encoding/json"
	"fmt"
	"os"
)

// структура для парсинга протокола
type Protocol struct {
	Commands map[string]interface{} `json:"commands"`
	Structs  map[string]interface{} `json:"structs"`
}

type Command struct {
	CommandCode string                 `json:"command_code"`
	
}

func GenerateProtocol(filePath string) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening protocol file:", err)
		return
	}
	defer jsonFile.Close()

	var protocol map[string]interface{}
	err = json.NewDecoder(jsonFile).Decode(&protocol)
	if err != nil {
		fmt.Println("Error decoding protocol file:", err)
		return
	}

	fmt.Println(protocol)
}
