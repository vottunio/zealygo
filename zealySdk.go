package zealygo

import (
	"encoding/json"

	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
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

func (z *ZealySdk) GetUserBySocialId(email, discordId, twitterId, discordHandle, ethAddress *string) (interface{}, error) {
	userReq := userRequestDto{email: email, discordId: discordId, twitterId: twitterId, discordHandle: discordHandle, ethAddress: ethAddress}
	return makeReqApi(z.apiKey, fmt.Sprintf(ApiUrl, z.subdomain)+"users", METHOD_GET, userReq)

}

func (z *ZealySdk) GetCommunityQuests() {

}

func (z *ZealySdk) GetCommunityClaimedQuests() {

}

func makeReqApi(apiKey, endpoint, method string, data interface{}) (interface{}, error) {
	var req *http.Request
	var res *http.Response
	var statuscode int = 0
	var answer interface{}
	var builder strings.Builder
	var err error

	builder.WriteString(endpoint)

	if data != nil {
		builder.WriteString("?")
		refData := reflect.ValueOf(data)
		for i := 0; i < refData.NumField(); i++ {
			field := refData.Type().Field(i)
			fieldValue := refData.Field(i)
			if fieldValue.Interface() != nil {
				builder.WriteString(field.Name)
				builder.WriteString("=")
				builder.WriteString(fmt.Sprintf("%v", fieldValue.Interface()))
			}

		}
		endpoint = builder.String()
	}

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

				err = json.Unmarshal(body, &answer)
				if err != nil {
					log.Printf("Error unmarshaling token information received from api: %+v", err)
					return nil, errors.New(ErrorParsingJson, fmt.Sprintf("Error unmarshaling token information received from api: %+v", err))
				}

				return answer, nil

			case http.StatusUnauthorized:
				return nil, errors.New(ErrorUnauthorized, "The token used in not authorized to perform the requested operation")

			default:
				errorMsg := ErrorDTO{}
				err := json.Unmarshal(body, &errorMsg)
				if err != nil {
					log.Printf("Error unmarshaling token information received from api: %+v", err)
					return nil, fmt.Errorf(ErrorHttpStatus, statuscode)
				}
				return nil, errors.New(errorMsg.Code, errorMsg.Message)
			}
		} else {
			log.Printf("error executing request with error %+v", err)
			return nil, err
		}
	} else {
		log.Printf("error creating request to send to server %+v", err)
		return nil, err
	}
}

func setReqHeaders(req *http.Request, apiKey, method string) {
	req.Header.Set(X_API_KEY, apiKey)
	if method == METHOD_POST {
		req.Header.Set(CONTENT_TYPE, MIME_TYPE_JSON)
	}
}
