package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

var services = []string{"Facebook",
	"YouTube",
	"Twitter",
	"Instagram",
	"Blogger",
	"GooglePlus",
	"Twitch",
	"Reddit",
	"Ebay",
	"Wordpress",
	"Pinterest",
	"Yelp",
	"Slack",
	"Github",
	"Basecamp",
	"Tumblr",
	"Flickr",
	"Pandora",
	"ProductHunt",
	"Steam",
	"MySpace",
	"Foursquare",
	"OkCupid",
	"Vimeo",
	"UStream",
	"Etsy",
	"SoundCloud",
	"BitBucket",
	"Meetup",
	"CashMe",
	"DailyMotion",
	"Aboutme",
	"Disqus",
	"Medium",
	"Behance",
	"Photobucket",
	"Bitly",
	"CafeMom",
	"coderwall",
	"Fanpop",
	"deviantART",
	"GoodReads",
	"Instructables",
	"Keybase",
	"Kongregate",
	"LiveJournal",
	"StumbleUpon",
	"AngelList",
	"LastFM",
	"Slideshare",
	"Tripit",
	"Fotolog",
	"Vine",
	"PayPal",
	"Dribbble",
	"Imgur",
	"Tracky",
	"Flipboard",
	"Vk",
	"kik",
	"Codecademy",
	"Roblox",
	"Gravatar",
	"Trip",
	"Pastebin",
	"Coinbase",
	"BlipFM",
	"Wikipedia",
	"Ello",
	"StreamMe",
	"IFTTT",
	"WebCredit",
	"CodeMentor",
	"Soupio",
	"Fiverr",
	"Trakt",
	"Hackernews",
	"five00px",
	"Spotify",
	"POF",
	"Houzz",
	"Contently",
	"BuzzFeed",
	"TripAdvisor",
	"HubPages",
	"Scribd",
	"Venmo",
	"Canva",
	"CreativeMarket",
	"Bandcamp",
	"Wikia",
	"ReverbNation",
	"Wattpad",
	"Designspiration",
	"ColourLovers",
	"eyeem",
	"KanoWorld",
	"AskFM",
	"Smashcast",
	"Badoo",
	"Newgrounds",
	"younow",
	"Patreon",
	"Mixcloud",
	"Gumroad",
	"Quora"}

const endPoint string = "https://namechk.com/"

func main() {

	type ValidToken struct {
		Valid string `json:"valid"`
	}

	type userNameCheckResponseObject struct {
		Username     string `json:"username"`
		Available    bool   `json:"available"`
		Status       string `json:"status"`
		FailedReason string `json:"failed_reason"`
		CallbackURL  string `json:"callback_url"`
	}

	username := flag.String("u", "test", "Username")
	flag.Parse()

	res, err := http.PostForm(endPoint, url.Values{"q": {*username}})
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, _ := ioutil.ReadAll(res.Body)

	token := ValidToken{}
	json.Unmarshal(body, &token)

	for _, service := range services {

		res, err := http.PostForm("https://namechk.com/services/check", url.Values{"token": {string(token.Valid)}, "fat": {"xwSgxU58x1nAwVbP6+mYSFLsa8zkcl2q6NcKwc8uFm+TvFbN8LaOzmLOBDKza0ShvREINUhbwwljVe30LbKcQw=="}, "service": {strings.ToLower(service)}})
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		body, _ := ioutil.ReadAll(res.Body)

		responseObject := userNameCheckResponseObject{}
		json.Unmarshal(body, &responseObject)

		output := fmt.Sprintf("Service %s : user %s is %s @ %s", service, responseObject.Username, responseObject.Status, responseObject.CallbackURL)
		fmt.Println(output)

	}

}
