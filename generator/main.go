package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type StubDetails struct {
	Name, FileName, Destination string
	Values                      map[string]interface{}
}

func (s *StubDetails) CreateStub() bool {
	contentsBuff, err := ioutil.ReadFile(s.Name)
	if err != nil {
		log.Fatalf("Unable to read file: %s", s.Name)
	}

	//f, err := os.OpenFile(s.Destination+s.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755) for append, if data exist before
	f, err := os.OpenFile(s.Destination+s.FileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Unable to open file: %s", s.FileName)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Unable to close file: %s", s.FileName)
		}
	}(f)

	temp, err := template.New(s.FileName).Parse(string(contentsBuff))
	if err != nil {
		log.Fatalf("Unable to parse template: %s", s.Name)
	}
	err = temp.Execute(f, s.Values)
	if err != nil {
		return false
	}
	return err == nil
}
func main() {
	var modelName = "jadwal"
	stubModel := StubDetails{
		Name:        "./stubs/model.go.stub",
		FileName:    modelName + ".go",
		Destination: "../models/",
		Values: map[string]interface{}{
			"Model": strings.Title(modelName),
			"Fields": []map[string]string{
				{"Name": "ID", "Type": "int", "Json": "id"},
				{"Name": "Name", "Type": "string", "Json": "name"},
				{"Name": "Day", "Type": "string", "Json": "day"},
			},
		},
	}
	stubHandler := StubDetails{
		Name:        "./stubs/handler.go.stub",
		FileName:    modelName + "Handler.go",
		Destination: "../handlers/",
		Values: map[string]interface{}{
			"Model": strings.Title(modelName),
			"Data":  modelName,
		},
	}
	stubRepository := StubDetails{
		Name:        "./stubs/repository.go.stub",
		FileName:    modelName + "Repository.go",
		Destination: "../repositories/",
		Values: map[string]interface{}{
			"Model": strings.Title(modelName),
			"Data":  modelName,
		},
	}
	statusMode, statusHandler, statusRepository := stubModel.CreateStub(), stubHandler.CreateStub(), stubRepository.CreateStub()
	if statusMode && statusHandler && statusRepository {
		fmt.Println("Success Create Model, Handler, Repository for " + modelName + " ")
	}
}
