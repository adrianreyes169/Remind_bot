package main

import (
	"github.com/go-co-op/gocron/v2"
)

func scheduler(reminders []Reminder, config Config) error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	for i := 0; i < len(reminders); i++ {
		_, err = s.NewJob(gocron.CronJob(reminders[i].Time, false), gocron.NewTask(sendMessage, reminders[i].Message, config))
		if err != nil {
			return err
		}
	}
	s.Start()

	return nil
}
