package main


import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)



type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}



var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}



func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/add-todo", addTodo)
	router.Run("localhost:8080")
}



func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}



func addTodo(context *gin.Context) {
	var newTodo todo

	if error := context.BindJSON(&newTodo); error != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}



func getTodoByID(id string) (*todo, error) {
	for _, todoItem := range todos {
		if todoItem.ID == id {
			return &todoItem, nil
		}
	}

	return nil, errors.New("todo not found")
}



func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, error := getTodoByID(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}



func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}