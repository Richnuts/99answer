package listing

import (
	"net/http"
	"net/url"
)

type ListingClient struct {
	BaseURL string
	Client  *http.Client
}

func NewListingClient(baseURL string) *ListingClient {
	return &ListingClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (lc *ListingClient) GetListings(params url.Values) (*http.Response, error) {
	u := lc.BaseURL + "/listings?" + params.Encode()
	return lc.Client.Get(u)
}

func (lc *ListingClient) CreateListing(params url.Values) (*http.Response, error) {
	u := lc.BaseURL + "/listings"
	return lc.Client.PostForm(u, params)
}
