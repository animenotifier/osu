package osu

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aerogo/http/client"
)

// APIKey is the Osu API key.
var APIKey = ""

// User is a user in Osu.
type User struct {
	UserID        string        `json:"user_id"`
	UserName      string        `json:"username"`
	Count300      string        `json:"count300"`
	Count100      string        `json:"count100"`
	Count50       string        `json:"count50"`
	PlayCount     string        `json:"playcount"`
	RankedScore   string        `json:"ranked_score"`
	TotalScore    string        `json:"total_score"`
	PPRank        string        `json:"pp_rank"`
	Level         string        `json:"level"`
	PPRaw         string        `json:"pp_raw"`
	Accuracy      string        `json:"accuracy"`
	CountryRankSS string        `json:"count_rank_ss"`
	CountryRankS  string        `json:"count_rank_s"`
	CountryRankA  string        `json:"count_rank_a"`
	Country       string        `json:"country"`
	PPCountryRank string        `json:"pp_country_rank"`
	Events        []interface{} `json:"events"`
}

// GetUser ...
func GetUser(nick string) (*User, error) {
	if APIKey == "" {
		return nil, errors.New("API key not set")
	}

	users := []*User{}
	response, err := client.Get("https://osu.ppy.sh/api/get_user?u="+nick+"&type=string&k="+APIKey).Header("Accept", "application/json").EndStruct(&users)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %d", response.StatusCode())
	}

	if len(users) == 0 {
		return nil, errors.New("User not found")
	}

	return users[0], nil
}
