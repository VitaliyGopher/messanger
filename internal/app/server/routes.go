package server

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func (s *server) configureRouter() {
	s.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.router.GET("/ping", s.Ping)

	s.router.POST("/sendCode", s.SendCodeHandler)
	s.router.POST("/createJwt", s.GetJWT)
	s.router.POST("/getNewJWT", s.GetNewJWT)
}
