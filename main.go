// ginBlog project main.go
package main

import (
	_ "fmt"
	"ginBlog/routers"
	_ "net/http"

	_ "github.com/gin-gonic/gin"
)

func main() {

	router := routers.InitRouter()

	router.Run()
}
