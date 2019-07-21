package DataImporter

import (
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

type DbConfig struct {
	DatabaseType string
	DatabasePort int
	DatabaseUser string
	DatabaseName string
	DatabasePW   string
}

type Database struct {
	Config     DbConfig
	Connection *sql.DB
}

func NewDb() (*Database, error) {
	db := &Database{}
	config := DbConfig{}

	err := config.Load()
	if err != nil {
		return db, err
	}

	db.Config = config
	connection, err := db.Connect()
	if err != nil {
		return db, err
	}

	db.Connection = connection

	return db, nil
}

func (dbc *DbConfig) Load() error {
	viper.SetConfigType("yaml")
	fileHandler, err := os.Open("config/database.yaml")
	if err != nil {
		return err
	}

	err = viper.ReadConfig(bufio.NewReader(fileHandler))
	if err != nil {
		return err
	}

	dbc.DatabaseType = "mysql"
	dbc.DatabasePort = viper.GetInt("mysql_port")
	dbc.DatabaseName = viper.GetString("mysql_database_name")
	dbc.DatabaseUser = viper.GetString("mysql_database_user")
	dbc.DatabasePW = viper.GetString("mysql_database_pw")

	return nil
}

func (dbc *DbConfig) GetDsn() string {
	return dbc.DatabaseUser + ":" + dbc.DatabasePW + "@/" + dbc.DatabaseName + "?parseTime=true"
}

func (db *Database) Connect() (*sql.DB, error) {
	dbConnection, err := sql.Open(db.Config.DatabaseType, db.Config.GetDsn())
	if err != nil {
		return nil, err
	}
	if err = dbConnection.Ping(); err != nil {
		return nil, err
	}
	return dbConnection, nil
}
