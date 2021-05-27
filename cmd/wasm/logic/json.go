package logic

import (
	"encoding/json"
	"fmt"
	"time"
)

func PrettyJSON(input string) (string, error) {
	start := time.Now()
	defer func() {
		fmt.Printf("\n\nDone in %v\n", time.Since(start))
	}()
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}

	pretty, err := json.MarshalIndent(raw, "", " ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}
