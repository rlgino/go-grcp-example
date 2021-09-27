package main

import (
	"github.com/rlgino/go-grcp-example/internal/platform"
	"github.com/rlgino/go-grcp-example/internal/platform/server"
	"log"
)

func main() {
	cfg := server.Config{
		Protocol: "tcp",
		Host:     "localhost",
		Port:     "3333",
	}

	srv := platform.NewServer(cfg)
	log.Printf("gRPC server running at %s://%s:%s ...\n", cfg.Protocol, cfg.Host, cfg.Port)
	log.Fatalln(srv.Serve())
}
