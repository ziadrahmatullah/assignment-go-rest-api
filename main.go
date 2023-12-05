package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	db := database.ConnectDB()

	addr := os.Getenv("APP_PORT")

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)

	wr := repository.NewWalletRepository(db)

	tr := repository.NewTransactionRepository(db)
	tu := usecase.NewTransactionUsecase(tr, wr)
	th := handler.NewTransactionHandler(tu)

	gr := repository.NewGameRepository(db)
	gu := usecase.NewGameUsecase(gr, wr)
	gh := handler.NewGameHandler(gu)

	opts := server.RouterOpts{
		UserHandler:        uh,
		TransactionHandler: th,
		GameHandler:        gh,
	}
	r := server.NewRouter(opts)

	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
