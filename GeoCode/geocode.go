package geocode

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {
	log.SetFlags(0)
}

const bUrl = "https://maps.googleapis.com/maps/api/geocode/json?latlng=%v,%v"

type apiTripper struct {
	http.RoundTripper
	k string
}

func (o apiTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Print("-> ", r.URL)

	q := r.URL.Query()
	q.Set("key", o.k)
	r.URL.RawQuery = q.Encode()
	return o.RoundTripper.RoundTrip(r)
}

type Result struct {
	AddressComponents []struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	} `json:"address_components"`
	FormattedAddress string `json:"formatted_address"`
	Geometry         struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
	PlaceId string `json:"place_id"`
}

type Client struct {
	client *http.Client
	tick   *time.Ticker
}

func NewClient(tick time.Duration, apiKey string) *Client {
	return &Client{
		client: &http.Client{Transport: apiTripper{http.DefaultTransport, apiKey}},
		tick:   time.NewTicker(tick),
	}
}

func (c *Client) ReverseGeocode(lat, lng float64) ([]Result, error) {
	<-c.tick.C

	rsp, err := c.client.Get(fmt.Sprintf(bUrl, lat, lng))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var v struct {
		Results []Result `json:"results"`
		Status  string   `json:"status"`
	}
	if err := json.NewDecoder(rsp.Body).Decode(&v); err != nil {
		return nil, err
	}
	if v.Status != "OK" {
		return nil, errors.New(v.Status)
	}
	return v.Results, nil
}
