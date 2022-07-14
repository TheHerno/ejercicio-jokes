package database

import (
	"sync"
	"test-project-hernan/src/libs/env"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func resetOnceTestProject() {
	once = sync.Once{}
}

func TestSetupTestProjectGormDB(t *testing.T) {
	t.Run("Should success on", func(t *testing.T) {
		t.Run("Should success on", func(t *testing.T) {
			resetOnceTestProject()
			db := SetupTestProjectGormDB()
			sqlDB, _ := db.DB()
			errPing := sqlDB.Ping()
			//Data Assertion
			assert.NotNil(t, db)
			assert.NoError(t, errPing)
			t.Cleanup(func() {
			})
		})
		t.Run("Wait for postgres", func(t *testing.T) {
			// Smaller time & wrong DB name
			oldDelta := env.TestProjectServiceSecondsBetweenAttempts
			env.TestProjectServiceSecondsBetweenAttempts = time.Second / 2
			oldValue := env.TestProjectServicePostgresqlNameTest
			env.TestProjectServicePostgresqlNameTest = "TestProject_SERVICE_POSTGRESQL_NAME_not_found"
			var db *gorm.DB
			var errPing error
			wait := make(chan bool)
			go func() {
				resetOnceTestProject()
				db = SetupTestProjectGormDB()
				sqlDB, _ := db.DB()
				errPing = sqlDB.Ping()
				wait <- true
			}()
			time.Sleep(env.TestProjectServiceSecondsBetweenAttempts)
			env.TestProjectServicePostgresqlNameTest = oldValue
			<-wait

			//Data Assertion
			assert.NotNil(t, db)
			assert.NoError(t, errPing)
			t.Cleanup(func() {
				env.TestProjectServicePostgresqlNameTest = oldValue
				env.TestProjectServiceSecondsBetweenAttempts = oldDelta
			})
		})
	})
}

func TestGetTestProjectGormConnection(t *testing.T) {
	t.Run("Should success when the connection is already open", func(t *testing.T) {
		resetOnceTestProject()
		db := SetupTestProjectGormDB()
		dbSingleton := GetTestProjectGormConnection()
		sqlDB, _ := dbSingleton.DB()
		errPing := sqlDB.Ping()
		//Data Assertion
		assert.Equal(t, db, dbSingleton)
		assert.NoError(t, errPing)
	})
}
