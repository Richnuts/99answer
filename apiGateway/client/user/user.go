package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type UserClient struct {
	BaseURL string
	Client  *http.Client
}

func NewUserClient(baseURL string) *UserClient {
	return &UserClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (uc *UserClient) CreateUser(input CreateUserInput) (*http.Response, error) {
	u := uc.BaseURL + "/users"

	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	return uc.Client.Post(u, "application/json", bytes.NewReader(body))
}

func (uc *UserClient) GetUser(id string) (*http.Response, error) {
	u := uc.BaseURL + "/users/" + id
	return uc.Client.Get(u)
}

func (uc *UserClient) GetUsers(params url.Values) (*http.Response, error) {
	u := uc.BaseURL + "/users?" + params.Encode()
	return uc.Client.Get(u)
}
