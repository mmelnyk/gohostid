package gohostid

import (
	"crypto/sha256"
	"encoding/base64"
)

var (
	iscached bool
	cached   string
)

// GetID returns persistent unique host id based
func GetID() (string, error) {
	if !iscached {
		rawid, err := getID()

		if err != nil {
			return "", err
		}

		hashedid := sha256.Sum256([]byte(rawid))
		cached = base64.RawURLEncoding.EncodeToString(hashedid[:])[:24]
		iscached = true
	}

	return cached, nil
}
