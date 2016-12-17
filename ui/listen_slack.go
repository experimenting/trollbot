package ui

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/Sirupsen/logrus"

	"github.com/experimenting/trollbot/troll"
)

// ListenSlack is Listening
func ListenSlack(token string, tr *troll.Troll) error {

	if token == "" {
		return fmt.Errorf("Needed %s\n", "TROLL_SLACK_TOKEN")
	}
	// start a websocket-based Real Time API session
	ws, id := slackConnect(token)
	fmt.Println("mybot ready, ^C exits")
	fmt.Println("current id=", id)

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(id, m)
		// see if we're mentioned

		if m.Type == "message" && (strings.Contains(m.Text, "<@"+id+">") ||
			strings.Contains(m.Text, " troll") ||
			strings.HasPrefix(m.Text, "troll")) {
			// if so try to parse if
			if m.Text == "<@"+id+"> help" || m.Text == "trollbot help" {
				m.Text = fmt.Sprintf("use 'trollbot [keywords|username]'\n Keywords are [help,%s]", strings.Join(tr.GetKeywords(), ","))
				postMessage(ws, m)
			} else {
				go func(m Message) {
					re := regexp.MustCompile("<@([^>.]*)>")
					to := re.FindAllString(m.Text, -1)
					m.Text, err = tr.Troll(m.Text, to)
					if err != nil {
						logrus.Error("return from troll", m, err)
					}
					postMessage(ws, m)
				}(m)
			}
		}
	}
}
