package handlers

import(
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"

)

func apiResponse(status int, body interface{})(*events.APIGatwayProxyResponse){
	resp := resp.APIGatwayProxyResponse{Headers:map[string]string["contentent type":"application/json"]}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}