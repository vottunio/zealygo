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

type ActivityAnswer struct {
	Data []ClaimedActivityZealy `json:"data"`
}

type ClaimedActivityZealy struct {
	User struct {
		ID               string      `json:"id"`
		Name             string      `json:"name"`
		Avatar           interface{} `json:"avatar"`
		Addresses        interface{} `json:"addresses"`
		DiscordHandle    interface{} `json:"discordHandle"`
		TwitterUsername  string      `json:"twitterUsername"`
		TwitterFollowers int         `json:"twitterFollowers"`
		DiscordID        interface{} `json:"discordId"`
	} `json:"user"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Status     string    `json:"status"`
	Xp         int       `json:"xp"`
	Deleted    bool      `json:"deleted"`
	Submission struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"submission"`
	CommunityID string `json:"communityId"`
	Reward      []struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"reward"`
	ID             string      `json:"id"`
	QuestID        string      `json:"questId"`
	TwitterID      interface{} `json:"twitterId"`
	SubmissionHash interface{} `json:"submissionHash"`
	Mark           interface{} `json:"mark"`
	RewardStatus   interface{} `json:"rewardStatus"`
	DeletedAt      interface{} `json:"deletedAt"`
	RetriedAt      interface{} `json:"retriedAt"`
	LastReviewerID interface{} `json:"lastReviewerId"`
	Name           string      `json:"name"`
	ValidationData struct {
		Value   string `json:"value"`
		Enabled bool   `json:"enabled"`
	} `json:"validationData"`
	Recurrence string `json:"recurrence"`
	Email      string `json:"email"`
}

type ActivityZealy struct {
	Name       string        `json:"name"`
	Content    []interface{} `json:"content"`
	Recurrence string        `json:"recurrence"`
	Deleted    bool          `json:"deleted"`
	CreatedAt  time.Time     `json:"createdAt"`
	UpdatedAt  time.Time     `json:"updatedAt"`
	Reward     []struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"reward"`
	Position       int    `json:"position"`
	Published      bool   `json:"published"`
	CommunityID    string `json:"communityId"`
	SubmissionType string `json:"submissionType"`
	Condition      []struct {
		Type     string `json:"type"`
		Value    string `json:"value"`
		Operator string `json:"operator"`
	} `json:"condition"`
	ValidationData struct {
		NoOfImages int `json:"noOfImages"`
	} `json:"validationData"`
	AutoValidate      bool   `json:"autoValidate"`
	ID                string `json:"id"`
	CategoryID        string `json:"categoryId"`
	ConditionOperator string `json:"conditionOperator"`
	ClaimCounter      int    `json:"claimCounter"`
	RetryAfter        int    `json:"retryAfter"`
	Description       struct {
		Type    string `json:"type"`
		Content []struct {
			Type  string `json:"type"`
			Attrs struct {
				Indent int `json:"indent"`
			} `json:"attrs"`
			Content []struct {
				Text string `json:"text"`
				Type string `json:"type"`
			} `json:"content,omitempty"`
		} `json:"content"`
	} `json:"description"`
	Tasks      []interface{} `json:"tasks"`
	V2         bool          `json:"v2"`
	ClaimLimit int           `json:"claimLimit"`
	Archived   bool          `json:"archived"`
}
