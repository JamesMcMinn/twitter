package twitter

import (
	"encoding/json"
)

type Tweet struct {
	Id             int64       `json:"id"`
	Text           string      `json:"text"`
	Coordinates    interface{} `json:"coordinates"`
	Created_at_str string      `json:"created_at"`
	Lang           string      `json:"lang"`
	Timestamp      uint64
	Entities       struct {
		Hashtags []struct {
			Text    string
			Indices []int
		} `json:"hashtags"`
		Urls []struct {
			Url          string
			Display_url  string
			Expanded_url string
			Indices      []int
		} `json:"urls"`
		User_mentions []struct {
			Id          int
			Screen_name string
			Name        string
			Indices     []int
		} `json:"user_mentions"`
		Media []struct {
			Media_url_https string
			Media_url       string
		} `json:"media"`
	} `json:"entities"`
	Place interface{} `json:"place"`
	User  struct {
		Statuses_count          int
		Name                    string
		Description             string
		Location                string
		Verified                bool
		Followers_count         int
		Profile_image_url_https string
		Url                     string
		Screen_name             string
		Friends_count           int
		Created_at              string
		Id                      uint64
	} `json:"user"`
	Retween bool
}

func JSONtoTweet(line []byte) Tweet {
	var t Tweet
	json.Unmarshal(line, &t)
	return t
}

func TweetToJSON(tweet Tweet) ([]byte, error) {
	return json.Marshal(&tweet)
}
