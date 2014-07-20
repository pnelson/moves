package moves

import (
	"net/url"
	"strconv"
)

type Place struct {
	Id                    uint64   `json:"id,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Type                  string   `json:"type"`
	FoursquareId          string   `json:"foursquareId,omitempty"`
	FoursquareCategoryIds []string `json:"foursquareCategoryIds,omitempty"`
	Location              struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}

func (c *Client) Places(period string) ([]Storyline, error) {
	return c.StorylineQuery("/user/places/daily/"+period, nil)
}

func (c *Client) PlacesRange(from, to string) ([]Storyline, error) {
	query := url.Values{"from": {from}, "to": {to}}
	return c.StorylineQuery("/user/places/daily", query)
}

func (c *Client) PlacesPast(past int) ([]Storyline, error) {
	query := url.Values{"pastDays": {strconv.Itoa(past)}}
	return c.StorylineQuery("/user/places/daily", query)
}
