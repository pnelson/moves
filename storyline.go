package moves

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Storyline struct {
	Date         string    `json:"date"` // yyyyMMdd
	Summary      []Summary `json:"summary"`
	Segments     []Segment `json:"segments"`
	CaloriesIdle int64     `json:"caloriesIdle,omitempty"`
	LastUpdate   Time      `json:"lastUpdate,omitempty"`
}

func (c *Client) Storyline(period string) ([]Storyline, error) {
	return c.StorylineQuery("/user/storyline/daily/"+period, nil)
}

func (c *Client) StorylineRange(from, to string) ([]Storyline, error) {
	query := url.Values{"from": {from}, "to": {to}}
	return c.StorylineQuery("/user/storyline/daily", query)
}

func (c *Client) StorylinePast(past int) ([]Storyline, error) {
	query := url.Values{"pastDays": {strconv.Itoa(past)}}
	return c.StorylineQuery("/user/storyline/daily", query)
}

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
