package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJSON(language string, data map[string]interface{}) map[string]interface{} {
	JSON, err := ioutil.ReadFile("app/translations/" + language + ".json")
	if err != nil {
		fmt.Println(err, JSON)
	}
	err = json.Unmarshal(JSON, &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
