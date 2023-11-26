package database

import (
  "ttd/config"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "log"
  "time"
)

func GetConnection() *sql.DB {
  db , err := sql.Open(config.Dialect,config.Connect)
  if err != nil {
    log.Fatal(err.Error())
    return nil
  }
  db.SetConnMaxLifetime(time.Minute * 60)
  db.SetConnMaxIdleTime(time.Minute * 5)
  db.SetMaxOpenConns(100)
  db.SetMaxIdleConns(10)
  return db
}