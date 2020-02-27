package main

import (
	"database/sql"
	"fmt"

	//As we use here sqlite3 here we need to define this package
	_ "github.com/mattn/go-sqlite3"
)

const pathToTheDatabase = "./sqlite/test-database.sqlite"

var (
	DB  *sql.DB
	err error
)

func main() {
	DB, err = sql.Open("sqlite3", pathToTheDatabase)
	if err != nil {
		panic(err)
	}

	//Get all results from the database
	var result interface{}
	result, err = getAllRows()
	if err != nil {
		panic(err)
	}

	fmt.Println("Total results:")
	fmt.Println(result)

	lastInsertID, err := insertRow("Testing")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Inserted one row with ID: %d", lastInsertID))

	//Get result by ID
	result, err = getNameByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Name is : %s", result))

	rowsAffected, err := deleteRow(lastInsertID)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Deleted item with ID: %d; Total rows affected: %d", lastInsertID, rowsAffected))
}

func insertRow(name string) (int64, error) {
	result, err := DB.Exec("insert into test_table (name) values (?)", name)
	if err != nil {
		return int64(0), err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return int64(0), err
	}

	return lastInsertID, nil
}

func getAllRows() (map[int64]string, error) {
	var result = map[int64]string{}

	//Select all rows
	rows, err := DB.Query("select id, name from test_table")
	if err == sql.ErrNoRows {
		fmt.Println("There is no rows in the database")
		return map[int64]string{}, nil
	} else if err != nil {
		return map[int64]string{}, err
	}

	var (
		id   int64
		name string
	)

	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			return map[int64]string{}, err
		}

		result[id] = name
	}

	return result, nil
}

func getNameByID(id int64) (string, error) {
	var name string
	if err := DB.QueryRow("select name from test_table where id = ?", id).Scan(&name); err != nil {
		return "", err
	}

	return name, nil
}

func deleteRow(ID int64) (int64, error) {
	result, err := DB.Exec("delete from test_table where id = ?", ID)
	if err != nil {
		return int64(0), err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return int64(0), err
	}

	return rowsAffected, nil
}
