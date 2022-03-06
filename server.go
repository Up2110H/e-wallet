package e_wallet

import (
	"context"
	"github.com/Up2110H/e-wallet/pkg/handler"
	"github.com/Up2110H/e-wallet/pkg/store"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
	handlers   *handler.Handler
	store      *store.Store
}

func NewServer() *Server {
	return &Server{handlers: handler.NewHandler()}
}

func (s *Server) Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	s.configureServer()
	if err := s.configureStore(); err != nil {
		log.Fatalf("DB Init Error: %s", err.Error())
	}

	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Server Run Error: %s", err.Error())
	}

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) configureServer() {
	s.httpServer = &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        s.handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}

func (s *Server) configureStore() error {
	storeConfig := store.NewConfig()
	st := store.NewStore(storeConfig)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
