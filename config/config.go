package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//SetEnvironment set environment variables from a env.json file in runtime, if the file exists
func SetEnvironment() {
	file, err := ioutil.ReadFile("./config/env.json")
	if err == nil {

		jsonMap := make(map[string]interface{})
		json.Unmarshal(file, &jsonMap)

		env := "development"
		if jsonMap[env] == nil {
			return
		}
		database := jsonMap[env].(map[string]interface{})

		for key, value := range database {

			switch value.(type) {
			case string:
				os.Setenv(key, value.(string))
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				os.Setenv(key, fmt.Sprintf("%d", value.(int)))
			case float32, float64:
				val := fmt.Sprintf("%.2f", value.(float64))
				strings := strings.Split(val, ".")
				if strings[1] != "00" {
					os.Setenv(key, val)
				} else {
					os.Setenv(key, strings[0])
				}
			case bool:
				os.Setenv(key, fmt.Sprintf("%v", value.(bool)))
			}
		}
	}
}
