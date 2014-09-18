package main

import (
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/dynamodb"
	"log"
	"os/user"
	"time"
)

var authCreds aws.Auth
var awsRegion aws.Region
var defaultExpiration time.Duration
var defaultAwsRegionName = "ap-southeast-1"
var dynamoDBTable dynamodb.Table
var dynamoDBServer dynamodb.Server
var dynamoDBTableName = "status_table"

func main() {
	hashKey := "java"
	rangeKey := "de1ece0a-28e9-429c-a9a5-b2fc5e2fc3f9"
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

	dynamoDBServer = dynamodb.Server{}
	dynamoDBServer.Auth = authCreds
	dynamoDBServer.Region = awsRegion
	dynamoDBTable = dynamodb.Table{}
	dynamoDBTable.Name = dynamoDBTableName
	dynamoDBTable.Server = &dynamoDBServer
	primaryKey := dynamodb.PrimaryKey{}
	primaryKey.KeyAttribute = dynamodb.NewStringAttribute("languageId", "S")
	primaryKey.RangeAttribute = dynamodb.NewStringAttribute("cuid", "S")
	attributes := make([]dynamodb.Attribute, 4)
	attributes[0] = *dynamodb.NewStringAttribute("languageId", "S")
	attributes[1] = *dynamodb.NewStringAttribute("cuid", "SK")
	attributes[2] = *dynamodb.NewStringAttribute("myKey", "1234")
	attributes[3] = *dynamodb.NewStringAttribute("another_Attr", "Sample")
	created, err := dynamoDBTable.PutItem(hashKey, rangeKey, attributes)
	if err != nil {
		log.Fatal("Error creating the Dynamo DB Table : ", err.Error())
	} else {
		log.Println("Create Item Status : ", created)
	}
	dynamoDBTable.Key = primaryKey
	key := dynamodb.Key{}
	key.HashKey = hashKey
	key.RangeKey = rangeKey
	item, err := dynamoDBTable.GetItem(&key)
	if err != nil || item == nil {
		log.Println("Error getting Item (Key = ", key, ") : ", err)
	} else {
		log.Println(" Got the Item ", item)
	}
	log.Println("Deleting the same item now !")
	deleted, err := dynamoDBTable.DeleteItem(&key)
	if err != nil {
		log.Fatal("Error deleting the item : ", err.Error())
	} else {
		log.Println("Item Delete Status := ", deleted)
	}
}
