package users

import (
	"fmt"
	"time"

	"github.com/genesixx/coalibot/utils"

	"github.com/nlopes/slack"
)

func Calba(option string, event *utils.Message) bool {
	loc, _ := time.LoadLocation("Europe/Paris")
	stage, _ := time.ParseInLocation("2006-01-02 15:04", "2018-12-03 17:15", loc)
	ts := -time.Since(stage)

	ts += +Decisecond / 2
	d := (ts / Day)
	ts = ts % Day
	h := ts / time.Hour
	ts = ts % time.Hour
	m := ts / time.Minute
	ts = ts % time.Minute
	s := ts / time.Second
	utils.PostMsg(event, slack.MsgOptionText(fmt.Sprintf("Libération dans %02d jours, %02d heures %02d minutes %02d secondes", d, h, m, s), false))
	return true
}
