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

func (z *ZealySdk) GetCommunityClaimesByUserId(status, userId string, page, limit *string) (*ActivityAnswer, error) {

	endpoint := fmt.Sprintf(ApiUrl, z.subdomain) + "claimed-quests?status=" + status + "&user_id=" + userId
	if page != nil && limit != nil {
		endpoint = fmt.Sprintf(ApiUrl, z.subdomain) + "claimed-quests?status=" + status + "&user_id=" + userId + "&page=" + *page + "&limit=" + *limit
	}
	activitiesZealy := ActivityAnswer{}
	err := makeReqApi(z.apiKey, endpoint, METHOD_GET, &activitiesZealy)
	if err != nil {
		return nil, err
	}
	return &activitiesZealy, nil

}

func (z *ZealySdk) GetCommunityClaimsByQuestId(status, questId string, page, limit *string) (*ActivityAnswer, error) {

	endpoint := fmt.Sprintf(ApiUrl, z.subdomain) + "claimed-quests?status=" + status + "&quest_id=" + questId
	if page != nil && limit != nil {
		endpoint = fmt.Sprintf(ApiUrl, z.subdomain) + "claimed-quests?status=" + status + "&quest_id=" + questId + "&page=" + *page + "&limit=" + *limit
	}
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
				var err error
				if isJSON(body) {
					err = json.Unmarshal(body, &errorMsg)
					if err != nil {
						log.Printf("Error unmarshaling token information received from api: %+v", err)
						return fmt.Errorf(ErrorHttpStatus, statuscode)
					}
				} else {
					errorMsg.Message = string(body)
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

func isJSON(data []byte) bool {
	var js map[string]interface{}
	return json.Unmarshal(data, &js) == nil
}
