package client

import (
	"fmt"
	"log"
)

func SimpleScanRead() {
	// Scanでdbを読み込むサンプル
	sampleTableName := "<テーブル名>"
	table := ConnectionTable(sampleTableName)

	var rows []PlaygroundRow
	snapshot := table.Scan().All(&rows)

	if snapshot != nil {
		log.Fatalln("snapshot error: ", snapshot)
	}

	log.Println("simpleScanRead logs: ", rows)
}

func SimpleQueryRead() {
	// Queryでdbを読み込むサンプル
	sampleTableName := "<テーブル名>"
	table := ConnectionTable(sampleTableName)

	var rows []PlaygroundRow

	targetColumn, targetValue := "<指定カラム>", "<指定の値>"
	snapshot := table.Get(targetColumn, targetValue).All(&rows)

	if snapshot != nil {
		log.Fatalln("snapshot error: ", snapshot)
	}

	fmt.Println(rows)
}
