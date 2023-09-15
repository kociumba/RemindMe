package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
	"gopkg.in/ini.v1"
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

	settings := `# all the main settings go here
	show_notification = true
	test_setting = false`

	if _, err := os.Stat("settings.ini"); err == nil {
		// TODO: add a check if the settings file is valid
		return
	} else if errors.Is(err, os.ErrNotExist) {
		f, _ := os.Create("settings.ini")
		defer f.Close()

		f.WriteString(settings)

	} else {
		f, _ := os.Create("settings.ini")
		defer f.Close()
	}

}

func (a *App) ScheduleNotification(timeInSeconds int, title string, messageBody string) string {
	go sendNotification(timeInSeconds, title, messageBody)

	return "notification scheduled"
}

// TODO: this is wrong, handle the schedualing in js frontend couse this will trigger a notif even if the app is closed and the reminder is deleted

func sendNotification(timeInSeconds int, title string, messageBody string) {

	time.Sleep(time.Duration(timeInSeconds) * time.Second)

	err := beeep.Notify(title, messageBody, "assets/information.png")
	if err != nil {
		panic(err)
	}
}

func (a *App) UpdateSettings(section string, key string, newValue string) string {

	cfg, err := ini.Load("settings.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	cfg.Section(section).Key(key).SetValue(newValue)
	cfg.SaveTo("settings.ini")

	return "setting changed"
}

func (a *App) LoadSettings(section string, key string) string {

	cfg, err := ini.Load("settings.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	return cfg.Section(section).Key(key).String()
}

// this is for an easter egg couse i'm bored

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {

	messages := [7]string{"hello", "easter", "egg", "fuck", "you", "go", "away"}

	mutex.Lock()
	counter++
	fmt.Fprintf(w, messages[counter])
	mutex.Unlock()

}

func (a *App) OpenEasterEgg(port string) {
	http.HandleFunc("/", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
