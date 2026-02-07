package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/controller"
	"github.com/rlapz/mmweb/middleware"
	"github.com/rlapz/mmweb/service"
	"github.com/rlapz/mmweb/util"

	_ "modernc.org/sqlite"

	repoSqlite "github.com/rlapz/mmweb/repo/sqlite"
)

func serve(mux *middleware.Middleware, cfg *config.Config) error {
	srv := new(http.Server)
	srv.Handler = mux
	srv.Addr = fmt.Sprintf("%s:%s", cfg.ListenHost, cfg.ListenPort)

	go func() {
		log.Printf("serving on: [%s]\n", srv.Addr)
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("error: serve: ListenAndServe", err.Error())
		}

		log.Println("connection listener stopped")
	}()

	moeChan := make(chan os.Signal, 1)
	signal.Notify(moeChan, syscall.SIGINT)
	<-moeChan

	shutCtx, shutRel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutRel()

	err := srv.Shutdown(shutCtx)
	if err != nil {
		log.Println("error: serve: Shutdown", err.Error())
		return err
	}

	return nil
}

func run(cfg *config.Config) error {
	defer func() {
		log.Println("gracefully stopped :-)")
	}()

	db, err := util.SqlitePoolNew(cfg.DbPath, cfg.DbPoolInitSize)
	if err != nil {
		log.Println("error: Run: SqlitePoolNew:", err)
		return err
	}
	defer db.Destory()

	mux := middleware.New(cfg)
	repp := repoSqlite.New(db)
	srvv := service.New(repp)

	controller.Init(cfg, mux, srvv)

	err = serve(mux, cfg)
	if err != nil {
		log.Println("error: Run: serve:", err)
	}

	return err
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println("main: config.Load:", err.Error())
		os.Exit(1)
	}

	if run(cfg) != nil {
		os.Exit(1)
	}
}
