package cmd

import (
	"fmt"

	ofModule "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/handler"
	ofRepo "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/repository"
	ofService "bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/offer/service"

	"github.com/sirupsen/logrus"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/internal/base/handler"

	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/db"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/httpclient"
	"bitbucket.org/bitbucketnobubank/paylater-cms-api/pkg/metric"

	"os"
	"strconv"
)

var (
	params map[string]string

	baseHandler  *handler.BaseHTTPHandler
	offerHandler *ofModule.HTTPHandler
	httpClient   httpclient.Client

	mysqlClientRepo  *db.MySQLClientRepository
	statsdMonitoring metric.StatsdMonitoring
)

func initMySQL() {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbname := os.Getenv("DB_NAME")
	uname := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	mysqlClientRepo, _ = db.NewMySQLRepository(host, uname, pass, dbname, port)
}

func initInfrastructure() {
	initMySQL()
	initLog() // Init log after baseHandler

	httpClientFactory := httpclient.New()
	httpClient = httpClientFactory.CreateClient()

	var err error

	if err != nil {
		logrus.Errorln(err)
		fmt.Println(err)
	}
}

func isProd() bool {
	return os.Getenv("APP_ENV") == "production"
}

func initHTTP() {
	params = initParams()
	initInfrastructure()

	params["mysql_tz"] = mysqlClientRepo.TZ

	offerRepo := ofRepo.NewRepository(mysqlClientRepo.DB)

	offerService := ofService.NewService(offerRepo)

	baseHandler = handler.NewBaseHTTPHandler(mysqlClientRepo.DB, httpClient, params,
		offerService, statsdMonitoring)

	offerHandler = ofModule.NewHTTPHandler(baseHandler, offerService)

	fmt.Println("INFO: Init and load module completed. Server started.\n---")
}

func initLog() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	lv := os.Getenv("LOG_LEVEL_DEV")
	level := logrus.InfoLevel
	switch lv {
	case "PanicLevel":
		level = logrus.PanicLevel
	case "FatalLevel":
		level = logrus.FatalLevel
	case "ErrorLevel":
		level = logrus.ErrorLevel
	case "WarnLevel":
		level = logrus.WarnLevel
	case "InfoLevel":
		level = logrus.InfoLevel
	case "DebugLevel":
		level = logrus.DebugLevel
	case "TraceLevel":
		level = logrus.TraceLevel
	default:
	}

	// Only log above level
	if isProd() {
		// Only Warn and Error log for prod
		logrus.SetLevel(logrus.WarnLevel) // Set default InfoLevel
	} else {
		// Set log level for staging.
		if lv == "" && os.Getenv("APP_DEBUG") == "True" {
			level = logrus.DebugLevel
		}
		logrus.SetLevel(level)
	}
}
