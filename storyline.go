package moves

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Storyline represents activity and location storyline.
type Storyline struct {
	Date         string    `json:"date"` // yyyyMMdd
	Summary      []Summary `json:"summary"`
	Segments     []Segment `json:"segments"`
	CaloriesIdle int64     `json:"caloriesIdle,omitempty"`
	LastUpdate   Time      `json:"lastUpdate,omitempty"`
}

// Storyline retrieves the activity and location breakdown for the
// authenticated user for the day, week, or month provided. See the API
// documentation for expected period formats.
func (c *Client) Storyline(period string) ([]Storyline, error) {
	return c.StorylineQuery("/user/storyline/daily/"+period, nil)
}

// StorylineRange retrieves the activity and location breakdown for the
// authenticated user over the provided date range. See the API documentation
// for expected date formats and limitations.
func (c *Client) StorylineRange(from, to string) ([]Storyline, error) {
	query := url.Values{"from": {from}, "to": {to}}
	return c.StorylineQuery("/user/storyline/daily", query)
}

// StorylinePast retrieves the activity and location breakdown for the
// authenticated user over the past provided days, including today.
// See the API documentation for limitations.
func (c *Client) StorylinePast(past int) ([]Storyline, error) {
	query := url.Values{"pastDays": {strconv.Itoa(past)}}
	return c.StorylineQuery("/user/storyline/daily", query)
}

// StorylineQuery retrieves the activity and location breakdown for the
// authenticated user. This method exists so that other query string
// parameters may be specified. See the API documentation for additional
// query parameters.
func (c *Client) StorylineQuery(path string, query url.Values) ([]Storyline, error) {
	resp, err := c.get(path, query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var rv []Storyline
	err = json.NewDecoder(resp.Body).Decode(&rv)
	if err != nil {
		return nil, err
	}

	return rv, nil
}
