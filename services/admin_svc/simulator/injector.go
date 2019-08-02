package simulator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

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
	token := getToken(name)
	request, err := http.NewRequest("POST", "http://localhost:8080/user/in", nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.AccessToken))
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func getToken(name string) *token {
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
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	t := &token{}
	_ = json.Unmarshal(body, t)
	return t
}

func getRandPlayerName() string {
	fn := firstNames[rand.Intn(len(firstNames))]
	ln := lastNames[rand.Intn(len(lastNames))]

	return fmt.Sprintf("%s %s", fn, ln)
}
