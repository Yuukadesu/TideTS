// TideTS CLI：通过 Session 执行最简 SQL（INSERT / SELECT）。
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hanami/tidets/cli/repl"
	"github.com/hanami/tidets/client/session"
)

func main() {
	host := flag.String("host", session.DefaultHost, "DataNode host")
	port := flag.Int("port", session.DefaultPort, "DataNode port")
	user := flag.String("user", session.DefaultUsername, "username")
	pass := flag.String("password", session.DefaultPassword, "password")
	fetchSize := flag.Int("fetch-size", session.DefaultFetchSize, "session fetch size for SELECT default limit")
	flag.Parse()

	sess, err := session.New(
		session.WithHost(*host),
		session.WithPort(*port),
		session.WithUsername(*user),
		session.WithPassword(*pass),
		session.WithFetchSize(*fetchSize),
	)
	if err != nil {
		log.Fatalf("session: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := sess.Open(ctx); err != nil {
		log.Fatalf("open session: %v", err)
	}
	defer func() {
		if err := sess.Close(); err != nil {
			log.Printf("close session: %v", err)
		}
	}()

	fmt.Printf("connected %s:%d as %s (session opened)\n", sess.Host(), sess.Port(), sess.Username())

	if err := repl.Run(ctx, sess, repl.Options{
		In:     os.Stdin,
		Out:    os.Stdout,
		Err:    os.Stderr,
		Prompt: "tidets> ",
	}); err != nil {
		log.Fatalf("repl: %v", err)
	}
}
