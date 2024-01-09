package utils

import (
	"encoding/json"
	"fmt"
)

func EntityToMap(T any) map[string]interface{} {
	result := make(map[string]interface{})

	data, err := json.Marshal(T)
	if err != nil {
		return make(map[string]interface{})
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return make(map[string]interface{})
	}

	return result
}

func MapToEntity(T any, K any) {

	data, err := json.Marshal(T)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(data, &K)

	/*if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, K); err != nil {
		panic(err)
	} */
}

func ToJson(T any) string {
	json, err := json.Marshal(T)

	if err != nil {
		return ""
	}

	return string(json)
}
