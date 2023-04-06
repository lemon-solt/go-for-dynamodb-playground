package client

import (
	"fmt"
	"log"
	"playground/constants"
	"time"
)

func SimplePut() {
	log.Println("simplePut start at: ", time.Now())
	// constants.TableNameTableName := "<テーブル名>"

	table := ConnectionTable(constants.TableName)

	p := &PlaygroundRow{Id: "sample_user_1", Suffix: 20230101, Name: "user1"}

	log.Println(p)

	if err := table.Put(p).Run(); err != nil {
		log.Fatalln("put error: ", err)
	}
}

func SimppleBulkPut() {

	log.Println("simple bulk start at: ", time.Now())
	// constants. := "<テーブル名>"

	table := ConnectionTable(constants.TableName)

	itemList := make([]interface{}, 0)

	for i, v := range []string{"foo", "bar", "baz"} {
		fmt.Println(i)
		itemList = append(itemList, PlaygroundRow{Id: v, Suffix: i + 1, Name: v + fmt.Sprintf("user_%v", v)})
	}

	log.Println("itemlist: ", len(itemList))

	bw := table.Batch()
	wr, err := bw.Write().Put(itemList...).Run()

	if err != nil {
		log.Fatalln("err: ", err)
	}

	log.Println("success: ", wr)
}
