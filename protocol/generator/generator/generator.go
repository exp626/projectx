package generator

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"text/template"
)

// структура для парсинга протокола
type Protocol struct {
	Version  string    `json:"version"`
	Protocol string    `json:"protocol_name"`
	Commands []Command `json:"commands"`
}

type Command struct {
	CommandCode int    `json:"command_code"`
	Body        Body   `json:"body"`
	Description string `json:"description"`
}

type Body struct {
	Type        string  `json:"type"`
	Direction   string  `json:"direction"`
	Description string  `json:"description"`
	Name        string  `json:"name"`
	Options     Options `json:"options"`
}

type Options struct {
	Fields []Field `json:"fields"`
}

type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Types struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Options Options `json:"options"`
}

func GenerateProtocol(filePath string, templatePath string, outputPath string) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening protocol file:", err)
		return
	}
	defer jsonFile.Close()

	var protocol Protocol
	err = json.NewDecoder(jsonFile).Decode(&protocol)
	if err != nil {
		fmt.Println("Error decoding protocol file:", err)
		return
	}

	// генерируем фаил с командами
	generateCommands(protocol, templatePath, outputPath)

	// генерируем фаил с типами
	generateTypes(protocol, templatePath, outputPath)
}

func generateCommands(protocol Protocol, templatePath string, outputPath string) {
	// генерируем фаил с командами

	slog.Info("Generating commands", "templatePath", templatePath, "outputPath", outputPath)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	tmpl.Execute(os.Stdout, protocol)
}

func generateTypes(protocol Protocol, templatePath string, outputPath string) {
	// генерируем фаил с типами

	slog.Info("Generating types", "templatePath", templatePath, "outputPath", outputPath)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	tmpl.Execute(os.Stdout, protocol)
}
