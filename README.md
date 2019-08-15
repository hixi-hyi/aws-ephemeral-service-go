## aws-ephemeral-service-go (awseph)

awseph is create aws services for testing easily.


### Example

```go

func TestMain() {
    ls := localstack.NewLocalStack()
    awsclient.UseLocalStack(ls)
	sess := session.Must(session.NewSession())

    name := "ephemeral"

    ephSession, teardown := awseph.NewScopedSession(sess)
    defer teardown()

    awsephservice.S3BucketForceDelete = true
    ephSession.S3BucketMustCreate(name)
    // write the test code for s3

    ephSession.SQSQueueMustCreate(name)
    // write the test code for sqs
}
```

### Services
* [session](./awseph/session.go)
* [service](./awseph/service)
