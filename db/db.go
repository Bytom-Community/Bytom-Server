package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type DB struct {
	Engine  *xorm.Engine
	Address string
	Port    string
	User    string
	Pass    string
	DBName  string
}

func NewDB(addr, user, pass, port, dbname string) (db *DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, addr, port, dbname)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return
	}

	return &DB{
		Engine:  engine,
		Address: addr,
		Port:    port,
		User:    user,
		DBName:  dbname,
	}, nil
}

func (db *DB) Close() error {
	return db.Engine.Close()
}
