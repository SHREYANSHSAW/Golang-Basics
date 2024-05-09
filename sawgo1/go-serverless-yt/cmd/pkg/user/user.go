package user

import(
	
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/dynmaodb"
	"github.com/aws/aws-sdk-go/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

)

var(
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord = "failed to fetch record"
	ErrorInvalidUserData = "invalid user data"
	ErrorInvalidEmail = "invalid email"
	ErrorCouldNotMarshalItem = "could not marshal item"
	ErrorCouldNotDeleteItem = "cold delete item"
	ErrorcouldnotDynamoPutitem = "could not dynamo put item"
    ErrorUserAlreadyExists = "user.User already exists"
	ErrorUuserDoesNotExist = "user.user does not exists"
)

type User struct{
	Emai  string `json:"email"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

func FetchUser(email, tableName string, dynaclient dynamodbiface.DynamiDBABI)(*user, error){

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email":{
				S: aws.String(email)
			}
		},
		TableName: aws.String(tableNmae)
	}

	result, err := dynaClient.GetItem(input)
	if err!= nil{
		return nil, errors.New(ErrorfailedToFatchRecord)
	}

}

func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*[]User, error){
	input := &dynamodb.ScanInput{
		TableName := aws.String(tableName),
	}
	result, err := dynaClient.Scan(input)
	if err != nil{
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]user)
	err = dynamodbattribute.unmarshalMap(result.Item, item)
	return item, nil

}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient Dynamodbiface.DynamoDBABI)(*User, error)
{
	var u user

	if err := json.Unmarshal([]byte(req.body), &u); err !=nil{
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(u.Email){
		return nil, errors.New(ErrorInvalidUserData)
	}
	currentUser, _ := FetchUser(u.Email, tableName, dynaClient)
    if currentUser != nil && len(currentUser.Email) !=0 {
		return nil, errors.New(ErrorUserAlreadyExists)
	} 
	av, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	} 
	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(tableName)
	}

	_, err = dynaClient.PutItem(input)
	if err != nil{
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &u, nil
}

func UpdateUser(req events.APIGatwayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*User, error)
{
	var u User
	if err := json.Unmarshal([]byte(req.Body), &u); err !=nil{
		return nil, errors.New(ErrorinvalidEmail)
	}
   currentUser, _ := FetchUser(u.Email, tableName, dynaClient)
   if currentUser != nil && len (currentUser.Email)==0{
	return nil, errors.New(ErrorUserDoesNotExist)
   }

   av, err := dynamodbattribute.Marshal(u)
   if err != nil {
	return nil, errors.New(ErrorCouldNotmarshalItem)
   }

   input := &dynamodb.PutItemInput{
	Item: av,
	TableName: aws.String(tableName),
   }

   _, err = dynaClient.PutItem(input)
   if err != nil{
	return nil, errors.New(ErrorCouldNotDynamoPutItem)
   }
   return &u, nil
	

}
func DeleteUser() error{
	
}