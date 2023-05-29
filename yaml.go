package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Name string   `yaml:"name"`
	Tags []string `yaml:"tags"`
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

	ndata, err := yaml.Marshal(metadata)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("data/metadata.yaml", ndata, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
