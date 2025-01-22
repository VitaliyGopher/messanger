package server

func (s *server) configureRouter() {
	s.router.GET("/ping", s.Ping)

	s.router.POST("/sendSms", s.SendSmsCodeHandler)
}
