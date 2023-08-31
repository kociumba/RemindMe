package main

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CreateSettingsOrAppend(append string) error {
	f, err := os.OpenFile("settings.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.WriteString(append + "\n"); err != nil {
		panic(err)
	}

	return nil
}

func (a *App) RemoveSettings(IndexOfSettingsToRemove int) string {
	returnedSettings, err := os.ReadFile("settings.txt")
	if err != nil {
		return err.Error()
	}

	settings := strings.Split(string(returnedSettings), "\n")

	if IndexOfSettingsToRemove >= 0 && IndexOfSettingsToRemove < len(settings) {

		settings = append(settings[:IndexOfSettingsToRemove], settings[IndexOfSettingsToRemove+1:]...)
	} else {
		return "Invalid index"
	}

	updatedSettings := strings.Join(settings, "\n")

	err = os.WriteFile("settings.txt", []byte(updatedSettings), 0644)
	if err != nil {
		return err.Error()
	}

	return "Settings removed successfully"
}

func (a *App) ReadSettings() string {
	returnedSettings, _ := os.ReadFile("settings.txt")

	return string(returnedSettings)
}
