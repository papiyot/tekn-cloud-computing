package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type propinsi struct {
	id       int
	propinsi string
}

func (p propinsi) String() string {
	return fmt.Sprintf("%d, %s", p.id, p.propinsi)
}

func main() {
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/laravel")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected")
		// var id = 1
		rows, err := db.Query("select id, propinsi from propinsi ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()

		var result []propinsi

		for rows.Next() {
			var each = propinsi{}
			var err = rows.Scan(&each.id, &each.propinsi)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result = append(result, each)
		}

		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, each := range result {
			fmt.Println(each.String())
		}
	}
	defer db.Close()
}
