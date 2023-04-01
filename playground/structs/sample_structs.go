package structs

type PlaygroundRow struct {
	Id     string `dynamo:"id"`
	Suffix int    `dynamo:"suffix"`
	Name   string `dynamo:"name"`
}
