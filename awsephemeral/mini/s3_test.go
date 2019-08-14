package mini_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/aws-client-go/awsclient"
	"github.com/hixi-hyi/aws-ephemeral-service-go/awsephemeral/mini"
	"github.com/hixi-hyi/localstack-go/localstack"
)

func TestCreateS3Bucket(t *testing.T) {
	t.Run("mini", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		awsclient.UseLocalStack(ls)
		sess := session.Must(session.NewSession())

		bucket := "awsephemeral"

		if exists, _ := mini.S3BucketExists(sess, bucket); exists == true {
			t.Errorf("bucket already exists")
		}
		teardown, err := mini.S3BucketCreate(sess, bucket)
		if err != nil {
			t.Fatal(err)
		}
		if exists, _ := mini.S3BucketExists(sess, bucket); exists == false {
			t.Errorf("bucket not exists")
		}
		teardown()
		if exists, _ := mini.S3BucketExists(sess, bucket); exists == true {
			t.Errorf("bucket could not delete")
		}
	})
}
