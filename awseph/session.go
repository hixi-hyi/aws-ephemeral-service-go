package awseph

import (
    "github.com/aws/aws-sdk-go/aws/session"
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
		panic(errs)
	}
}

func (m *Session) AddService(f func() error) {
	m.Defers = append(m.Defers, f)
}

func (m *Session) S3BucketMustCreate(bucket string) {
	m.AddService(MustCreate(awsephservice.S3BucketCreate(m.AwsSession, bucket)))
}
