package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

const Path string = "conf.json"

type Configuration struct {
	Token             string
	Group             string
	RequestsPerMinute uint
	MembersLimit      uint
	Prefix            string
	Postfix           string
}

func LoadConfig() (Configuration, error) {
	if !ConfigIsExists() {
		CreateConfig()
		log.Fatal("The configuration file does not exist, a new one has been created, please fill it out")
	}

	file, fErr := os.Open(Path)

	if fErr != nil {
		log.Fatal("File not opened!")
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		log.Fatal("Unable to decode configuration file!")
	}

	return configuration, nil
}

func ConfigIsExists() bool {
	if _, err := os.Stat(Path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func CreateConfig() {
	data := []byte("{\n  \"Token\": \"YOU-TG-TOKEN\",\n  \"Group\": \"YOU-SECRET-GROUP-ID\",\n  \"RequestsPerMinute\": 30,\n  \"MembersLimit\": 1,\n  \"Prefix\": \"\",\n  \"Postfix\": \"\"\n}")
	fo, err := os.Create(Path)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	err = os.WriteFile(Path, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
