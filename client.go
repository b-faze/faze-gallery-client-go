package fazegalleryclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "http://localhost:19090"

type ClientCtx struct {
	VisApi *VisualisationClient
}

func NewClientCtx(host *string) (*ClientCtx, error) {
	c, err := newClient(host)
	if err != nil {
		return nil, err
	}

	return &ClientCtx{
		VisApi: newVisualisationClient(c),
	}, nil
}

type client struct {
	HostURL    string
	HTTPClient *http.Client
}

func newClient(host *string) (*client, error) {
	c := client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	return &c, nil
}

func (c *client) doRequest(req *http.Request) ([]byte, error) {

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
