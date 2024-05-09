package main

import(
	"os"
	"github.com/aws/aws-skd-go/aws/session"
)

var(
	dynaClient dynamodbiface.DynamoDBAPI
)

func main(){
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},)

		if err != nil{
			return
		}
		dynaClient = dynamodb.New(awsSession)
		lambda.Start(handler)
}

func handler(req events.APIGatwayProxyRequest) (*event.APIGatwayProxyResponse, error){
	switch req.HTTPMethod{
	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynaClient)
	case "put":
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "Delete":
		return handlers.DeleteUser(req, tableName, dynaClient)
	}
default:
	return handlers.UnhandleMethod()
}

const tableName = "LambdaInGoUser"

