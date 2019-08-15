package awseph_test

import (
	"testing"
	"time"

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
			_ = ephemeralSession.S3BucketMustCreate(bucket)
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
			ephemeralSession.AddService(awsephservice.S3BucketCreate(sess, bucket))
			assert.Equal(t, true, awseph.MustExists(awsephservice.S3BucketExists(sess, bucket)))
		}()
		assert.Equal(t, false, awseph.MustExists(awsephservice.S3BucketExists(sess, bucket)))
	})
	t.Run("services", func(t *testing.T) {
		ls := localstack.NewLocalStack()
		awsclient.UseLocalStack(ls)
		sess := session.Must(session.NewSession())

		name := "ephemeral"
		func() {
			ephemeralSession, teardown := awseph.NewScopedSession(sess)
			defer teardown()
			{
				t.Log("S3")
				assert.Equal(t, false, awseph.MustExists(awsephservice.S3BucketExists(sess, name)))
				ephemeralSession.S3BucketMustCreate(name)
				assert.Equal(t, true, awseph.MustExists(awsephservice.S3BucketExists(sess, name)))
			}
			{
				t.Log("SQS")
				assert.Equal(t, false, awseph.MustExists(awsephservice.SQSQueueExists(sess, name)))
				ephemeralSession.SQSQueueMustCreate(name)
				assert.Equal(t, true, awseph.MustExists(awsephservice.SQSQueueExists(sess, name)))
			}
			{
				t.Log("DynamoDB")
				type Dynamo struct {
					Id   int64     `dynamo:",hash"`
					Date time.Time `dynamo:",range"`
				}
				assert.Equal(t, false, awseph.MustExists(awsephservice.DynamoDBTableExists(sess, name)))
				ephemeralSession.DynamoDBTableMustCreateByDynamo(name, Dynamo{})
				assert.Equal(t, true, awseph.MustExists(awsephservice.DynamoDBTableExists(sess, name)))
			}
		}()
		assert.Equal(t, false, awseph.MustExists(awsephservice.S3BucketExists(sess, name)))
		assert.Equal(t, false, awseph.MustExists(awsephservice.SQSQueueExists(sess, name)))
		assert.Equal(t, false, awseph.MustExists(awsephservice.DynamoDBTableExists(sess, name)))

	})
}
