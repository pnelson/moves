package moves

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Summary struct {
	Date    string `json:"date"` // yyyyMMdd
	Summary []struct {
		Activity string  `json:"activity,omitempty"`
		Group    string  `json:"group,omitempty"`
		Duration float64 `json:"duration"`
		Distance float64 `json:"distance,omitempty"`
		Steps    int64   `json:"steps,omitempty"`
		Calories int64   `json:"calories,omitempty"`
	} `json:"summary"`
	CaloriesIdle int64 `json:"caloriesIdle,omitempty"`
	LastUpdate   Time  `json:"lastUpdate,omitempty"`
}

func (c *Client) Summary(period string) ([]Summary, error) {
	return c.summary("/user/summary/daily/"+period, nil)
}

func (c *Client) SummaryRange(from, to string) ([]Summary, error) {
	query := url.Values{"from": {from}, "to": {to}}
	return c.summary("/user/summary/daily", query)
}

func (c *Client) SummaryPast(past int) ([]Summary, error) {
	query := url.Values{"pastDays": {strconv.Itoa(past)}}
	return c.summary("/user/summary/daily", query)
}

func (c *Client) summary(path string, query url.Values) ([]Summary, error) {
	resp, err := c.get(path, query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var rv []Summary
	err = json.NewDecoder(resp.Body).Decode(&rv)
	if err != nil {
		return nil, err
	}

	return rv, nil
}
