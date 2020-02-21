package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UploadFile :
func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			msg := fmt.Errorf("%v", err).Error()
			log.Println("Came here", msg)
			log.Println(err)
			byteArr, _ := json.Marshal(&response{Statuscode: "400", Errormsg: msg})
			w.Write(byteArr)
			w.WriteHeader(http.StatusBadRequest)
		}
	}()

	email := p.ByName("email")

	message := "User not Found"

	username := ""

	if email == "test1@gmail.com" {
		username = "Innov Lab"
	}

	fmt.Println(db)
	err = db.QueryRow("SELECT name FROM POCDB.USER_TABLE WHERE email=?", email).Scan(&username)

	if err != nil {
		log.Println("Error in SQL:", err.Error())
	}

	//log.Println(db.)

	if len(username) > 0 {
		message = "Welcome to" + username + " !"
	}
	var byteArr []byte

	byteArr, _ = json.Marshal(&response{Statuscode: "200", Name: email, Msg: message})

	w.Write(byteArr)

}

// Welcome :
func Welcome(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			msg := fmt.Errorf("%v", err).Error()
			log.Println("Came here", msg)
			log.Println(err)
			byteArr, _ := json.Marshal(&response{Statuscode: "400", Errormsg: msg})
			w.Write(byteArr)
			w.WriteHeader(http.StatusBadRequest)
		}
	}()

	var byteArr []byte

	byteArr, _ = json.Marshal(&response{Statuscode: "200", Name: "", Msg: "Welcome Team."})

	w.Write(byteArr)

}
