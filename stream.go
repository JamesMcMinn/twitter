package twitter

import (
	"github.com/araddon/httpstream"
	"github.com/mrjones/oauth"
)

func FillStream(channel chan Tweet, consumerKey, consumerSecret, ot, osec string) {
	httpstream.OauthCon = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	at := oauth.AccessToken{
		Token:  ot,
		Secret: osec,
	}

	client := httpstream.NewOAuthClient(&at, httpstream.OnlyTweetsFilter(func(line []byte) {
		channel <- JSONtoTweet(line)
	}))

	err := client.Sample(nil)
	if err != nil {
		httpstream.Log(httpstream.ERROR, err.Error())
	}
}
