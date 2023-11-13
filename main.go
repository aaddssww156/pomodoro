package main

import (
	"log"
	"time"

	"github.com/go-toast/toast"
)

func main() {
	notificationWork := toast.Notification{
		AppID:   "Pomodoro",
		Title:   "Work",
		Message: "Started 25 min work!",
	}

	notificationRest := toast.Notification{
		AppID:   "Pomodoro",
		Title:   "Rest",
		Message: "Have a rest for 5 minutes!",
	}

	for {
		err := notificationWork.Push()
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(25 * time.Minute)

		err = notificationRest.Push()
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(5 * time.Minute)
	}
}
