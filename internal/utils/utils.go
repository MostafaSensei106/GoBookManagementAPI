package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, x); err != nil {
		return err
	}
	return nil
}

func MapToJSON(m map[string]string) string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}
