package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/huey-emma/cms/db"
	"github.com/huey-emma/cms/internal/person"
	"github.com/huey-emma/cms/internal/utils/jsonlog"
	"github.com/huey-emma/cms/internal/utils/lib"
)

type Router struct {
	port   string
	mux    *http.ServeMux
	logger *jsonlog.Logger
	*db.Database
}

func New(port string, logger *jsonlog.Logger, db *db.Database) *Router {
	mux := new(http.ServeMux)
	return &Router{port, mux, logger, db}
}

func (router *Router) Use() *Router {
	personR := person.NewRepository(router.DB)
	personS := person.NewService(personR)
	personH := person.NewHandler(personS)

	router.mux.Handle("/api", router.recoverPanic(router.logRequests(personH)))
	router.mux.Handle("/api/", router.recoverPanic(router.logRequests(personH)))

	return router
}

func (router *Router) Run() error {
	server := &http.Server{
		Addr:         "0.0.0.0:" + router.port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      router.mux,
	}

	errch := make(chan error)

	go func() {
		quitch := make(chan os.Signal, 1)
		signal.Notify(quitch, syscall.SIGTERM, syscall.SIGINT)
		router.logger.PrintInfo("signal recieved", lib.Map[string]{"signal": (<-quitch).String()})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			errch <- err
		}

		errch <- nil
	}()

	router.logger.PrintInfo("server is starting", lib.Map[string]{"port": router.port})

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	if err := <-errch; err != nil {
		return err
	}

	return nil
}

func (router *Router) recoverPanic(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				router.logger.PrintError(fmt.Errorf("%v", r), nil)
				http.Error(w, "internal server error", 500)
			}
		}()

		n.ServeHTTP(w, r)
	})
}

func (router *Router) logRequests(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router.logger.PrintInfo(r.Method, lib.Map[string]{"path": r.RequestURI})
		n.ServeHTTP(w, r)
	})
}
