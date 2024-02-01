package zealygo

import (
	"encoding/json"

	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/vottun-com/utils/errors"
	"github.com/vottunio/log"
)

type ZealySdk struct {
	apiKey    string
	subdomain string
}

func New(apiKey, subdomain string) *ZealySdk {
	return &ZealySdk{apiKey: apiKey, subdomain: subdomain}
}

// func (z *ZealySdk) GetUserBySocialId(email, discordId, twitterId, discordHandle, ethAddress *string) (interface{}, error) {
// 	var builder strings.Builder
// 	params := make([]string, 0)

// 	if email == nil && discordId == nil && twitterId == nil && discordHandle == nil && ethAddress == nil {
// 		return nil, errors.New(ErrorIncorrectParamas, "At least one of the parameters has to be ")
// 	}

// 	builder.WriteString(fmt.Sprintf(ApiUrl, z.subdomain) + "users?")

// 	if email != nil{
// 		params = append(params, *email)
// 	}
// 	if discordId != nil{
// 		params = append(params, *discordId)
// 	}
// 	if twitterId != nil{
// 		params = append(params, *twitterId)
// 	}
// 	if discordHandle != nil{
// 		params = append(params, *discordHandle)
// 	}
// 	if ethAddress != nil{
// 		params = append(params, *ethAddress)
// 	}

// 	for i:=0;i<len(params);i++{
// 		if i != len(params) - 1{

// 		}

// 	}

// 	endpoint := builder.String()

// 	return makeReqApi(z.apiKey, endpoint, METHOD_GET)

// }

func (z *ZealySdk) GetUserByEmail(email string) (*UserZealy, error) {
	endpoint := fmt.Sprintf(ApiUrl, z.subdomain) + "users?email=" + email
	userZealy := UserZealy{}
	err := makeReqApi(z.apiKey, endpoint, METHOD_GET, &userZealy)
	if err != nil {
		return nil, err
	}
	return &userZealy, nil
}

func (z *ZealySdk) GetCommunityQuests() (*[]ActivityZealy, error) {
	endpoint := fmt.Sprintf(ApiUrl, z.subdomain) + "quests"
	activitiesZealy := []ActivityZealy{}
	err := makeReqApi(z.apiKey, endpoint, METHOD_GET, &activitiesZealy)
	if err != nil {
		return nil, err
	}
	return &activitiesZealy, nil

}

func (z *ZealySdk) GetQuestById(questId string) (*ActivityZealy, error) {
	endpoint := fmt.Sprintf(ApiUrl, z.subdomain) + "quests"
	activitiesZealy := []ActivityZealy{}
	err := makeReqApi(z.apiKey, endpoint, METHOD_GET, &activitiesZealy)
	if err != nil {
		return nil, err
	}

	for _, actZealy := range activitiesZealy {
		if actZealy.ID == questId {
			return &actZealy, nil
		}
	}

	return nil, errors.New(ErrorActivityNotFound, "Activity does not exist")

}

func (z *ZealySdk) GetCommunityClaimesByUserId(status, userId string) (*ActivityAnswer, error) {

	endpoint := fmt.Sprintf(ApiUrl, z.subdomain) + "claimed-quests?status=" + status + "&user_id=" + userId
	activitiesZealy := ActivityAnswer{}
	err := makeReqApi(z.apiKey, endpoint, METHOD_GET, &activitiesZealy)
	if err != nil {
		return nil, err
	}
	return &activitiesZealy, nil

}

func (z *ZealySdk) GetCommunityClaimsByQuestId(status, questId string) (*ActivityAnswer, error) {

	endpoint := fmt.Sprintf(ApiUrl, z.subdomain) + "claimed-quests?status=" + status + "&quest_id=" + questId
	activitiesZealy := ActivityAnswer{}
	err := makeReqApi(z.apiKey, endpoint, METHOD_GET, &activitiesZealy)
	if err != nil {
		return nil, err
	}
	return &activitiesZealy, nil

}

func makeReqApi(apiKey, endpoint, method string, data interface{}) error {
	var req *http.Request
	var res *http.Response
	var statuscode int = 0

	var err error

	// builder.WriteString(endpoint)

	// if data != nil {
	// 	builder.WriteString("?")
	// 	refData := reflect.ValueOf(data)
	// 	s := refData.Elem()
	// 	fmt.Printf("s.Kind(): %v\n", s.Kind())
	// 	for i := 0; i < refData.NumField(); i++ {
	// 		field := refData.Type().Field(i)
	// 		fmt.Printf("field.Name: %v\n", field.Name)
	// 		fieldValue := refData.Field(i)

	// 		if fieldValue.Interface() != nil {
	// 			builder.WriteString(field.Name)
	// 			builder.WriteString("=")
	// 			builder.WriteString(fmt.Sprintf("%v", fieldValue.Interface()))
	// 		}

	// 	}
	// 	endpoint = builder.String()
	// }

	if req, err = http.NewRequest(method, endpoint, nil); err == nil {
		setReqHeaders(req, apiKey, method)
		client := &http.Client{
			Timeout: 30 * time.Second,
		}

		res, err = client.Do(req)
		if err == nil {
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			statuscode = res.StatusCode
			log.Tracef("Received statuscode %d", statuscode)
			switch statuscode {
			case http.StatusOK, http.StatusCreated, http.StatusAccepted:

				err = json.Unmarshal(body, &data)
				if err != nil {
					log.Printf("Error unmarshaling token information received from api: %+v", err)
					return errors.New(ErrorParsingJson, fmt.Sprintf("Error unmarshaling token information received from api: %+v", err))
				}

				return nil

			case http.StatusUnauthorized:
				return errors.New(ErrorUnauthorized, "The token used in not authorized to perform the requested operation")

			default:
				errorMsg := ErrorDTO{}
				err := json.Unmarshal(body, &errorMsg)
				if err != nil {
					log.Printf("Error unmarshaling token information received from api: %+v", err)
					return fmt.Errorf(ErrorHttpStatus, statuscode)
				}
				return errors.New(errorMsg.Code, errorMsg.Message)
			}
		} else {
			log.Printf("error executing request with error %+v", err)
			return err
		}
	} else {
		log.Printf("error creating request to send to server %+v", err)
		return err
	}
}

func setReqHeaders(req *http.Request, apiKey, method string) {
	req.Header.Set(X_API_KEY, apiKey)
	if method == METHOD_POST {
		req.Header.Set(CONTENT_TYPE, MIME_TYPE_JSON)
	}
}
