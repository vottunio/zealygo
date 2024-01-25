package zealygo

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
