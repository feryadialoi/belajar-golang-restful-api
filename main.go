package main

import (
	"feryadialoi/belajar-golang-restful-api/app"
	"feryadialoi/belajar-golang-restful-api/controller"
	"feryadialoi/belajar-golang-restful-api/helper"
	"feryadialoi/belajar-golang-restful-api/middleware"
	"feryadialoi/belajar-golang-restful-api/repository"
	"feryadialoi/belajar-golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
