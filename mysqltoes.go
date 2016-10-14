package mysqltoes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/olivere/elastic.v3"
)

var (
	db       *sql.DB
	rows     *sql.Rows
	esClient *elastic.Client
	c        Config
	err      error
)

// Config struct contain the user managed configuration
type Config struct {
	DatabaseConn       string
	ElasticsearchConn  string
	DatabaseQuery      string
	ElasticsearchIndex string
	ElasticsearchType  string
}

// LoadMysqlToElasticsearch is loading the query result into the database
func LoadMysqlToElasticsearch(c Config) {
	c.connectToDatabase()
	c.connectToElasticsearch()
	c.getQueryRows()

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	resultID := 0
	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)

		tmpStruct := map[string]string{}

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			tmpStruct[col] = fmt.Sprintf("%s", v)
		}
		b, _ := json.Marshal(tmpStruct)
		fmt.Println(strconv.Itoa(resultID) + ". data was loaded into the Elasticsearch")
		loadToElasticsearch(string(b), strconv.Itoa(resultID))
		resultID++
	}
}

//SetConfig is setting up the configuration variables
func SetConfig(dbconn string, esconn string, dbquery string, esindex string, estype string) Config {
	c.DatabaseConn = dbconn
	c.ElasticsearchConn = esconn
	c.DatabaseQuery = dbquery
	c.ElasticsearchIndex = esindex
	c.ElasticsearchType = estype
	return c
}

func loadToElasticsearch(json string, id string) {
	_, err = esClient.Index().
		Index(c.ElasticsearchIndex).
		Type(c.ElasticsearchType).
		Id(id).
		BodyString(json).
		Refresh(true).
		Do()
	if err != nil {
		fmt.Println(err)
	}
}

func (c Config) connectToDatabase() {
	db, err = sql.Open("mysql", c.DatabaseConn)
	if err != nil {
		fmt.Println(err)
	}
}

func (c Config) connectToElasticsearch() {
	esClient, err = elastic.NewSimpleClient(elastic.SetURL(c.ElasticsearchConn))
	if err != nil {
		fmt.Println(err)
	}
}

func (c Config) getQueryRows() {
	rows, err = db.Query(c.DatabaseQuery)
	if err != nil {
		fmt.Println(err)
	}
}
