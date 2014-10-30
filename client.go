package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
)

const (
	DATABASES     = "SHOW DATABASES;"
	TABLES        = "SHOW TABLES;"
	TABLE_SCHEMA  = "SHOW COLUMNS IN`%s`;"
	TABLE_INDEXES = "SHOW INDEX IN `%s`;"
)

type Client struct {
	db      *sqlx.DB
	history []string
}

type Result struct {
	Columns []string        `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}

func NewError(err error) Error {
	return Error{err.Error()}
}

func NewClient() (*Client, error) {
	db, err := sqlx.Open("mysql", getConnectionString())

	if err != nil {
		return nil, err
	}

	return &Client{db: db}, nil
}

func (client *Client) Test() error {
	return client.db.Ping()
}

func (client *Client) recordQuery(query string) {
	client.history = append(client.history, query)
}

func (client *Client) Databases() ([]string, error) {
	res, err := client.Query(DATABASES)

	if err != nil {
		return nil, err
	}

	var tables []string

	for _, row := range res.Rows {
		tables = append(tables, row[0].(string))
	}

	return tables, nil
}

func (client *Client) Tables() ([]string, error) {
	res, err := client.Query(TABLES)

	if err != nil {
		return nil, err
	}

	var tables []string

	for _, row := range res.Rows {
		tables = append(tables, row[0].(string))
	}

	return tables, nil
}

func (client *Client) Table(table string) (*Result, error) {
	return client.Query(fmt.Sprintf(TABLE_SCHEMA, table))
}

func (client *Client) TableIndexes(table string) (*Result, error) {
	res, err := client.Query(fmt.Sprintf(TABLE_INDEXES, table))

	if err != nil {
		return nil, err
	}

	return res, err
}

func (client *Client) Query(query string) (*Result, error) {
	rows, err := client.db.Queryx(query)

	client.recordQuery(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cols, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	result := Result{
		Columns: cols,
	}

	for rows.Next() {
		obj, err := rows.SliceScan()

		for i, item := range obj {
			if item == nil {
				obj[i] = nil
			} else {
				t := reflect.TypeOf(item).Kind().String()

				if t == "slice" {
					obj[i] = string(item.([]byte))
				}
			}
		}

		if err == nil {
			result.Rows = append(result.Rows, obj)
		}
	}

	return &result, nil
}

func (res *Result) Format() []map[string]interface{} {
	var items []map[string]interface{}

	for _, row := range res.Rows {
		item := make(map[string]interface{})

		for i, c := range res.Columns {
			item[c] = row[i]
		}

		items = append(items, item)
	}

	return items
}

func (res *Result) CSV() []byte {
	buff := &bytes.Buffer{}
	writer := csv.NewWriter(buff)

	writer.Write(res.Columns)

	for _, row := range res.Rows {
		record := make([]string, len(res.Columns))

		for i, item := range row {
			if item != nil {
				record[i] = fmt.Sprintf("%v", item)
			} else {
				record[i] = ""
			}
		}

		err := writer.Write(record)

		if err != nil {
			fmt.Println(err)
			break
		}
	}

	writer.Flush()
	return buff.Bytes()
}
