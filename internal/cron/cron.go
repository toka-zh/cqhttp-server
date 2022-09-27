package cron

import "github.com/robfig/cron/v3"

//var Cron cron.Cron

func init() {
	cron := cron.New()
	cron.AddFunc("@daily", SyncPicture)
}
