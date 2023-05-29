package main

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Name string   `yaml:"name" json:"name"`
	Tags []string `yaml:"tags" json:"tags"`
}

func main() {
	data, err := os.ReadFile("data/metadata.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var metadata Metadata

	if err = yaml.Unmarshal(data, &metadata); err != nil {
		log.Fatal(err)
	}

	metadata.Name += "!!!"
	metadata.Tags = append(metadata.Tags, "java")

	ndata, err := json.Marshal(metadata)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("data/convert.json", ndata, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
