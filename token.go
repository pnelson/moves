package moves

import (
	"encoding/json"
	"os"
	"time"
)

// Token contains an end-user's tokens.
// This is the data you must store to persist authentication.
type Token struct {
	Access  string
	Refresh string
	Expiry  time.Time
	UserId  uint64
}

// Expired returns true if the token is expired.
func (t *Token) Expired() bool {
	if t.Expiry.IsZero() {
		return false
	}

	return t.Expiry.Before(time.Now().UTC())
}

// Cache specifies the methods that implement a Token cache.
type Cache interface {
	Token() (*Token, error)
	PutToken(token *Token) error
}

// CacheFile implements Cache. Its value is the name of the file in which
// the Token is stored in JSON format.
type CacheFile string

func (f CacheFile) Token() (*Token, error) {
	file, err := os.Open(string(f))
	if err != nil {
		return nil, err
	}

	defer file.Close()

	token := &Token{}
	err = json.NewDecoder(file).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (f CacheFile) PutToken(token *Token) error {
	file, err := os.OpenFile(string(f), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	err = json.NewEncoder(file).Encode(token)
	if err != nil {
		file.Close()
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
