//go:build android || ios

package wails

import "github.com/wailsapp/wails/v3/pkg/application"

// SecureGet gets a value from mobile secure storage.
func SecureGet(key string) string {
	return application.Mobile.SecureGet(key)
}

// SecureDelete deletes a value from mobile secure storage.
func SecureDelete(key string) {
	application.Mobile.SecureDelete(key)
}
