package dbconnect

import (
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func CreateDBConnection() *sqlx.DB {
	viper.AddConfigPath("config")
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.name"))

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create DB Connection Failed")
		panic(err)
	}
	return db
}

func CreateTestingDBConnection(confDir string) *sqlx.DB {
	viper.AddConfigPath(fmt.Sprintf("%sconfig", confDir))
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", viper.GetString("database-test.host"), viper.GetString("database-test.port"), viper.GetString("database-test.user"), viper.GetString("database-test.password"), viper.GetString("database-test.name"))

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Create DB Connection Failed")
		panic(err)
	}

	return db
}
