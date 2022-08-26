package user_controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ghgsnaidu/go-db-demo/pkg/models"
	"github.com/ghgsnaidu/go-db-demo/pkg/services"
	"github.com/gorilla/mux"
)

var (
	id int64 = 0
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := services.GetUsers()
	n := len(users)
	for i := 0; i < n; i++ {
		writeToResponse(w, users[i])
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing ID")
		return
	}
	user, suc := services.GetUser(id)
	if !suc {
		w.WriteHeader(http.StatusNoContent)
		fmt.Println("id not found in data-base")
		return
	} else {
		writeToResponse(w, user)
	}

}
func PostUser(w http.ResponseWriter, r *http.Request) {
	id++
	var user models.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("error occured while unmarshaling ")
		fmt.Println(err)
		return
	}

	if !services.AddUser(id, user) {
		fmt.Println("cannot insert into data-base")
	}
}

func writeToResponse(w http.ResponseWriter, user models.User) {
	res, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error occured while marshaling")
		return
	}
	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
