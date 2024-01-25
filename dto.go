package zealygo

import "time"

type userRequestDto struct {
	email           *string
	discordId       *string
	twitterId       *string
	discordHandle   *string
	twitterUsername *string
	ethAddress      *string
}

type ErrorDTO struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type UserZealy struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ID        string    `json:"id"`
	Addresses struct {
		Other string `json:"other"`
	} `json:"addresses"`
	Avatar                interface{}   `json:"avatar"`
	DiscordHandle         interface{}   `json:"discordHandle"`
	TwitterUsername       string        `json:"twitterUsername"`
	Country               string        `json:"country"`
	City                  string        `json:"city"`
	TwitterFollowersCount int           `json:"twitterFollowersCount"`
	TweetCount            int           `json:"tweetCount"`
	SocialAccounts        []interface{} `json:"socialAccounts"`
	Guilds                []interface{} `json:"guilds"`
	DisplayedInformation  []string      `json:"displayedInformation"`
	Goal                  string        `json:"goal"`
	DeletedAt             interface{}   `json:"deletedAt"`
	RestoredAt            interface{}   `json:"restoredAt"`
	Referrer              interface{}   `json:"referrer"`
	Interests             []string      `json:"interests"`
	Xp                    int           `json:"xp"`
	Rank                  int           `json:"rank"`
	Invites               int           `json:"invites"`
	Role                  string        `json:"role"`
	Level                 int           `json:"level"`
	IsBanned              bool          `json:"isBanned"`
	Karma                 int           `json:"karma"`
	ReferrerURL           interface{}   `json:"referrerUrl"`
}
