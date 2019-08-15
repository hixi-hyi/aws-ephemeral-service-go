package awsephservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hixi-hyi/aws-client-go/awsclient"
)

var S3BucketForceDelete = false

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
		if S3BucketForceDelete {
			svc := awsclient.S3(sess)
			input := &s3.ListObjectsInput{}
			input.SetBucket(bucket)
			ret, err := svc.ListObjects(input)
			if err != nil {
				return err
			}
			for _, item := range ret.Contents {
				input := &s3.DeleteObjectInput{}
				input.SetBucket(bucket)
				input.SetKey(aws.StringValue(item.Key))
				_, err := svc.DeleteObject(input)
				if err != nil {
					return err
				}
			}

			//iter := s3manager.NewDeleteListIterator(svc, input)
			//if err := s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter); err != nil {
			//	fmt.Printf("Unable to delete objects from bucket %q, %v", bucket, err)
			//	return err
			//}
		}
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
