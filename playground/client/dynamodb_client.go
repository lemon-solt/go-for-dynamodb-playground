package client

import (
	"playground/repository"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type PlaygroundRow struct {
	Id     string `dynamo:"id"`
	Suffix int    `dynamo:"suffix"`
	Name   string `dynamo:"name"`
}

func ConnectionTable(tableName string) (table dynamo.Table) {
	settings := repository.EnvSetting
	c := credentials.NewStaticCredentials(settings.AwsAccessKey, settings.AwsSeacretKey, "")
	db := dynamo.New(session.New(), &aws.Config{Credentials: c, Region: aws.String(settings.Region)})
	table = db.Table(tableName)
	return
}
