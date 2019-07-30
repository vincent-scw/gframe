package simulator

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// InjectPlayers injects amount of players
func InjectPlayers(amount int) {
	for i := 0; i < amount; i++ {
		name := getRandPlayerName()
		go inject(name)
	}
}

func inject(name string) {
	formData := url.Values{
		"client_id":     {"player_api"},
		"client_secret": {"999999"},
		"grant_type":    {"password"},
		"username":      {name},
		"password":      {"123"},
	}
	resp, err := http.PostForm("http://localhost:9096/token", formData)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

func getRandPlayerName() string {
	fn := firstNames[rand.Intn(len(firstNames))]
	ln := lastNames[rand.Intn(len(lastNames))]

	return fmt.Sprintf("%s %s", fn, ln)
}
