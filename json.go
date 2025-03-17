package log

import (
	"encoding/json"
)

func jsonPrettyPrint(data any) string {
	ret, _ := json.MarshalIndent(data, "", "  ")
	return string(ret)
}
