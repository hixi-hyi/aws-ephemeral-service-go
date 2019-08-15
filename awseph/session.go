package awseph

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
	awsephservice "github.com/hixi-hyi/aws-ephemeral-service-go/awseph/service"
)

type Session struct {
	AwsSession *session.Session
	Defers     []func() error
}

func NewSession(sess *session.Session) *Session {
	return &Session{
		AwsSession: sess,
	}
}

func NewScopedSession(sess *session.Session) (*Session, func()) {
	m := NewSession(sess)
	return m, m.Defer()
}

func (m *Session) Defer() func() {
	return func() {
		m.Teardown()
	}
}

func (m *Session) Teardown() {
	var errs []error
	for _, f := range m.Defers {
		err := f()
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		panic(fmt.Sprintf("%#v", errs))
	}
}

func (m *Session) AddService(i interface{}, f func() error, err error) interface{} {

	if err != nil {
		panic(err)
	}
	m.Defers = append(m.Defers, f)
	return i
}

func (m *Session) S3BucketMustCreate(name string) *s3.CreateBucketOutput {
	ret, f, err := awsephservice.S3BucketCreate(m.AwsSession, name)
	m.AddService(nil, f, err)
	return ret
}

func (m *Session) SQSQueueMustCreate(name string) *sqs.CreateQueueOutput {
	ret, f, err := awsephservice.SQSQueueCreate(m.AwsSession, name)
	m.AddService(nil, f, err)
	return ret
}

func (m *Session) DynamoDBTableMustCreateByDynamo(name string, from interface{}) {
	_, f, err := awsephservice.DynamoDBTableCreateByDynamo(m.AwsSession, name, from)
	m.AddService(nil, f, err)
	return
}
