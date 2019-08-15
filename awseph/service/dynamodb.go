package awsephservice

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	"github.com/hixi-hyi/aws-client-go/awsclient"
)

func DynamoDBTableCreateByDynamo(sess *session.Session, name string, from interface{}) (interface{}, func() error, error) {

	db := dynamo.NewFromIface(awsclient.DynamoDB(sess))
	err := db.CreateTable(name, from).Run()
	if err != nil {
		return nil, nil, err
	}

	f := func() error {
		return db.Table(name).DeleteTable().Run()
	}
	return nil, f, nil
}

func DynamoDBTableExists(sess *session.Session, name string) (bool, error) {
	db := dynamo.NewFromIface(awsclient.DynamoDB(sess))
	_, err := db.Table(name).Describe().Run()
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
