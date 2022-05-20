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

// func (p *PostMessageEndpoint) handler(w http.ResponseWriter, r *http.Request) {
// 	message := types.Message{}
// 	err := http_helper.HttpHelper{}.DecodePostRequest(r, &message)
// 	if err != nil {
// 		p.log.Error("can not decode post message", logger.NewParameter("request", r.GetBody))
// 		return
// 	}
