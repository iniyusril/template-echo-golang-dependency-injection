package main

import "github.com/labstack/echo/v4"

type Server struct {
	Server *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	return &Server{e}
}
