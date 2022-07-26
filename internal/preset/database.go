package preset

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"ozonTask/pkg/link"
	"time"
)

type config struct {
	Port string `yaml:"port"`
	Db   struct {
		User    string `yaml:"user"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Dbname  string `yaml:"dbname"`
		Sslmode string `yaml:"sslmode"`
	} `yaml:"db"`
}

const DbPassword = "DB_PASSWORD"

func GetStorage() link.LinkStorage {
	var storage link.LinkStorage
	mode := flag.String("mode", "memory", "choose mode (memory/db)")

	if *mode == "memory" {
		log.Println("used in-memory storage")
		storage = link.NewLinkMemory()
	} else if *mode == "db" {
		log.Println("used database storage")
		storage = InitLinkSQL()
	} else {
		log.Fatalf("wrong mode error")
	}
	return storage
}

func InitLinkSQL() link.LinkStorage {

	var configure config

	f, err := os.ReadFile("configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(f, &configure); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())

	}

	environment := fmt.Sprintf(`host=%s port=%s user=%s
 	password=%s dbname=%s sslmode=%s`,
		configure.Db.Host,
		configure.Db.Port,
		configure.Db.User,
		os.Getenv(DbPassword),
		configure.Db.Dbname,
		configure.Db.Sslmode)
	db, err := sql.Open("postgres", environment)

	if err != nil {
		log.Fatalf("error opening db: %s", err.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalf("pinging database error: %s", err.Error())
	}

	return link.NewLinkSQL(db)
}
