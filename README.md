# MySQL to ElasticSearch in Golang

This is a small Golang tool for loading MySQL data to ElasticSearch for improve the search in the database.

## Usage
```
package main

import (
	"github.com/pumpkinseed/mysqltoes"
)

func main() {
	config := mysqltoes.SetConfig(
		DATABASE_CONNECTION,
		ELASTICSEARCH_CONNECTION,
		DATABASE_QUERY,
		ELASTICSEARCH_INDEX,
		ELASTICSEARCH_TYPE)
	mysqltoes.LoadMysqlToElasticsearch(config)
}
```
More details in the `./example`

## Dependencies

```
go get github.com/go-sql-driver/mysql
go get gopkg.in/olivere/elastic.v3
```

## Install

```
go get github.com/PumpkinSeed/golang-mysqltoes
```
