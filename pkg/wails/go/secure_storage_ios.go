//go:build ios

package wails

import "github.com/wailsapp/wails/v3/pkg/application"

// SecureSet stores a value in iOS secure storage.
func SecureSet(key, value string) error {
	application.IOS.SecureSet(key, value)
	return nil
}
