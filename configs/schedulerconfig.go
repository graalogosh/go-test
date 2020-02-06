package configs

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func scheduledTask() {
	fmt.Println("Are you still there?")
}

func SetupScheduler() {
	gocron.Every(3).Seconds().Do(scheduledTask)
	<-gocron.Start()
}
