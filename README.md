## aws-ephemeral-service-go (awseph)

awseph is create aws services for testing easily.


### Example

```go

func TestMain() {
    ls := localstack.NewLocalStack()
    awsclient.UseLocalStack(ls)
	sess := session.Must(session.NewSession())

    name := "ephemeral"

	ephemeralSession, teardown := awseph.NewScopedSession(sess)
    defer teardown()

    ephemeralSession.S3BucketMustCreate(name)
    // write the test code for s3

    ephemeralSession.SQSQueueMustCreate(name)
    // write the test code for sqs
}
```

### Services
* [session](./blob/master/awseph/session.go)
* [service](./tree/master/awseph/service)
