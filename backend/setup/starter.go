package setup

import (
	"context"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"time"
)

func StartServer(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		panic(err)
	}
	address, err := cmd.Flags().GetString("addr")
	if err != nil {
		panic(err)
	}

	e := NewServer(debug)

	// Start server
	go func() {
		if err := e.Start(address); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
