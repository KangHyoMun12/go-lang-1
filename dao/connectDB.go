package dao

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const cmdSelect string = "SELECT * FROM "

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/trees")
	checkerr(err)
	return db
}

/*
GetListByTableName return list
*/
func GetListByTableName(tablename string) *sql.Rows {
	result, err := ConnectDB().Query(cmdSelect + tablename)
	checkerr(err)
	defer ConnectDB().Close()
	return result
}

func getElementByIDSingleTable(tablename string, id int) *sql.Row {
	result := ConnectDB().QueryRow(cmdSelect + tablename + " WHERE ID = " + strconv.Itoa(id))
	defer ConnectDB().Close()
	return result
}

/*
GetElementByID return list
*/
func GetElementByID(tablenameMain string, tablenameSub string, propKeyPrimaryMain string, propKeyUniqueSub string, id int) *sql.Row {
	var stringQuery = "SELECT "
	stringQuery += " p.Id, p.Name, p.Price, p.`Type`, c.Name "
	stringQuery += " FROM "
	stringQuery += tablenameMain + " as p "
	stringQuery += " LEFT JOIN "
	stringQuery += tablenameSub + " as c "
	stringQuery += " ON "
	stringQuery += " p." + propKeyPrimaryMain + " = c." + propKeyUniqueSub
	stringQuery += " WHERE "
	stringQuery += " p.Id = " + strconv.Itoa(id)
	result := ConnectDB().QueryRow(stringQuery)
	defer ConnectDB().Close()
	return result
}

func checkerr(err error) error {
	if err != nil {
		panic(err.Error())
	}
	return nil
}

/*
CheckAccount check account
*/
func CheckAccount(username string, password string) *sql.Row {
	//var stringQuery = "SELECT * FROM account WHERE username = '" + username + "' AND password = '" + password + "'"
	sqlprepare, err := ConnectDB().Prepare("SELECT * FROM account WHERE username = ? AND password = ?")
	if err != nil {
		panic(err.Error())
	}
	result := sqlprepare.QueryRow(username, password)
	defer ConnectDB().Close()
	return result
}

/*
AddAccount add acc
*/
func AddAccount(username string, email string, password string, phone string, showname string, birthday string) error {
	sqlprepare, err := ConnectDB().Prepare("INSERT INTO account(username,email,`password`,`phone`,showname,birthday) VALUES(?,?,?,?,?,?)")
	defer ConnectDB().Close()
	if err != nil {
		panic(err.Error())
	}
	_, err = sqlprepare.Exec(username, email, password, phone, showname, birthday)
	return err
}
