package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ngohoang211020/blog/internal/config"
	"github.com/ngohoang211020/blog/internal/jsonlog"
	models2 "github.com/ngohoang211020/blog/internal/models"
	"log"
	"os"
	"time"
)

const version = "1.0.0"

func main() {
	start()
}

func start() {
	err := config.Config.LoadConfig()
	if err != nil {
		log.Fatal("Can't load configuration from env file")
	}

	logger := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()
	logger.PrintInfo("database connection pool established", nil)

	models := models2.NewModels(db)
	app := &application{
		logger: logger,
		model:  models,
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Config.DSN)
	if err != nil {
		return nil, err
	}
	// Set the maximum number of open (in-use + idle) connections in the pool. Note that
	// passing a value less than or equal to 0 will mean there is no limit.
	db.SetMaxOpenConns(config.Config.MaxOpenConn)
	// Set the maximum number of idle connections in the pool. Again, passing a value
	// less than or equal to 0 will mean there is no limit.
	db.SetMaxIdleConns(config.Config.MaxIdleConn)
	// Use the time.ParseDuration() function to convert the idle timeout duration string
	// to a time.Duration type.
	duration, err := time.ParseDuration(config.Config.MaxIdleTimeout)
	if err != nil {
		return nil, err
	}
	// Set the maximum idle timeout.
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Use PingContext() to establish a new connection to the database, passing in the
	// context we created above as a parameter. If the connection couldn't be
	// established successfully within the 5 second deadline, then this will return an
	// error.

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	// Return the sql.DB connection pool.
	return db, nil

}
