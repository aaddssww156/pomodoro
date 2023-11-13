package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getlantern/systray"
)

const (
	// workTime = 15
	workTime = 25 * time.Second
	restTime = 5 * time.Second
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
	systray.AddSeparator()

	startMenu := systray.AddMenuItem("Начать", "")
	pauseMenu := systray.AddMenuItem("Остановить", "")
	stopMenu := systray.AddMenuItem("Сбросить", "")
	exitMenu := systray.AddMenuItem("Выход", "")

	seconds := make(chan int)
	seconds <- 0
	go func() {
		for {
			select {
			case <-startMenu.ClickedCh:
				startTimer(seconds)
				return
			case <-pauseMenu.ClickedCh:
				pauseTimer()
				return
			case <-stopMenu.ClickedCh:
				stopTimer()
				return
			case <-exitMenu.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	go func() {
		for {
			// systray.SetTooltip(time.Now().String())
			systray.SetTooltip(getTime(seconds))
			time.Sleep(1 * time.Second)
		}
	}()
}

func stopTimer() {
	// startTime := time.Now()

}

func getTime(seconds chan int) string {
	log.Println("getTime")
	return fmt.Sprintf("%d", <-seconds)
}

func pauseTimer() {
}

func startTimer(seconds chan int) {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		select {
		case t := <-ticker.C:
			seconds <- +1
			// log.Println(<-seconds)
			log.Println(t)
		}
	}()
}

func onExit() {
	log.Println("exited")
}
