package env

import (
	"os"
	"strconv"
	"time"
)

var (

	// AppEnv Application Environment
	AppEnv string

	// EnvironmentName Environment Name
	EnvironmentName string

	// ServiceName Service Name
	ServiceName string

	// ServiceVersion Service version
	ServiceVersion string

	// TestProjectServiceSecondsBetweenAttempts TestProjectService Interval in Seconds between attempts
	TestProjectServiceSecondsBetweenAttempts time.Duration

	// TestProjectServicePostgresqlHost TestProjectService PostgreSQL host
	TestProjectServicePostgresqlHost string

	// TestProjectServicePostgresqlPort TestProjectService PostgreSQL port
	TestProjectServicePostgresqlPort string

	// TestProjectServicePostgresqlName TestProjectService PostgreSQL name
	TestProjectServicePostgresqlName string

	// TestProjectServicePostgresqlNameTest TestProjectService PostgreSQL name Test
	TestProjectServicePostgresqlNameTest string

	// TestProjectServicePostgresqlUsername TestProjectService PostgreSQL app username
	TestProjectServicePostgresqlUsername string

	// TestProjectServicePostgresqlPassword TestProjectService PostgreSQL app password
	TestProjectServicePostgresqlPassword string

	// TestProjectServicePostgresqlSSLMode TestProjectService PostgreSQL ssl mode
	TestProjectServicePostgresqlSSLMode string

	// TestProjectServiceGrpcPort TestProjectService gRPC port
	TestProjectServiceGrpcPort string

	// TestProjectServiceRestPort TestProjectService Rest port
	TestProjectServiceRestPort string

	// WhiteList White List
	WhiteList string

	// External services

	// EventLoggerURL Logger service URL
	EventLoggerURL string

	// EventLoggerUser Logger service user
	EventLoggerUser string

	// EventLoggerPassword Logger service password
	EventLoggerPassword string
)

func init() {
	// App Environment
	AppEnv = os.Getenv("APP_ENV")

	// Environment Name
	EnvironmentName = os.Getenv("ENVIRONMENT_NAME")
	// Service Name
	ServiceName = os.Getenv("SERVICE_NAME")
	// Service Version
	ServiceVersion = os.Getenv("VERSION")

	// TestProjectService - gRPC
	TestProjectServiceGrpcPort = os.Getenv("TEST_PROJECT_HERNAN_GRPC_PORT")

	// TestProjectService - Rest
	TestProjectServiceRestPort = os.Getenv("TEST_PROJECT_HERNAN_REST_PORT")

	// TestProjectService Interval in Seconds Between Attempts
	var seconds int
	processIntEnvVar(&seconds, "TEST_PROJECT_HERNAN_SECONDS_BETWEEN_ATTEMPTS", 60)
	TestProjectServiceSecondsBetweenAttempts = time.Duration(seconds) * time.Second

	// TestProjectService - PostgreSQL
	TestProjectServicePostgresqlHost = os.Getenv("TEST_PROJECT_HERNAN_POSTGRESQL_HOST")
	TestProjectServicePostgresqlPort = os.Getenv("TEST_PROJECT_HERNAN_POSTGRESQL_PORT")
	TestProjectServicePostgresqlName = os.Getenv("TEST_PROJECT_HERNAN_POSTGRESQL_NAME")
	TestProjectServicePostgresqlNameTest = os.Getenv("TEST_PROJECT_HERNAN_POSTGRESQL_NAME_TEST")
	TestProjectServicePostgresqlUsername = os.Getenv("TEST_PROJECT_HERNAN_POSTGRESQL_USERNAME")
	TestProjectServicePostgresqlPassword = os.Getenv("TEST_PROJECT_HERNAN_POSTGRESQL_PASSWORD")
	TestProjectServicePostgresqlSSLMode = os.Getenv("TEST_PROJECT_HERNAN_POSTGRESQL_SSLMODE")

	// Logger service
	EventLoggerURL = os.Getenv("EVENT_LOGGER_URL")
	EventLoggerUser = os.Getenv("EVENT_LOGGER_USER")
	EventLoggerPassword = os.Getenv("EVENT_LOGGER_PASSWORD")

	// White list
	WhiteList = os.Getenv("WHITE_LIST")
}

// processIntEnvVar gets environment variable from os and parses it to int
func processIntEnvVar(intVar *int, envKey string, defaultValue int) {
	var err error
	*intVar, err = strconv.Atoi(os.Getenv(envKey))
	if err != nil {
		*intVar = defaultValue
	}
}
