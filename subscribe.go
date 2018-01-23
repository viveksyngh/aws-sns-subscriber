package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
)

const httpEndpoint = "https://6303c682.ngrok.io"
const httpProtocol = "https"
const snsTopicARN = "arn:aws:sns:ap-south-1:725344396561:open-faas-test"
const awsRegion = "ap-south-1"

//subscribe sends a subscribe request to SNS topic and initiates the subscritption process
func subscribe(endPoint string, protocol string, topicARN string) {

	input := &sns.SubscribeInput{
		Endpoint: &endPoint,
		Protocol: &protocol,
		TopicArn: &topicARN,
	}

	sess, err := session.NewSession(&aws.Config{Region: aws.String(getenv("AWS_REGION", awsRegion))})
	if err != nil {
		fmt.Printf("Unable to initiate session")
		return
	}
	svc := sns.New(sess)

	out, err := svc.Subscribe(input)
	if err != nil {
		fmt.Println("Unable to Subscribe")
		return
	}
	fmt.Printf(*out.SubscriptionArn)
}

//getenv returns env varibale if set otherwise returns default value
func getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func main() {
	httpEndpoint := getenv("HTTP_ENDPOINT", httpEndpoint)
	httpProtocol := getenv("HTTP_PROTOCOL", httpProtocol)
	snsTopicARN := getenv("SNS_TOPIC_ARN", snsTopicARN)
	subscribe(httpEndpoint, httpProtocol, snsTopicARN)
}
