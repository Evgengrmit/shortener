package preset

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"ozonTask/pkg/link"
	"time"
)

func InitLinkMemory() link.LinkStorage {
	return link.NewLinkMemory()

}

func InitLinkSQL() link.LinkStorage {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}
	if err := godotenv.Load(""); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())

	}
	environment := fmt.Sprintf(`host=%s port=%s user=%s
 	password=%s dbname=%s sslmode=%s`,
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.user"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.dbname"),
		viper.GetString("sslmode"))
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
