package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb" // sqlserver driver
)

type MssqlDB struct {
	db *sql.DB
}

func NewMssql() *sql.DB {
	// initial database
	db, err := sql.Open("sqlserver", fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		CF.Mssql.Host,
		CF.Mssql.Username,
		CF.Mssql.Password,
		CF.Mssql.Port,
		CF.Mssql.Database,
	))
	if err != nil {
		log.Println("Error creating connection pool: ", err.Error())
		os.Exit(1) // <----------  exit before run
	}

	err = db.PingContext(context.Background()) // ping background
	if err != nil {
		log.Printf("mssql ping not found!, ip=%s", CF.Mssql.Host)
		os.Exit(1) // <----------  exit before run
	} else {
		log.Printf("mssql is connected! ip=%s", CF.Mssql.Host)
	}
	return db
}

func (h *MssqlDB) GetProfile(mid int) (rt ProfileForm) {
	db := h.db
	ctx := context.Background()
	err := db.QueryRowContext(ctx, os.Getenv("tsql_GetProfile"), sql.Named("p1", mid)).Scan(
		&rt.Username,
		&rt.EmailAddress,
		&rt.Firstname,
		&rt.Lastname,
		&rt.Gender,
		&rt.CountryCode,
		&rt.DateOfBirth,
		&rt.MonthOfBirth,
		&rt.YearOfBirth,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", mid)
	case err != nil:
		log.Printf("query error: %v\n", err)
	default:
		log.Print("has user accounts")
	}
	return
}

func (h *MssqlDB) ExistUsername(username string) bool {
	db := h.db
	ctx := context.Background()

	var result string

	err := db.QueryRowContext(ctx, os.Getenv("tsql_ExistUsername"), sql.Named("p1", username)).Scan(&result)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %s\n", result)
		return false
	case err != nil:
		log.Printf("query error: %v\n", err)
		return false
	default:
		log.Printf("username is %q, account created\n", result)
		return true
	}
}
func (h *MssqlDB) ExistEmail(email string) bool {
	db := h.db
	ctx := context.Background()

	var result string
	err := db.QueryRowContext(ctx, os.Getenv("tsql_ExistEmail"), sql.Named("p1", email)).Scan(&result)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no email name %s\n", result)
		return false
	case err != nil:
		log.Printf("query error: %v\n", err)
		return false
	default:
		log.Printf("email : %q, already exist \n", result)
		return true
	}
}
func (h *MssqlDB) ExistPhone(phone string) bool {
	db := h.db
	ctx := context.Background()

	var result string
	err := db.QueryRowContext(ctx, os.Getenv("tsql_ExistPhone"), sql.Named("p1", phone)).Scan(&result)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no email name %s\n", result)
		return false
	case err != nil:
		log.Printf("query error: %v\n", err)
		return false
	default:
		log.Printf("email : %q, already exist \n", result)
		return true
	}
}
