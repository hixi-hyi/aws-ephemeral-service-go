package awsephservice

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hixi-hyi/aws-client-go/awsclient"
)

// func S3BucketCreateFull(sess *session.Session, input *s3.BucketCreateInput) (func() error, error)

/*
Notice: the method can not delete bucket if bucket has objects.
*/
func S3BucketCreate(sess *session.Session, bucket string) (*s3.CreateBucketOutput, func() error, error) {
	input := &s3.CreateBucketInput{}
	input.SetBucket(bucket)
	svc := awsclient.S3(sess)
	ret, err := svc.CreateBucket(input)
	if err != nil {
		return nil, nil, err
	}
	f := func() error {
		input := &s3.DeleteBucketInput{}
		input.SetBucket(bucket)
		svc := awsclient.S3(sess)
		_, err := svc.DeleteBucket(input)
		return err
	}
	return ret, f, nil
}

func S3BucketExists(sess *session.Session, bucket string) (bool, error) {
	input := &s3.HeadBucketInput{}
	input.SetBucket(bucket)
	svc := awsclient.S3(sess)
	_, err := svc.HeadBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket, "NotFound":
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
