package coda

import "strings"

func removeFromJsonString(jsonString string) string {
	jsonString = strings.Replace(jsonString, "'", `"`, -1)
	jsonString = strings.Replace(jsonString, "None", "0", -1)
	jsonString = strings.Replace(jsonString, "null", "0", -1)
	jsonString = strings.Replace(jsonString, `{"data":`, "", -1)
	jsonString = jsonString[:len(jsonString)-1]
	return jsonString
}
