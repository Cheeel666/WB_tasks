package main

import (
	"dev11/pkg/handler"
	"dev11/pkg/server"
)

func main() {
	// Можно было бы обойтись без объекта handlers, но сначала сделал с использованием gin и потом увидел в задании,
	// что юзается только стандартная библиотека. В данном случае просто инициализируем ручки в handlers.InitRoutes(), они
	// автоматически прикручиваются к объекту srv.
	handlers := new(handler.Handler)
	handlers.InitRoutes()
	srv := new(server.Server)
	srv.Run("8080")
}
