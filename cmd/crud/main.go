package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shohrukh56/mcDonalds/cmd/crud/app"
	"github.com/shohrukh56/mcDonalds/pkg/crud/services/burgers"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var (
	host = flag.String("host", "0.0.0.0", "Server host")
	port = flag.String("port", "9999", "Server port")
	dsn  = flag.String("dsn", "postgres://user:pass@localhost:5432/app", "Postgres DSN")
)

func main() {
	flag.Parse()
	portEnv, ok := os.LookupEnv("PORT")
	if ok {
		log.Print()
		*port = portEnv
	}
	dsnEnv, ok := os.LookupEnv("DATABASE_URL")
	if ok {
		*dsn = dsnEnv
	}
	log.Print(*dsn)
	addr := net.JoinHostPort(*host, *port)
	start(addr, *dsn)
}

func start(addr string, dsn string) {
	router := app.NewExactMux()
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	burgersSvc := burgers.NewBurgersSvc(pool)
	err = burgersSvc.InitDB()
	if err != nil {
		panic(err)
	}

	server := app.NewServer(
		router,
		pool,
		burgersSvc, // DI + Containers
		filepath.Join("web", "templates"),
		filepath.Join("web", "assets"),
	)
	server.InitRoutes()

	panic(http.ListenAndServe(addr, server))
}
