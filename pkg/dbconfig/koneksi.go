package dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func GetDBConnection() (*sql.DB, error) {

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}

	dbUser := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	dbParams := viper.GetString("DB_PARAMS")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbParams )

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %s", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %s", err)
	}

	return db, nil
}
