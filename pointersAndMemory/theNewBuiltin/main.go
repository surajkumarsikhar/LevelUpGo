package main

import "fmt"

type Config struct {
	Port      *int
	EnableTLS *bool
}

type Server struct {
	Port      int
	EnableTLS bool
	Addr      string // listen address, e.g. "localhost:8080"
}

// newServer defaults any unset Config field (Port to 8080, EnableTLS to true),
// then returns a Server holding the resolved values plus its listen Addr
// ("localhost:<port>"). Leave fields the caller already set unchanged.
func newServer(cfg *Config) *Server {
	// Your code here
	port := 8080
	if cfg.Port != nil {
		port = *cfg.Port
	}
	enableTLS := true
	if cfg.EnableTLS != nil {
		enableTLS = *cfg.EnableTLS
	}
	return &Server{
		Port:      port,
		EnableTLS: enableTLS,
		Addr:      fmt.Sprintf("localhost:%d", port),
	}
}

func (s *Server) start() {
	scheme := "http"
	if s.EnableTLS {
		scheme = "https"
	}
	fmt.Printf("listening on %s://%s\n", scheme, s.Addr)
}

func main() {
	srv := newServer(&Config{})
	srv.start()
}
