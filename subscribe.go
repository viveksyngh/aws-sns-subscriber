package main

import (
	"github.com/aws/aws-sdk-go/service/sns"
 	"github.com/aws/aws-sdk-go/aws"
 	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
)
const endPoint = "https://39cb8fd6.ngrok.io"
const protocol = "https"
const topicARN = "arn:aws:sns:ap-south-1:725344396561:open-faas-test"

//subscribe sends a subscribe request to SNS topic and initiates the subscritption process
func subscribe(endPoint string, protocol string, topicARN string) {
	
	input := &sns.SubscribeInput{
		Endpoint: &endPoint,
		Protocol: &protocol,
		TopicArn: &topicARN,
	
	}

	sess, err := session.NewSession(&aws.Config{Region: aws.String("ap-south-1"),})
	if(err != nil) {
		fmt.Printf("Unable to initiate session")
		return 
	}
	svc := sns.New(sess)

	out, err := svc.Subscribe(input)
	if(err != nil){
		fmt.Println("Unable to Subscribe")
		return 
	}
	fmt.Printf(*out.SubscriptionArn)
}

func main() {
	subscribe(endPoint, protocol, topicARN)
}
