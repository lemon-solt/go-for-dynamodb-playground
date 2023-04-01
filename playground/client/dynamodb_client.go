package client

import (
	"fmt"
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

func CreateSession() {
	table := ConnectionTable("")
	var rows []PlaygroundRow
	// err := table.Get("id", "1").All(&rows)
	err := table.Scan().All(&rows)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	fmt.Println(rows) // [{100 2017-06-02T09:25:56.000147134 +0000 UTC}]と出力される。

	for i := range rows {
		fmt.Println(rows[i]) // {100 2017-06-02T09:25:56.000147134 +0000 UTC}
	}
}

func ConnectionTable(tableName string) (table dynamo.Table) {
	settings := repository.EnvSetting
	c := credentials.NewStaticCredentials(settings.AwsAccessKey, settings.AwsSeacretKey, "")
	db := dynamo.New(session.New(), &aws.Config{Credentials: c, Region: aws.String(settings.Region)})
	table = db.Table(tableName)
	return
}
