package config

import (
	"io"
	"test-project-hernan/src/libs/database"
	"test-project-hernan/src/libs/logger"
	"test-project-hernan/src/libs/sentry"
)

/*
slice of dependencies, io.Closes is an interface with method Close() error
all package that makes connections implements it
*/
var dependenciesToClose []io.Closer

/*
SetupCommonDependencies calls setup for each necessary dependencies
and registers them on one slice to be closed later
*/
func SetupCommonDependencies() {
	logger.SetupLogger()
	sentry.SetupSentry()
	database.SetupTestProjectGormDB()
	dependenciesToClose = []io.Closer{}
}

/*
TearDownCommonDependencies iterates each dependency and calls Close method
*/
func TearDownCommonDependencies() {
	for _, dependecy := range dependenciesToClose {
		dependecy.Close()
	}
}
