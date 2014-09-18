package main

import (
	//  "fmt"
	"log"
	//  "time"
	//  "encoding/json"
	//  "encoding/base64"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sqs"
	"os/user"
	"time"
)

var authCreds aws.Auth
var awsRegion aws.Region
var sqsService *sqs.SQS
var sqsQueue *sqs.Queue
var defaultExpiration time.Duration
var defaultAwsRegionName = "ap-southeast-1"
var sqsQueueName = "MyQueue"

func main() {
	defaultExpiration = 10000
	usr, err := user.Current()
	homeDir := "."
	if err != nil {
		log.Println("Error gettingcurrent user ", err)
	} else {
		homeDir = usr.HomeDir
		log.Println(" User ", usr, " Home Dir := ", homeDir)
	}
	authCreds, err = aws.CredentialFileAuth(homeDir+"/.aws/credentials", "default", defaultExpiration)
	awsRegion = aws.Regions[defaultAwsRegionName]
	if err != nil {
		log.Println(" Error getting Authentication Creds : ", err)
	}

	sqsService = sqs.New(authCreds, awsRegion)
	sqsQueue, err = sqsService.GetQueue(sqsQueueName)
	if err != nil {
		log.Println("Error getting SQS Queue (", sqsQueueName, ") : ", err)
	}
	log.Println("Sending Message...")
	messageBody := "Hello World!"
	sendResponse, err := sqsQueue.SendMessage(messageBody)
	if err != nil {
		log.Fatal("Message Sending Error : ", err)
	} else {
		log.Println("Message Send Response = ", sendResponse)
	}
	log.Println("Receiving Message...")
	receiveResponse, err := sqsQueue.ReceiveMessage(1)
	if err != nil {
		log.Println("Error receiving next message : ", err)
		return
	}
	log.Println(" Receive Response : ", receiveResponse)
	if len(receiveResponse.Messages) > 0 {
		for _, msg := range receiveResponse.Messages {
			msgAttrbs := msg.MessageAttribute
			log.Println(" Message := ", msg, " Attributes = ", msgAttrbs)
		}
	} else {
		log.Println("No Messages found in the Queue")
	}
	log.Println("Deleting the received Message after processing !")
	if len(receiveResponse.Messages) > 0 {
		for _, msg := range receiveResponse.Messages {
			deleteResponse, err := sqsQueue.DeleteMessage(&msg)
			log.Println(" Delete Response := ", deleteResponse, " Error = ", err)
		}
	}
}
