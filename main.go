package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"./utils"
)

var conf *utils.Config = new(utils.Config)

func main() {
	conf = utils.LoadConfig("./config.toml1")
	if conf.Debug {
		if b, err := json.Marshal(conf); err == nil {
			var out bytes.Buffer
			if err := json.Indent(&out, b, "", "\t"); err == nil {
				fmt.Println(out.String())
			} else {
				// сюда по идее мы никогда не придем, мы же по сути из json делаем json
				log.Fatal("Что-то пошло не так..")
			}
		} else {
			// сюда по идее мы никогда не придем, мы же по сути из json делаем json
			log.Fatal("Что-то пошло не так..")
		}
	}
	fmt.Println("Ok")
}
