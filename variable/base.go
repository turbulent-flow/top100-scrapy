package variable

// The place that you can share the various variables accross the whole project.

import (
	"os"
	"fmt"
)

var (
	Env =  os.Getenv("ENV")
	AppURI = os.Getenv("APP_URI")
	AMQPURL = os.Getenv("AMQP_URL")
	dbName     = os.Getenv("DB_NAME")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("c")
	dbPort     = os.Getenv("DB_PORT")
	dbHost     = os.Getenv("DB_HOST")
	sslMode    = os.Getenv("SSL_MODE")
	maxPoolConns = os.Getenv("MAX_POOL_CONNECTIONS")
	minPoolConns = os.Getenv("MIN_POOL_CONNECTIONS")
	DBURL = buildDBURL()
	TestDBURL  = os.Getenv("TEST_DB_DSN")
	FixturesURI = os.Getenv("FIXTURES_URI")
)

func buildDBURL() (dbURL string) {
	dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=%s&pool_min_conns=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode, maxPoolConns, minPoolConns)
	return dbURL
}
