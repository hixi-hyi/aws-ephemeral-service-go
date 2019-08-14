package awsephservice_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/aws-client-go/awsclient"
	awsephservice "github.com/hixi-hyi/aws-ephemeral-service-go/awseph/service"
	"github.com/hixi-hyi/localstack-go/localstack"
)

func TestCreateS3Bucket(t *testing.T) {
	t.Run("awsephservice", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		awsclient.UseLocalStack(ls)
		sess := session.Must(session.NewSession())

		bucket := "awsephemeral"

		if exists, _ := awsephservice.S3BucketExists(sess, bucket); exists == true {
			t.Errorf("bucket already exists")
		}
		teardown, err := awsephservice.S3BucketCreate(sess, bucket)
		if err != nil {
			t.Fatal(err)
		}
		if exists, _ := awsephservice.S3BucketExists(sess, bucket); exists == false {
			t.Errorf("bucket not exists")
		}
		teardown()
		if exists, _ := awsephservice.S3BucketExists(sess, bucket); exists == true {
			t.Errorf("bucket could not delete")
		}
	})
}
