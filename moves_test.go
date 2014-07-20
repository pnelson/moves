package moves

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestTimeUnmarshalJSON(t *testing.T) {
	var tm Time

	err := json.Unmarshal([]byte(`"00000000T000000Z"`), &tm)
	if err == nil {
		t.Fatalf("expected error but got time %v", time.Time(tm))
	}

	if !strings.Contains(err.Error(), "range") {
		t.Errorf("unexpected error\nhave %v\nwant out of range error", err)
	}
}
