package main

import (
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fasthttp"
)

var server *fasthttp.Server
var shutdownCh chan os.Signal
var ormer orm.Ormer

// User model
type User struct {
	Id         int       `orm:"auto"`
	Name       string    `orm:"size(100)"`
	AccessedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	// Graceful termination channel attach OS signals
	shutdownCh = make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGTERM, syscall.SIGINT)

	// Set up the application from environment parameters
	initViper()
	// Enable the application logger (efficient uber/zap)
	initLogger()
	// Enable buffered custom logfile output (used when BACKEND=file)
	initLogFile()
	// Register prometheus metrics
	registerMetrics()
	// Register ORM backend
	registerORM()
}

// Wait for OS termination singls and perform cleanup
func waitForShutdown() {
	var wg sync.WaitGroup
	defer logger.Sync()  // flush logger
	defer logfile.Sync() // flush logfile

	wg.Add(1)
	go func() {
		<-shutdownCh
		wg.Done()
	}()
	wg.Wait()

	// Gracefully shut down the server
	err := server.Shutdown()
	if err != nil {
		logger.Panicw("Failed shutting down fasthttp", "err", err.Error())
	}
}

// Register beego ORM with SQLite driver
func registerORM() {
	if strings.ToLower(appConfig.Backend) != "sqlite" {
		return
	}

	// Set up sqlite orm and the database `default'
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", appConfig.FilePath)
	orm.RegisterModel(new(User))
	orm.Debug = isDevelopment()
	ormer = orm.NewOrm()

	// Populate database tables and list to ORM commands
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logger.Panic(err)
	}
	logger.Warnw("Initialized sqlite backend", "filepath", appConfig.FilePath)
}

func main() {
	// Create router serving the defined HTTP endpoints
	router := NewRouter()

	// Start the fasthttp server in a background Goroutine
	handler := requestMetrics(router.Handler)
	go func() {
		server = &fasthttp.Server{Handler: handler}
		err := server.ListenAndServe(":" + appConfig.Port)
		if err != nil {
			logger.Panicw("Failed to start fasthttp", "err", err.Error())
		}
	}()

	logger.Warn("Started app")

	// Perform gracefull shutdown on SIGINT and SIGTERM
	waitForShutdown()
}
