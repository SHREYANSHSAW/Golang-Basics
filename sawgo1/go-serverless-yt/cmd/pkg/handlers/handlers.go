package handlers

import(
	"github.com/go-serverless-yt/pkg/user"
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-skd-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

)
 var ErrorMetodNotAllowed = "method not allowed"

type ErrorBody struct{
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIgatwayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIGatewayProxyResponse, error
){
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err!= nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusBadRequest, ErrorBody(aws.String(err.Error())))
	}
	return apiResponse(http.StatusOK, result)

}

result, err := user.FetchUsers(tableName, dynaClient)
if err != nil{
	return apiResponse(http.StatusBadRequest, ErrorBody{
		aws.String(err.Error()),
	})
}
return apiResponse(http.StatusOK, result)

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*events.APIgatwayProxyResponse, error
){
	result, err := user.CreateUser(req, tableName, dynaClient )
	if err != nil{
		return apiResponse(http.StatusBadRequest, errorBody{
           aws.String(err.Error()),
		})
		return apiResponse(http.StatusCreated, result)
	}

}

func UpdateUser(req, events.APIGatwayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*events.APIGatwayProxyResponse, error){
reult, err := user.UpdateUser(req, tableNmae, dynaClient)
if err != nil{
	return apiResponse(http.StatusBadRequest, ErrorBody{)
		aws.String(err.Error())
	}
}


func DeleteUser(){

}

func UnhandledMethod()(){
	 
}