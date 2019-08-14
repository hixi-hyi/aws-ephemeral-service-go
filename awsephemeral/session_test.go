package awsephemeral_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/aws-client-go/awsclient"
	"github.com/hixi-hyi/aws-ephemeral-service-go/awsephemeral"
	"github.com/hixi-hyi/aws-ephemeral-service-go/awsephemeral/mini"
	"github.com/hixi-hyi/localstack-go/localstack"
	"github.com/stretchr/testify/assert"
)

func TestSession(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		awsclient.UseLocalStack(ls)
		sess := session.Must(session.NewSession())

		bucket := "ephemeral"

		assert.Equal(t, false, awsephemeral.MustExists(mini.S3BucketExists(sess, bucket)))
		func() {
			ephemeralSession, teardown := awsephemeral.New()
			defer teardown()
			ephemeralSession.Add(awsephemeral.MustCreate(mini.S3BucketCreate(sess, bucket)))
			assert.Equal(t, true, awsephemeral.MustExists(mini.S3BucketExists(sess, bucket)))
		}()
		assert.Equal(t, false, awsephemeral.MustExists(mini.S3BucketExists(sess, bucket)))
	})
}
