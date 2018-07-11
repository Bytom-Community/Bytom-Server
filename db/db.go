package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
)

type DB struct {
	Address string
	Port    string
	User    string
	Pass    string
	DBName  string
}

func NewDB(addr, user, pass, port, dbname string) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, addr, port, dbname)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return
	}

	Engine = engine
	return nil
}

func IsInit() bool {
	if Engine != nil {
		return true
	}
	return false
}

func (db *DB) Close() error {
	if Engine != nil {
		return Engine.Close()
	}
	return nil
}
