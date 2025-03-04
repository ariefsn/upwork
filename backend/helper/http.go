package helper

import (
	"io"
	"net/http"
)

type Http struct {
	Method  string
	Url     string
	Payload io.Reader
	Headers map[string]string
}

func (h *Http) Send() (*http.Response, error) {
	req, err := http.NewRequest(h.Method, h.Url, h.Payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	for k, v := range h.Headers {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
