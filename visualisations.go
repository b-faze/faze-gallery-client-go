package fazegalleryclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type VisualisationClient Client

func (vc *VisualisationClient) GetAll(authToken *string) (*[]Visualisation, error) {
	c := (*Client)(vc)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/visualisations", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	orders := []Visualisation{}
	err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (c *Client) Get(id string) (*Visualisation, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/visualisations/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := Visualisation{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) Create(vis *Visualisation) (*Visualisation, error) {
	rb, err := json.Marshal(vis)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/visualisations", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	result := Visualisation{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateOrder(id string, vis *Visualisation) (*Visualisation, error) {
	rb, err := json.Marshal(vis)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/visualisations/%s", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	result := Visualisation{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Delete(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/visualisations/%s", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted visualisation" {
		return errors.New(string(body))
	}

	return nil
}
