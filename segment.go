package moves

type Segment struct {
	Type       string     `json:"type"`
	StartTime  Time       `json:"startTime"`
	EndTime    Time       `json:"endTime"`
	Place      Place      `json:"place,omitempty"`
	Activities []Activity `json:"activities,omitempty"`
	LastUpdate Time       `json:"lastUpdate,omitempty"`
}
