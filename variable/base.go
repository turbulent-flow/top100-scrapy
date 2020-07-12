package variable

// The place that you can share the various variables accross the whole project.

import (
	"os"
	"fmt"
	"net"
	"net/http"
	"time"
	"strconv"
)

var (
	Env =  os.Getenv("ENV")
	AppName = os.Getenv("APP_NAME")
	AppURI = os.Getenv("APP_URI")
	AMQPURL = os.Getenv("CLOUDAMQP_URL")
	dbName     = os.Getenv("DB_NAME")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbPort     = os.Getenv("DB_PORT")
	dbHost     = os.Getenv("DB_HOST")
	sslMode    = os.Getenv("SSL_MODE")
	maxPoolConns = os.Getenv("MAX_POOL_CONNECTIONS")
	minPoolConns = os.Getenv("MIN_POOL_CONNECTIONS")
	DBURL = buildDBURL()
	MigrationURL = buildMigrationURL()
	PopulationURL = buildPopulationURL()
	TestDBURL  = os.Getenv("TEST_DB_DSN")
	FixturesURI = buildFixturesURI()
	TestFixturesURI = buildTestFixturesURI()
	AWSregion = os.Getenv("AWS_S3_REGION")
    BucketName = os.Getenv("AWS_S3_BUCKET_NAME")
    AWSaccessId = os.Getenv("AWS_ACCESS_KEY_ID")
	AWSaccessSecret = os.Getenv("AWS_SECRET_ACCESS_KEY")
	S3BucketEndpoint = os.Getenv("AWS_S3_BUCKET_ENDPOINT")
	httpClientMaxIdleConnsPerHost = os.Getenv("HTTP_CLIENT_MAX_IDLE_CONNECTIONS_PER_HOST")
	HTTPclientPreconfigs = buildHTTPclientPreconfigs()
	Concurrency = os.Getenv("GOROUTINE_CONCURRENCY")
	NewRelicLicenseKey = os.Getenv("NEW_RELIC_LICENSE_KEY")
)

func buildDBURL() (dbURL string) {
	switch Env {
	case "development":
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=%s&pool_min_conns=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode, maxPoolConns, minPoolConns)
	default:
		dbURL = fmt.Sprintf("%s?sslmode=require&pool_max_conns=%s&pool_min_conns=%s", os.Getenv("DATABASE_URL"), maxPoolConns, minPoolConns)
	}
	return dbURL
}

func buildHTTPclientPreconfigs() *http.Transport {
	n, _ := strconv.Atoi(httpClientMaxIdleConnsPerHost)
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).DialContext,
		MaxIdleConns: 100,
		MaxIdleConnsPerHost: n,
	}
}

func buildMigrationURL() (migrationURL string) {
	switch Env {
	case "development":
		migrationURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)
	default:
		migrationURL = fmt.Sprintf("%s?sslmode=require", os.Getenv("DATABASE_URL"))
	}
	return migrationURL
}

func buildPopulationURL() (PopulationURL string) {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)
}

func buildTestFixturesURI() (uri string) {
	return fmt.Sprintf("%s/test/fixtures", AppURI)
}

func buildFixturesURI() (uri string) {
	return fmt.Sprintf("%s/fixtures", AppURI)
}
