package acme

import (
	"encoding/json"
	"os"
)

type AcmeContainer struct {
	Acmetls struct {
		Certificates []Certificate `json:"Certificates"`
	} `json:"acmetls"`
}

type Certificate struct {
	Domain struct {
		Main string `json:"main"`
	} `json:"domain"`
	Certificate string `json:"certificate"`
	Key         string `json:"key"`
	Store       string `json:"Store"`
}

func ReadAcmeFile(path string) (AcmeContainer, error) {
	var container AcmeContainer
	if content, err := os.ReadFile(path); err != nil {
		return container, err
	} else if err := json.Unmarshal(content, &container); err != nil {
		return container, err
	} else {
		return container, nil
	}
}
