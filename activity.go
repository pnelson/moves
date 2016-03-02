package moves

import (
	"net/url"
	"strconv"
)

// Activity represents the activity breakdown.
type Activity struct {
	Activity    string  `json:"activity"`
	Group       string  `json:"group,omitempty"`
	Manual      bool    `json:"manual"`
	StartTime   Time    `json:"startTime,omitempty"`
	EndTime     Time    `json:"endTime,omitempty"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance,omitempty"`
	Steps       int64   `json:"steps,omitempty"`
	Calories    int64   `json:"calories,omitempty"`
	TrackPoints []struct {
		Lat  float64 `json:"lat"`
		Lon  float64 `json:"lon"`
		Time Time    `json:"time"`
	} `json:"trackPoints,omitempty"`
}

// Activities retrieves the activity breakdown for the authenticated user
// for the day, week, or month provided. See the API documentation for
// expected period formats.
func (c *Client) Activities(period string) ([]Storyline, error) {
	return c.StorylineQuery("/user/activities/daily/"+period, nil)
}

// ActivitiesRange retrieves the activity breakdown for the authenticated user
// over the provided date range. See the API documentation for expected
// date formats and limitations.
func (c *Client) ActivitiesRange(from, to string) ([]Storyline, error) {
	query := url.Values{"from": {from}, "to": {to}}
	return c.StorylineQuery("/user/activities/daily", query)
}

// ActivitiesPast retrieves the activity breakdown for the authenticated user
// over the past provided days, including today. See the API documentation
// for limitations.
func (c *Client) ActivitiesPast(past int) ([]Storyline, error) {
	query := url.Values{"pastDays": {strconv.Itoa(past)}}
	return c.StorylineQuery("/user/activities/daily", query)
}
