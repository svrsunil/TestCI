package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB
var err error

type response struct {
	Statuscode string `json:"status"`
	RefNo      string `json:"RefNo"`
	Name       string `json:"Name"`
	Msg        string `json:"Msg"`
	Errormsg   string `json:"error"`
}

func initDB() {
	//initiate DB
	dbDriver := "mysql"
	dbHost := "mysqldb"
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	fmt.Println(dbName, dbPass)
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName+"?parseTime=true")

	if err != nil {
		log.Println("Error: ", err.Error())
	} else {
		log.Println("DB connection successful !!", db.Ping())
	}

}

//basicAuth is a interceptor before accessing a secured URL
func basicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, pass, ok := r.BasicAuth()

		if ok && checkUsernameAndPassword(user, pass) {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Deny access for the the handle
			http.Error(w, "Not Authorized", 401)

			return
		}
	}
}

func checkUsernameAndPassword(username, password string) bool {
	return username == "abc" && password == "123"
}

func main() {

	initDB()

	log.Println("Server started to accept request in 8080 port.")

	router := setupRoutes()

	http.ListenAndServe(":8080", router)
}

func setupRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/getuser/:email", basicAuth(getUser))

	return router
}

// #http://deployci-env.5mj32iv4cz.ap-southeast-1.elasticbeanstalk.com:8080/api/getuser/test1@gmail.com
