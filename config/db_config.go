package config

// Implement your configuration here
import (
        "fmt"
        "log"
        "os"
        "path/filepath"
        "github.com/joho/godotenv"
)

var WorkDir, _ = os.Getwd()
// var baseDir = filepath.Join(WorkDir, "..")//with Test
var baseDir = filepath.Join(WorkDir)//with main

var DBName string
var DBUser string
var DBPass string
var Dialect string
var DBHost string
var DBPort string
var TBName string
var Connect string
var Host string
var Port string
var Address string

func init() {
        err := godotenv.Load(filepath.Join(baseDir,".env"))
        if err != nil {
                log.Fatal(err.Error())
        }

        DBName = os.Getenv("dbname")
        DBUser = os.Getenv("dbuser")
        DBPass = os.Getenv("dbpass")
        DBHost = os.Getenv("dbhost")
        DBPort = os.Getenv("dbport")
        Host = os.Getenv("host")
        Port = os.Getenv("port")
        TBName = os.Getenv("tbname")
        Dialect = os.Getenv("dialect")
        Connect = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",DBUser,DBPass,DBHost,DBPort,DBName)
        Address = fmt.Sprintf("%v:%v",Host,Port)
}
