package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func decodeJSONMap(reader io.Reader) (map[string]interface{}, error) {
	var payload interface{}
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()

	if err := decoder.Decode(&payload); err != nil {
		return nil, err
	}

	switch v := payload.(type) {
	case map[string]interface{}:
		return v, nil
	case nil:
		return map[string]interface{}{}, nil
	default:
		return map[string]interface{}{"data": v}, nil
	}
}
func ConvertAnyToJSONString(data any) (string, error) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return string(jsonBytes), nil
}

func ReadJSONBody(r *http.Request) map[string]interface{} {
	if r.Body == nil {
		return nil
	}
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil
	}
	return data
}

func ReadJSONResponse(r *http.Response) map[string]interface{} {
	if r == nil || r.Body == nil {
		return nil
	}
	defer r.Body.Close()

	data, err := decodeJSONMap(r.Body)
	if err != nil {
		return nil
	}
	return data
}
