package config

import "errors"

type Data struct {
	Responses []Response
	Config    Config
}

type Response struct {
	Label  string
	Path   string
	Method string
	Code   int
	Body   string
}

type Config struct {
	Port    int
	LogPath string
}

func (d *Data) Contain(path string, method string) (Response, error) {
	for _, r := range d.Responses {
		if r.Path == path && (r.Method == method || method == "OPTIONS") {
			return r, nil
		}
	}
	return Response{}, errors.New("not found")
}
