package error

import "encoding/json"

type ResponseError struct {
	Message string `json:"error"`
}

func (re *ResponseError) Error() []byte {
	jsonResponse, err := json.Marshal(re)
	if err != nil {
		return []byte{}
	}
	return jsonResponse
}
