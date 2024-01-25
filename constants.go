package zealygo

const (
	ApiUrl string = "https://api-v1.zealy.io/communities/%s/"
)

const (
	CONTENT_TYPE   string = "Content-Type"
	X_API_KEY      string = "x-api-key"
	MIME_TYPE_JSON string = "application/json; charset=UTF-8"
	METHOD_GET            = "GET"
	METHOD_POST           = "POST"
)

const (
	ErrorParsingJson      string = "ERROR_PARSING_JSON"
	ErrorUnauthorized     string = "ERROR_UNAUTHORIZED"
	ErrorHttpStatus       string = "ERROR_HTTP_STATUS_%d"
	ErrorIncorrectParamas string = "ERROR_INCORRECT_PARAMS"
	//ErrorApiWrapperUrlNotSet string = "ERROR_API_WRAPPER_URL_NOT_SET"
)
