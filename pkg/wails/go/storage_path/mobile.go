//go:build android || ios

package storage_path

import (
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func DataDir() (string, error) {
	dataHome := application.Mobile.StoragePath()
	if dataHome == "" {
		return "", fmt.Errorf("application.PathDataHome is empty")
	}
	return dataHome, nil
}
