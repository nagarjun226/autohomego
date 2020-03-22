package main

import (
	"fmt"
	"io/ioutil"

	"github.com/nagarjun226/configmgr/domain"
)

func main() {
	c := domain.Config{}

	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	c.SetFromJSONParsed(configBytes)

	fmt.Println(c.GetConfig("registryapi"))
}
