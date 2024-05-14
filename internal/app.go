package app

import (
	"bashscripts/configs"
	"bashscripts/internal/adapters"
	"bashscripts/internal/repository"
	"bashscripts/internal/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func Start() error {
	s := &Server{
		router: mux.NewRouter(),
	}
	db, err := NewDB()
	if err != nil {
		return fmt.Errorf("can't connect to db: %v\n", err)
	}
	//httpserver
	repo := repository.NewRepo(db)
	service := service.NewService(repo)
	handler := adapters.NewHTTPHandle(service)
	handler.Register(s.router)

	log.Println("Starting Scripts at port: 8080")
	return http.ListenAndServe(":8080", s)
}

func NewDB() (*pgx.Conn, error) {
	c, err := configs.New("./configs/main.yaml")
	if err != nil {
		return nil, fmt.Errorf("can't receive config db: %v \n", err)
	}
	psqlInfo := fmt.Sprintf(c.DB.Conn)

	db, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error connection: %v\n", err)
	}

	return db, nil
}

// ServeHTTP
func (h *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
