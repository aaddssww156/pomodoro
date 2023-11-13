package main

import (
	"log"
	"os"

	"github.com/getlantern/systray"
)

const (
	workTime = 15
	restTime = 3
)

func main() {
	systray.Run(onReady, onExit)

	// Sets the icon of a menu item. Only available on Mac and Windows.
	// notificationWork := toast.Notification{
	// 	AppID:   "Pomodoro",
	// 	Title:   "Work",
	// 	Message: "Started 25 min work!",
	// }

	// notificationRest := toast.Notification{
	// 	AppID:   "Pomodoro",
	// 	Title:   "Rest",
	// 	Message: "Have a rest for 5 minutes!",
	// }

	// ticker := time.NewTicker(1 * time.Second)
	// time := 0
	// done := make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			log.Println("done")
	// 			// err := notificationRest.Push()
	// 			// if err != nil {
	// 			// log.Fatalln(err)
	// 			// }
	// 		case <-ticker.C:
	// 			time++
	// 			log.Println(time)
	// 			if time >= workTime {
	// 				done <- true
	// 				ticker.Stop()
	// 			}
	// 		}
	// 	}
	// }()
}

func getIcon(s string) []byte {
	file, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func onReady() {
	systray.SetIcon(getIcon("file.ico"))

	systray.SetTitle("Pomodoro")
	systray.SetTooltip("Pomodoro")

	systray.AddSeparator()

	startMenu := systray.AddMenuItem("Начать", "")
	pauseMenu := systray.AddMenuItem("Остановить", "")
	stopMenu := systray.AddMenuItem("Сбросить", "")

	go func() {
		for {
			select {
			case <-startMenu.ClickedCh:
				startTimer()
				return
			case <-pauseMenu.ClickedCh:
				pauseTimer()
				return
			case <-stopMenu.ClickedCh:
				stopTimer()
				return
			}
		}
	}()
}

func stopTimer() {
}

func pauseTimer() {
}

func startTimer() {

}

func onExit() {

}
