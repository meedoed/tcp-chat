package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	IP string
}

var (
	path = "./config/config.json"
	Conf Config
)

func init() {
	// чтение файла конфигурации
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	// получение IP-адреса и порта из файла
	err = json.Unmarshal(data, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}
