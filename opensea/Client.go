package opensea

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"austinkline/go-opensea/network"
	"austinkline/go-opensea/opensea/types"
)

const (
	BaseURLV1    = "https://api.opensea.io/api/v1"
	ActionAssets = "assets"
	ActionBundles = "bundles"
)

type IClient interface {
	RetrieveAssets(r types.RetrieveAssetsRequest) (res types.RetrieveAssetsResponse, err error)
	RetrieveBundles(r types.RetrieveBundlesRequest) (res types.RetrieveBundlesResponse, err error)
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL, httpClient: http.Client{}}
}

// TODO: Support for API Key
type Client struct {
	BaseURL    string
	httpClient http.Client
}

func (c *Client) RetrieveBundles(r types.RetrieveBundlesRequest) (res types.RetrieveBundlesResponse, err error) {
	bodyBytes, err := c.dispatchRequest(ActionBundles, r)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyBytes, &res)
	return
}

func (c *Client) RetrieveAssets(r types.RetrieveAssetsRequest) (res types.RetrieveAssetsResponse, err error) {
	bodyBytes, err := c.dispatchRequest(ActionAssets, r)
	if err != nil {
		return
	}

	err = json.Unmarshal(bodyBytes, &res)
	return
}

func (c *Client) buildURL(action string) (url string) {
	url = fmt.Sprintf("%s/%s", c.BaseURL, action)
	return
}

func (c *Client) dispatchRequest(action string, data interface{}) (resBytes []byte, err error) {
	url := c.buildURL(action)
	rb := network.RequestBuilder{}
	req, err := rb.BuildGetRequest(url, data)
	if err != nil {
		return
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = network.ErrNonSuccessResponse{Response: *res}
		return
	}

	resBytes, err = ioutil.ReadAll(res.Body)
	return
}
