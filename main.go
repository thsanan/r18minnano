package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/thsanan/r18minnano/pogo"
)

func main() {
	db, err := sqlx.Open("mysql", "thsanan:IntelliP24.X@tcp(103.212.36.66:3306)/javdb")
	if err != nil {
		panic(err.Error())
	}

	act := pogo.Actress{}

	err = db.Get(&act, "SELECT * FROM javdb.actress WHERE actress.act_id=?", 1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", act)

}
