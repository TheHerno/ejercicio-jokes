package main

import (
	"crypto/tls"
	"log"
	"os"
	"test-project-hernan/src/libs/env"

	"github.com/go-pg/pg/v9"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

const directory = "migrations/internal"

func main() {
	dbHost := env.TestProjectServicePostgresqlHost
	dbPort := env.TestProjectServicePostgresqlPort
	dbName := env.TestProjectServicePostgresqlName
	if env.AppEnv == "testing" {
		dbName = env.TestProjectServicePostgresqlNameTest
	}
	dbUser := env.TestProjectServicePostgresqlUsername
	dbPassword := env.TestProjectServicePostgresqlPassword
	dbSSLMode := env.TestProjectServicePostgresqlSSLMode

	options := &pg.Options{
		Addr:     dbHost + ":" + dbPort,
		User:     dbUser,
		Database: dbName,
		Password: dbPassword,
	}
	if dbSSLMode != "disable" {
		options.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	db := pg.Connect(options)

	err := migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
