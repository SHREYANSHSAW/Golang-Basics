package main

import (
	"fmt"
	"github.com/go-lambda-function"
)

type MyEvent struct{
	Name string `json:"what is your name"`
	Age int `json:"what is your age"`
}

type MyResponse struct{
	Message string `json:"Answer"`
}

func HandleLambdaEvent(event MyEvent)(MyResponse err){
	return MyResponse{Message: fmt.Sprintf("%s is %d is year old !" , event.Name, event.Age)}
}

func main(){
	lambda.Start(HandleLambdaEvent)
}