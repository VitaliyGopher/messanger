package server

func (s *server) configureRouter() {
	s.router.GET("/ping", s.Ping)

	s.router.POST("/sendCode", s.SendCodeHandler)
	s.router.POST("/createJwt", s.GetJWT)
}
