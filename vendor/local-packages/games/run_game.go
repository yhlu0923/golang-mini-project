package games

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitializeGames() {
	if err := godotenv.Load(); err != nil {
		//Do nothing
	}
	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	ResetDBTable(db)

	r := gin.Default()
	r.RedirectFixedPath = true
	// r.GET("/bookshelf", getBooks(db))
	// // [TODO] other method
	// r.GET("/bookshelf/:id", getBook(db))
	// r.POST("/bookshelf", addBook(db))
	// r.DELETE("/bookshelf/:id", deleteBook(db))
	// r.PUT("/bookshelf/:id", updateBook(db))

	r.Run(":" + port)
}

func ResetDBTable(db *sql.DB) {
	tmp_tableName := "games_Grab30"
	if _, err := db.Exec("DROP TABLE IF EXISTS " + tmp_tableName); err != nil {
		return
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS " + "tmp_tableName" + " (id SERIAL PRIMARY KEY, username VARCHAR(100), pages VARCHAR(10))"); err != nil {
		return
	}
}

// Guess num

//生成规定范围内的整数
//设置起始数字范围，0开始,endNum截止
func CreateRandomNumber(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}
