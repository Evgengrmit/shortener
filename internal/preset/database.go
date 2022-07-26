package preset

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"ozonTask/pkg/link"
	"time"
)

const (
	DbHost     = "db.host"
	DbPort     = "db.port"
	DbUser     = "db.user"
	DbPassword = "DB_PASSWORD"
	DbName     = "db.dbname"
	Mode       = "db.sslmode"
)

func GetStorage() link.LinkStorage {
	var storage link.LinkStorage
	mode := flag.String("mode", "", "choose mode (memory/db)")

	if *mode == "memory" {
		log.Println("used in-memory storage")
		storage = InitLinkMemory()
	} else if *mode == "db" {
		log.Println("used database storage")
		storage = InitLinkSQL()
	} else {
		log.Fatalf("wrong mode error")
	}
	return storage
}
func InitLinkMemory() link.LinkStorage {
	return link.NewLinkMemory()

}

func InitLinkSQL() link.LinkStorage {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())

	}
	environment := fmt.Sprintf(`host=%s port=%s user=%s
 	password=%s dbname=%s sslmode=%s`,
		viper.GetString(DbHost),
		viper.GetString(DbPort),
		viper.GetString(DbUser),
		os.Getenv(DbPassword),
		viper.GetString(DbName),
		viper.GetString(Mode))
	fmt.Println(environment)
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
