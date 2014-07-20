package moves

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Summary represents the activity summary.
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

// Summary retrieves the activity summaries for the authenticated user
// for the day, week, or month provided. See the API documentation for
// expected period formats.
func (c *Client) Summary(period string) ([]Summary, error) {
	return c.SummaryQuery("/user/summary/daily/"+period, nil)
}

// SummaryRange retrieves the activity summaries for the authenticated user
// over the provided date range. See the API documentation for expected
// date formats and limitations.
func (c *Client) SummaryRange(from, to string) ([]Summary, error) {
	query := url.Values{"from": {from}, "to": {to}}
	return c.SummaryQuery("/user/summary/daily", query)
}

// SummaryPast retrieves the activity summaries for the authenticated user
// over the past provided days, including today. See the API documentation
// for limitations.
func (c *Client) SummaryPast(past int) ([]Summary, error) {
	query := url.Values{"pastDays": {strconv.Itoa(past)}}
	return c.SummaryQuery("/user/summary/daily", query)
}

// SummaryQuery retrieves the activity summaries for the authenticated user.
// This method exists so that other query string parameters may be specified.
// See the API documentation for additional query parameters.
func (c *Client) SummaryQuery(path string, query url.Values) ([]Summary, error) {
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
