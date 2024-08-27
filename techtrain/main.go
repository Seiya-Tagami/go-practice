package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"techtrain-go-practice/db"
	"techtrain-go-practice/handler/router"
)

func main() {
	err := setupEnv()
	if err != nil {
		log.Fatalln("main: failed to set up environment variables, err =", err)
	}

	RunApp()
}

func RunApp() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH is not set")
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		log.Fatalf("main: failed to open db")
	}
	defer todoDB.Close()

	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := router.NewRouter(todoDB)

	srv := &http.Server{
		Addr:              port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Minute,
		WriteTimeout:      10 * time.Minute,
	}

	var wg sync.WaitGroup

	log.Printf("Starting server on localhost%s...\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Print(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()

	wg.Wait()
}

func setupEnv() error {
	err := os.Setenv("PORT", ":8080")
	if err != nil {
		return err
	}
	err = os.Setenv("DB_PATH", ".sqlite3/todo.db")
	if err != nil {
		return err
	}
	err = os.Setenv("BASIC_AUTH_USER_ID", "user")
	if err != nil {
		return err
	}
	err = os.Setenv("BASIC_AUTH_PASSWORD", "password")
	if err != nil {
		return err
	}
	return nil
}
