package awseph_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hixi-hyi/aws-client-go/awsclient"
	"github.com/hixi-hyi/aws-ephemeral-service-go/awseph"
	awsephservice "github.com/hixi-hyi/aws-ephemeral-service-go/awseph/service"
	"github.com/hixi-hyi/localstack-go/localstack"
	"github.com/stretchr/testify/assert"
)

func TestSession(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		awsclient.UseLocalStack(ls)
		sess := session.Must(session.NewSession())

		bucket := "ephemeral"

		assert.Equal(t, false, awseph.MustExists(awsephservice.S3BucketExists(sess, bucket)))
		func() {
			ephemeralSession, teardown := awseph.NewScopedSession(sess)
			defer teardown()
            ephemeralSession.S3BucketMustCreate(bucket)
			assert.Equal(t, true, awseph.MustExists(awsephservice.S3BucketExists(sess, bucket)))
		}()
		assert.Equal(t, false, awseph.MustExists(awsephservice.S3BucketExists(sess, bucket)))
	})
	t.Run("add service", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		awsclient.UseLocalStack(ls)
		sess := session.Must(session.NewSession())

		bucket := "ephemeral"

		assert.Equal(t, false, awseph.MustExists(awsephservice.S3BucketExists(sess, bucket)))
		func() {
			ephemeralSession, teardown := awseph.NewScopedSession(sess)
			defer teardown()
			ephemeralSession.AddService(awseph.MustCreate(awsephservice.S3BucketCreate(sess, bucket)))
			assert.Equal(t, true, awseph.MustExists(awsephservice.S3BucketExists(sess, bucket)))
		}()
		assert.Equal(t, false, awseph.MustExists(awsephservice.S3BucketExists(sess, bucket)))
	})
}
