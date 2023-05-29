package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Addr     string `json:"addr"`
	Protocol string `json:"protocol"`
	Type     string `json:"type"`
}

func main() {
	data, err := os.ReadFile("data/config.json")
	if err != nil {
		log.Fatal(err)
	}

	var configs []Config

	if err = json.Unmarshal(data, &configs); err != nil {
		log.Fatal(err)
	}

	configs = append(configs, Config{Addr: "begingo.com", Protocol: "https", Type: "website"})

	ndata, err := json.Marshal(configs)
	if err != nil {
		log.Fatal(err)
	}

	if err = os.WriteFile("data/config.json", ndata, 0755); err != nil {
		log.Fatal(err)
	}
}
