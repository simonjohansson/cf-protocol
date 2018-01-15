package in

import (
	"encoding/json"
	"errors"
)

func Execute(input []byte) (string, error) {
	var outdata struct {
		Version interface{} `json:"version"`
	}

	err := json.Unmarshal(input, &outdata)
	if err != nil {
		return "", err
	}
	if outdata.Version == nil {
		return "", errors.New("missing version")
	}
	outbytes, err := json.Marshal(outdata)
	return string(outbytes), err
}
