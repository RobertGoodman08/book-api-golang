package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct { // models table 
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{ // json 
	{ID: "1", Title: "Head First. Изучаем Go", Author: "Макгаврен Джей", Quantity: 1},
	{ID: "2", Title: "Black Hat Python", Author: "Джастин Зайтц", Quantity: 7},
	{ID: "3", Title: "Сияние", Author: "Стивен Кинг", Quantity: 12},
}

func getBooks(c *gin.Context) { // Список книг
	c.IndentedJSON(http.StatusOK, books)
}


func bookById(c *gin.Context) { //  Получение книги по идентификатору
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Книга не найдена"}) // 404 page 
		return 
	}

	c.IndentedJSON(http.StatusOK, book)
}



func checkoutBook(c *gin.Context) { // Оформить заказ книг и параметры запроса
	id, ok := c.GetQuery("id")


	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Отсутствующий параметр запроса идентификатора"})
		return 
	}


	book, err := getBookById(id)


	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Книга не найдена"}) // 404 page 
		return 
	}


	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Книга недоступна"}) 
		return 
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}





func returnBook(c *gin.Context) { // Возврат книги
	id, ok := c.GetQuery("id")


	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Отсутствующий параметр запроса идентификатора"})
		return 
	}


	book, err := getBookById(id)


	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Книга не найдена"}) // 404 page 
		return 
	}


	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}




func getBookById(id string) (*book, error) { // проверка существует ли id
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil 
		}
	}

	return nil, errors.New("Книга не найдена")
}



func createBook(c *gin.Context) { // Создание книг
	var newBook book 

	if err := c.BindJSON(&newBook); err != nil {
		return 
	}


	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


func deleteBook(c *gin.Context) { // Удаление книги
	id := c.Param("id")


	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}


	c.IndentedJSON(http.StatusOK, books)
}



func urls() { // маршрутизатор
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.GET("/delete/:id", deleteBook)
	router.Run("localhost:8080")
}



func main() {

	urls()	
}
