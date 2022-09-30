package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

func GetTwitterApi() *anaconda.TwitterApi {
    anaconda.SetConsumerKey("API Key")
    anaconda.SetConsumerSecret("API Secret")
    api := anaconda.NewTwitterApi("Access Token", "Access Token Secret")
    return api
}

func main(){
    api := GetTwitterApi()
    text := "失せろ"

    tweet, err := api.PostTweet(text, nil)
    if err != nil {
        panic(err)
        }
    fmt.Print(tweet.Text)
}