//go:build android

package wails

import (
	"encoding/json/v2"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// SecureSet stores a value in Android's secure storage.
func SecureSet(key, value string) error {
	payload, err := json.Marshal(struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return err
	}

	application.Android.SecureSet(string(payload))
	return nil
}
