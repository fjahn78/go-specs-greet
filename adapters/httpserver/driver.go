package httpserver

import (
	"io"
	"net/http"
)

type Driver struct {
	BaseURL string
	Client  *http.Client
}

// Curse implements specifications.MeanGreeter.
func (d Driver) Curse(name string) (string, error) {
	return newFunction(d, name)
}

func newFunction(d Driver, name string) (string, error) {
	res, err := d.Client.Get(d.BaseURL + "/curse?name=" + name)
	if err != nil {
		return "", err
	}
	defer res.Body.Close() //nolint:all

	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(greeting), nil
}

func (d Driver) Greet(name string) (string, error) {
	res, err := d.Client.Get(d.BaseURL + "/greet?name=" + name)
	if err != nil {
		return "", err
	}
	defer res.Body.Close() //nolint:all

	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(greeting), nil
}
