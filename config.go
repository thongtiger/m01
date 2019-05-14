package main

import (
	"os"
	"strconv"
)

// Config ...
type Config struct {
	Port int
	Jwt  struct {
		Private string
		Public  string
	}
	Mongodb struct {
		Host     string
		Database string
		Username string
		Password string
	}
	Mssql struct {
		Host     string
		Database string
		Username string
		Password string
		Port     int
		SCHEMA0  string
		SCHEMA2  string
		SCHEMA3  string
	}
	Socket struct {
		URL string
	}
	Redisdb struct {
		Addr     string
		Password string
		DB       int
	}
}

// GetByENV is get config by environment variable
func (c *Config) GetByENV() {
	// port
	if val, ok := os.LookupEnv("PORT"); ok {
		if Num, err := strconv.Atoi(val); err == nil {
			c.Port = Num
		}
	}
	// mssql
	c.Mssql.Port = 1433 // default

	if val, ok := os.LookupEnv("MSSQL_HOST"); ok {
		c.Mssql.Host = val
	}
	if val, ok := os.LookupEnv("MSSQL_DB"); ok {
		c.Mssql.Database = val
	}
	if val, ok := os.LookupEnv("MSSQL_USERNAME"); ok {
		c.Mssql.Username = val
	}
	if val, ok := os.LookupEnv("MSSQL_PASSWORD"); ok {
		c.Mssql.Password = val
	}

	// mongo
	if val, ok := os.LookupEnv("MONGO_HOST"); ok {
		c.Mongodb.Host = val
	}
	if val, ok := os.LookupEnv("MONGO_DB"); ok {
		c.Mongodb.Database = val
	}
	if val, ok := os.LookupEnv("MONGO_USERNAME"); ok {
		c.Mongodb.Username = val
	}
	if val, ok := os.LookupEnv("MONGO_PASSWORD"); ok {
		c.Mongodb.Password = val
	}
}
