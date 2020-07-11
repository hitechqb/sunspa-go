package sdk

import "os"

type Connection struct {
	Dialect    string
	DbUser     string
	DbPassWord string
	DbHost     string
	DbPort     string
	DbName     string
}

func NewConnection() *Connection {
	conn := &Connection{
		Dialect:    os.Getenv("GORM_DB_TYPE"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassWord: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
	}

	return conn
}
