package awsephservice_test

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/aws-client-go/awsclient"
	awsephservice "github.com/hixi-hyi/aws-ephemeral-service-go/awseph/service"
	"github.com/hixi-hyi/localstack-go/localstack"
)

type Table struct {
	Id   int64     `dynamo:",hash"`
	Date time.Time `dynamo:",range"`
}

func TestDynamoDBQueueCreate(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		awsclient.UseLocalStack(ls)
		sess := session.Must(session.NewSession())
		createFunc := awsephservice.DynamoDBTableCreateByDynamo
		existsFunc := awsephservice.DynamoDBTableExists

		name := "awsephemeral"

		if exists, err := existsFunc(sess, name); exists == true && err != nil {
			t.Errorf("service already exists")
		}
		_, teardown, err := createFunc(sess, name, Table{})
		if err != nil {
			t.Fatal(err)
		}
		if exists, err := existsFunc(sess, name); exists == false && err != nil {
			t.Errorf("service not exists")
		}
		teardown()
		if exists, err := existsFunc(sess, name); exists == true && err != nil {
			t.Errorf("service could not delete")
		}
	})
}
