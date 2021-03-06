package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

/*
 * Wrapper for sql.Query, meant to be the main query interface
 */
func Query( query string, parameters ...interface{} ) ( rows *sql.Rows, err error ) {
	var co *sql.DB

	co, err = sql.Open( "postgres", Cfg.DatabaseUrl )
	if err != nil {
		fmt.Printf( "can't connect to database : %v\n", err )
		return rows, err
	}
	defer co.Close()

	rows, err = co.Query( query, parameters... )
	if err != nil {
		return rows, err
	}

	return
}
