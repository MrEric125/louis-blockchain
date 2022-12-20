package controller


import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type AuthController struct {
}

var db =make(map[string]string)
func (auth *AuthController) GetUserValue(c *gin.Context) {

	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}

}

func main() {
	//ConnectSql()
}

func (auth *AuthController)  ConnectSql() {
	db, err := sql.Open("mysql", "root:!QAZ2wsx@tcp(127.0.0.1:3306)/louis_coub?parseTime=true")

	if err != nil {
		fmt.Print("connect mysql error:", err.Error())
		return

	}
	rows, err2 := db.Query("select 1 ")

	if err2 != nil {
		fmt.Print("ping mysql error:", err.Error())
		return
	}
	fmt.Print(rows)
}
