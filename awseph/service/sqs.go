package awsephservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/hixi-hyi/aws-client-go/awsclient"
)

func SQSQueueCreate(sess *session.Session, queue string) (*sqs.CreateQueueOutput, func() error, error) {
	input := &sqs.CreateQueueInput{}
	input.SetQueueName(queue)
	svc := awsclient.SQS(sess)
	ret, err := svc.CreateQueue(input)
	if err != nil {
		return nil, nil, err
	}
	f := func() error {
		input := &sqs.DeleteQueueInput{}
		input.SetQueueUrl(aws.StringValue(ret.QueueUrl))
		svc := awsclient.SQS(sess)
		_, err := svc.DeleteQueue(input)
		return err
	}
	return ret, f, nil
}

func SQSQueueExists(sess *session.Session, queue string) (bool, error) {
	input := &sqs.GetQueueUrlInput{}
	input.SetQueueName(queue)
	_, err := awsclient.SQS(sess).GetQueueUrl(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sqs.ErrCodeQueueDoesNotExist:
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
