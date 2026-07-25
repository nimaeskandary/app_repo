package main

import (
	"context"
	"embed"
	"log/slog"
	"sync"

	"github.com/nimaeskandary/app_repo/cmd/gordle_app/gordle_config"
	"github.com/nimaeskandary/app_repo/cmd/gordle_app/internal"
	"github.com/nimaeskandary/app_repo/cmd/gordle_app/internal/bridge"
	di "github.com/nimaeskandary/app_repo/pkg/di/go"
	greet_core "github.com/nimaeskandary/app_repo/pkg/greet/go/core"
	observability "github.com/nimaeskandary/app_repo/pkg/observability/go"
	obs_core "github.com/nimaeskandary/app_repo/pkg/observability/go/core"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	application.RegisterEvent[string]("gordle_app:navigate:path")
}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	ctx := context.Background()

	var greetService greet_core.GreetService
	var logger obs_core.Logger
	fxApp := di.CreateFxAppAndExtract(internal.ModuleList(gordle_config.Bytes), &greetService, &logger)
	if fxApp == nil {
		slog.Error("dependency injection system failed to initialize")
		return
	}
	if err := fxApp.Err(); err != nil {
		slog.Error("dependency injection system failed to initialize", "error", err)
		return
	}

	logger.Info(ctx, "starting dependency injection system")
	if err := fxApp.Start(ctx); err != nil {
		logger.Error(ctx, "dependency injection system failed to start", "error", err)
		return
	}
	var stopFxOnce sync.Once
	var stop = func() {
		stopFxOnce.Do(func() {
			logger.Info(ctx, "stopping dependency injection system")
			if err := fxApp.Stop(ctx); err != nil {
				logger.Error(ctx, "dependency injection system failed to stop gracefully", "error", err)
			}
		})
	}

	defer stop()

	wailsApp := application.New(application.Options{
		Name:        "Gordle",
		Description: "Wails3 react example app",
		Logger:      observability.AsSlog(logger),
		Services: []application.Service{
			application.NewService(&bridge.GreetService{GreetService: greetService}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		OnShutdown: stop,
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	wailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "Gordle",
		// using phone dimensions
		Width:  402,
		Height: 874,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarDefault,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		URL:              "/",
	})

	// Run the application. This blocks until the application has been exited.
	err := wailsApp.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		logger.Error(ctx, "wails application failed to run", "error", err)
	}
}
