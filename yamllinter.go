package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	generic string = "g"
	schema  string = "s"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}

}

func nodeToMap(node interface{}) map[string]interface{} {
	m, ok := node.(map[string]interface{})
	if !ok {
		panic(fmt.Sprintf("%v is not of type map", node))
	}
	return m
}

func nodeToList(node interface{}) []interface{} {
	m, ok := node.([]interface{})
	if !ok {
		panic(fmt.Sprintf("%v is not of type list", node))
	}
	return m
}

type sampleSchema struct {
	Service    service    `yaml:"service"`
	Credential credential `yaml:"credential"`
}

type service struct {
	Name string `yaml:"name"`
}

type credential struct {
	ClientId       string         `yaml:"clientId"`
	ClientSecret   string         `yaml:"clientSecret"`
	TokenURL       string         `yaml:"tokenURL"`
	EndpointParams endPointParams `yaml:"endpointParams"`
	AuthStyle      string         `yaml:"authStyle"`
}

type endPointParams struct {
	GrantType []string `yaml:"grant_type"`
}

func main() {

	argsWithProg := os.Args
	filename := argsWithProg[1]
	option := argsWithProg[2]
	fmt.Println(filename)
	if len(filename) > 0 {
		src, openError := ioutil.ReadFile(filename)
		CheckError(openError)
		switch option {
		case generic:

			//f, err := os.Open(filename)
			//CheckError(err)
			var conf map[string]interface{}
			err := yaml.Unmarshal([]byte(src), &conf)
			CheckError(err)
			//fmt.Printf("Value : %v\n", conf)
			//fmt.Println(conf)

			for key := range conf {
				fmt.Printf("%s : %v \n", key, conf[key])
				//fmt.Println(key)
			}

			/*value := nodeToList(nodeToMap(conf)["bar"])[0]
			fmt.Println("-------------:::-------------------")
			fmt.Printf("Value: %#v\n", value)
			*/
		case schema:
			var schemaBackedYaml sampleSchema
			err := yaml.Unmarshal([]byte(src), &schemaBackedYaml)
			CheckError(err)
			fmt.Println(schemaBackedYaml)

		}
	}

}
