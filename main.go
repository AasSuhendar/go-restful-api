package main

import (
	"fmt"
	"net/http"

	"github.com/AasSuhendar/go-restful-api/controller"
	"github.com/AasSuhendar/go-restful-api/helper"
	"github.com/AasSuhendar/go-restful-api/pkg/databases"
	"github.com/AasSuhendar/go-restful-api/pkg/execption"
	"github.com/AasSuhendar/go-restful-api/repository"
	"github.com/AasSuhendar/go-restful-api/services"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := databases.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = execption.ErrorHandler

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: router,
	}
	fmt.Println("Server Running")
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
