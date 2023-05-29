package main

import (
	"log"
	"os"
	"strings"
)

var (
	txt string
)

func main() {
	data, err := os.ReadFile("data/aprenda.txt")
	if err != nil {
		log.Fatal(err)
	}

	arrString := strings.Split(string(data), " ")

	for _, str := range arrString {
		if str == "bacana" {
			txt = strings.Replace(string(data), "bacana", "legal", 1)
		} else {
			txt = strings.Replace(string(data), "legal", "bacana", 1)
		}
	}

	if err := os.WriteFile("data/aprenda.txt", []byte(txt), 0755); err != nil {
		log.Fatal(err)
	}
}
