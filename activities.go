package moves

import "encoding/json"

// Activities represents the list of supported activities.
type Activities []struct {
	Activity string `json:"activity"`
	Group    string `json:"group,omitempty"`
	Geo      bool   `json:"geo"`
	Place    bool   `json:"place"`
	Color    string `json:"color"`
	Units    string `json:"units"`
}

// ActivityList retrieves the list of supported activities.
func (c *Client) ActivityList() (Activities, error) {
	resp, err := c.Get(c.BaseURI + "/activities")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var rv Activities
	err = json.NewDecoder(resp.Body).Decode(&rv)
	if err != nil {
		return nil, err
	}

	return rv, nil
}
