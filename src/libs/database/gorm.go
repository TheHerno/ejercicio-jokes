package database

import (
	"fmt"
	"sync"
	"test-project-hernan/src/libs/env"
	"test-project-hernan/src/libs/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	ormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db   *gorm.DB
	once sync.Once
)

//CreateTestProjectConnectionString returns the connection string based on environment variables
func CreateTestProjectConnectionString() string {
	//db config vars
	dbHost := env.TestProjectServicePostgresqlHost
	dbPort := env.TestProjectServicePostgresqlPort
	dbName := env.TestProjectServicePostgresqlName
	dbUser := env.TestProjectServicePostgresqlUsername
	dbPassword := env.TestProjectServicePostgresqlPassword
	dbSSLMode := env.TestProjectServicePostgresqlSSLMode
	if env.AppEnv == "testing" {
		dbName = env.TestProjectServicePostgresqlNameTest
	}
	//Make connection string with interpolation
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
	return connectionString
}

/*
SetupTestProjectGormDB open the pool connection in db var and return it
*/
func SetupTestProjectGormDB() *gorm.DB {
	once.Do(func() {
		config := &gorm.Config{
			Logger: ormlogger.Default.LogMode(ormlogger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
		//connect to db
		var dbError error
		db, dbError = gorm.Open(postgres.Open(CreateTestProjectConnectionString()), config)
		for dbError != nil {
			logger.GetInstance().Error("Failed to connect to own-database")
			time.Sleep(env.TestProjectServiceSecondsBetweenAttempts)
			logger.GetInstance().Info("Retrying...")
			db, dbError = gorm.Open(postgres.Open(CreateTestProjectConnectionString()), config)
		}
		logger.GetInstance().Info("Connected to own-database!")
		setConnectionMaxLifetime(db, 0) //To be reused forever
	})
	return db
}

/*
GetTestProjectGormConnection return db pointer which already have an open connection
*/
func GetTestProjectGormConnection() *gorm.DB {
	return SetupTestProjectGormDB()
}
