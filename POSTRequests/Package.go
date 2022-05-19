package http_helper

import (
	"encoding/json"
	"net/http"
)

type HttpHelper struct {
}

func (h HttpHelper) DecodePostRequest(r *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(obj)
	if err != nil {
		return err
	}
	return nil
}
