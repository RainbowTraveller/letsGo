/*
Execute with schema : GoRun % schema.yaml s
Execute generic yaml file : GoRun % schema.yaml g
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
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
	ClientId       int            `yaml:"clientId" validate:"required"`
	ClientSecret   string         `yaml:"clientSecret" validate:"required"`
	TokenURL       string         `yaml:"tokenURL"`
	EndpointParams endPointParams `yaml:"endpointParams"`
	AuthStyle      string         `yaml:"authStyle"`
}

type endPointParams struct {
	GrantType []string `yaml:"grant_type"`
}

func getLineNoAndMsg(errorMsg string) (no string, msg string) {
	from := strings.Index(errorMsg, "line")
	to := strings.LastIndex(errorMsg, ":")
	// fmt.Println(from)
	// fmt.Println(to)
	line := errorMsg[from:to]
	no = strings.Split(line, " ")[1]
	prev := errorMsg[:from]
	// fmt.Println(prev)

	post := errorMsg[to+1:]
	// fmt.Println(post)
	var msgBuilder strings.Builder
	msgBuilder.WriteString(prev)
	msgBuilder.WriteString(post)
	//# fmt.Println(line)
	//# fmt.Println(msgBuilder.String())
	//# fmt.Println(no)
	return no, msgBuilder.String()
}

func main() {

	argsWithProg := os.Args
	filename := argsWithProg[1]
	option := argsWithProg[2]
	//fmt.Println(filename)
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

		case schema:
			var schemaBackedYaml sampleSchema
			//err := yaml.Unmarshal([]byte(src), &schemaBackedYaml)
			err := yaml.UnmarshalStrict([]byte(src), &schemaBackedYaml)
			if err != nil {
				errorMsg := err.Error()
				fmt.Println("")
				fmt.Println("::::::::::::::::: Original Message::::::::::::::")
				fmt.Println(errorMsg)
				num, msg := getLineNoAndMsg(errorMsg)
				fmt.Println("::::::::::::::::: Unmarshalled Data ::::::::::::::")
				fmt.Println()
				fmt.Printf("Line no. : %s\n", num)
				fmt.Printf("mssage : %s\n", msg)
			}
			fmt.Println(schemaBackedYaml)
			// returns nil or ValidationErrors ( []FieldError )
			fmt.Println("VALIDATING THE DATA:----")
			validate := validator.New()
			errv := validate.Struct(schemaBackedYaml)
			if errv != nil {

				// this check is only needed when your code could produce
				// an invalid value for validation such as interface with nil
				// value most including myself do not usually have code like this.
				// if _, ok := err.(*validator.InvalidValidationError); ok {
				// 	fmt.Println("InvalidValidationError : ", err)
				// 	return
				// }

				fmt.Println(errv.Error())
				for _, err := range errv.(validator.ValidationErrors) {

					fmt.Println("Namespace :", err.Namespace())
					fmt.Println("Field :", err.Field())
					fmt.Println("StructNamespace : ", err.StructNamespace())
					fmt.Println("StructField : ", err.StructField())
					fmt.Println("Tag :", err.Tag())
					fmt.Println("ActualTag : ", err.ActualTag())
					fmt.Println("Kind :", err.Kind())
					fmt.Println("Type :", err.Type())
					fmt.Println("Value :", err.Value())
					fmt.Println("Param :", err.Param())
					message := fmt.Sprint(err)
					fmt.Println("Message : ", message)
					fmt.Println()
				}

				// from here you can create your own error messages in whatever language you wish
				return
			}
		}
	}

}
