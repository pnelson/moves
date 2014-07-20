package moves

import (
	"net/url"
	"strconv"
)

// Place represents a place the user has been to.
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

// Places retrieves the places for the authenticated user for the day, week,
// or month provided. See the API documentation for expected period formats.
func (c *Client) Places(period string) ([]Storyline, error) {
	return c.StorylineQuery("/user/places/daily/"+period, nil)
}

// PlacesRange retrieves the places for the authenticated user over the
// provided date range. See the API documentation for expected date formats
// and limitations.
func (c *Client) PlacesRange(from, to string) ([]Storyline, error) {
	query := url.Values{"from": {from}, "to": {to}}
	return c.StorylineQuery("/user/places/daily", query)
}

// PlacesPast retrieves the places for the authenticated user over the past
// provided days, including today. See the API documentation for limitations.
func (c *Client) PlacesPast(past int) ([]Storyline, error) {
	query := url.Values{"pastDays": {strconv.Itoa(past)}}
	return c.StorylineQuery("/user/places/daily", query)
}
