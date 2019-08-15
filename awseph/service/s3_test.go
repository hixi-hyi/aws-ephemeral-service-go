package awsephservice_test

import (
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hixi-hyi/aws-client-go/awsclient"
	awsephservice "github.com/hixi-hyi/aws-ephemeral-service-go/awseph/service"
	"github.com/hixi-hyi/localstack-go/localstack"
)

func TestS3QueueCreate(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		awsclient.UseLocalStack(ls)
		sess := session.Must(session.NewSession())
		createFunc := awsephservice.S3BucketCreate
		existsFunc := awsephservice.S3BucketExists

		name := "awsephemeral"

		if exists, err := existsFunc(sess, name); exists == true && err != nil {
			t.Errorf("service already exists. %v", err)
		}
		_, teardown, err := createFunc(sess, name)
		if err != nil {
			t.Fatal(err)
		}
		if exists, err := existsFunc(sess, name); exists == false && err != nil {
			t.Errorf("service not exists. %v", err)
		}
		if err := teardown(); err != nil {
			t.Errorf("count not do teardown. %v", err)
		}
		if exists, err := existsFunc(sess, name); exists == true && err != nil {
			t.Errorf("service could not delete. %v", err)
		}
	})
	t.Run("S3BucketForceDelete", func(t *testing.T) {
		awsephservice.S3BucketForceDelete = true
		ls := localstack.NewLocalStack()
		awsclient.UseLocalStack(ls)
		sess := session.Must(session.NewSession())
		createFunc := awsephservice.S3BucketCreate
		existsFunc := awsephservice.S3BucketExists

		name := "awsephemeral2"

		if exists, err := existsFunc(sess, name); exists == true && err != nil {
			t.Errorf("service already exists. %v", err)
		}
		_, teardown, err := createFunc(sess, name)
		if err != nil {
			t.Fatal(err)
		}
		if exists, err := existsFunc(sess, name); exists == false && err != nil {
			t.Errorf("service not exists. %v", err)
		}
		{
			t.Log("Put S3 Object")
			svc := awsclient.S3(sess)

			in := &s3.PutObjectInput{}
			in.SetBody(strings.NewReader("testtest"))
			in.SetBucket(name)
			in.SetKey("test.html")

			_, err := svc.PutObject(in)
			if err != nil {
				panic(err)
			}
		}
		if err := teardown(); err != nil {
			t.Errorf("count not do teardown. %v", err)
		}
		if exists, err := existsFunc(sess, name); exists == true && err != nil {
			t.Errorf("service could not delete. %v", err)
		}
	})

}
