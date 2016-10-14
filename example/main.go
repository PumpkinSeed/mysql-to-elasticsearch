package main

import (
	"github.com/pumpkinseed/mysqltoes"
)

func main() {
	config := mysqltoes.SetConfig(
		"{USER}:{PASSWORD}@tcp({ADDRESS}:{PORT})/{DATABASE}",
		"{ADDRESS}:{PORT}",
		"{QUERY}",
		"{INDEX}",
		"{TYPE}")
	mysqltoes.LoadMysqlToElasticsearch(config)
}
