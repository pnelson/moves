package moves

import "encoding/json"

// Profile represents the user profile information.
type Profile struct {
	UserId  uint64 `json:"userId"`
	Profile struct {
		FirstDate       string `json:"firstDate"` // yyyyMMdd
		CurrentTimeZone struct {
			Id     string `json:"id"`
			Offset int64  `json:"offset"`
		} `json:"currentTimeZone"`
		Localization struct {
			Language     string  `json:"language,omitempty"`
			Locale       string  `json:"locale,omitempty"`
			FirstWeekDay Weekday `json:"firstWeekDay,omitempty"`
			Metric       bool    `json:"metric,omitempty"`
		} `json:"localization,omitempty"`
		CaloriesAvailable bool   `json:"caloriesAvailable"`
		Platform          string `json:"platform"`
	} `json:"profile"`
}

// Profile retrieves the user profile information for the authenticated user.
func (c *Client) Profile() (*Profile, error) {
	resp, err := c.Get(c.BaseURI + "/user/profile")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	profile := &Profile{}
	err = json.NewDecoder(resp.Body).Decode(profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
