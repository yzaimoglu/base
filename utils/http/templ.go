package http_utils

import "encoding/json"

func AlpineData(s any) string {
	ss, _ := json.Marshal(s)
	return string(ss)
}
