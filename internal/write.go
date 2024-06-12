package internal

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(writer http.ResponseWriter, status int, obj any, headers ...map[string]string) error {
	if len(headers) != 0 {
		for _, value := range headers {
			for key, value2 := range value {
				writer.Header().Set(key, value2)
			}
		}
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	parse, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	writer.Write(parse)
	return nil
}

func WriteHTML(writer http.ResponseWriter, status int, obj string, headers ...map[string]string) error {
	if len(headers) != 0 {
		for _, value := range headers {
			for key, value2 := range value {
				writer.Header().Set(key, value2)
			}
		}
	}

	writer.Header().Set("Content-Type", "text/html")
	writer.WriteHeader(status)
	_, err := writer.Write([]byte(obj))
	if err != nil {
		return err
	}

	return nil
}

func WriteTo(writer http.ResponseWriter, status int, obj []byte, headers ...map[string]string) error {
	if len(headers) != 0 {
		for _, value := range headers {
			for key, value2 := range value {
				writer.Header().Set(key, value2)
			}
		}
	}

	writer.WriteHeader(status)
	_, err := writer.Write(obj)
	if err != nil {
		return err
	}

	return nil
}
