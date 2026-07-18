package main

import (
	"context"
	"embed"

	"log"
	"time"

	"github.com/nimaeskandary/wails3-react-polylith/bases/wails_app/app"
	"github.com/nimaeskandary/wails3-react-polylith/bases/wails_app/app/frontend_bridge"
	"github.com/nimaeskandary/wails3-react-polylith/components/greet/go/types"
	"github.com/nimaeskandary/wails3-react-polylith/components/util/go"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	// Register a custom event whose associated data type is string.
	// This is not required, but the binding generator will pick up registered events
	// and provide a strongly typed JS/TS API for them.
	application.RegisterEvent[string]("time")
}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	ctx := context.Background()

	var greetService greet_types.GreetService
	fxApp := util.CreateFxAppAndExtract(app.ModuleList(), &greetService)

	log.Println("starting dependency injection system...")
	if err := fxApp.Start(ctx); err != nil {
		log.Fatalf("dependency injection system failed to start: %v", err)
	}
	defer func() {
		log.Println("stopping dependency injection system...")
		if err := fxApp.Stop(ctx); err != nil {
			log.Printf("dependency injection system failed to stop gracefully: %v", err)
		}
	}()


	wailsApp := application.New(application.Options{
		Name:        "test-react",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(frontend_bridge.NewFrontendBridge(greetService)),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	wailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "Window 1",
		// Window sized to the golden ratio (1000 / 618 ≈ 1.618).
		Width:  1000,
		Height: 618,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		URL:              "/",
	})

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			wailsApp.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := wailsApp.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Print(err)
	}
}
