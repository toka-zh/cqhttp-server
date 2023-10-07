package cron

import "github.com/robfig/cron/v3"

//var Cron cron.Cron

func init() {
	c := cron.New()
	_, err := c.AddFunc("@daily", SyncPicture)
	if err != nil {
		return
	}
}
