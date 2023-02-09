package main

// Reference : YouTube c/LaithAcademy v/Build a REST API with Golang

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


type todo struct {
	ID string `json:"id"`
	Item string `json:"item"`
	Complted bool `json:"completed"`
}


var todos = []todo{
	{"1", "Clean Room", false},
	{"2", "Read Book", false},
	{"3", "Learn Go", false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	err := context.BindJSON((&newTodo))
	if err != nil {
		return 
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}


func getTodoById(id string) (*todo, error) {

	for i ,t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}


func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Todo not found"})
		return
	}
	todo.Complted = !todo.Complted
	context.IndentedJSON(http.StatusOK, todo)
}



func main() {

	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	
	router.Run("localhost:5000")
}