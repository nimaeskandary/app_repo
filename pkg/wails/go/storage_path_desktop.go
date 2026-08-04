//go:build !android && !ios

package wails

import (
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func DataDir() (string, error) {
	dataHome := application.Path(application.PathDataHome)
	if dataHome == "" {
		return "", fmt.Errorf("application.PathDataHome is empty")
	}
	return dataHome, nil
}
