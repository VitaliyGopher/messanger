package server

func (s *server) configureRouter() {
	s.router.GET("/ping", s.Ping)
}
