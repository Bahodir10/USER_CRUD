package api

import (
	"CONTACTAPP/models"
	"CONTACTAPP/service"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)
func GetAllUsers(w http.ResponseWriter, r *http.Request){
	users,err := service.GetUsers()
	if err != nil{
		http.Error(w,"Failed to get users from data base",http.StatusInternalServerError)
		return
	}
	userBytes,err := json.Marshal(users)
	if err != nil{
		http.Error(w,"Failed marshal users",http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)
}


func GetUser(w http.ResponseWriter, r *http.Request){
	userId := r.URL.Query().Get("id")
	if userId == ""{
		http.Error(w,"need user id",http.StatusBadRequest)
		return
	}
	userUUID,err := uuid.Parse(userId)
	if err != nil{
		http.Error(w,"Invalid UUID",http.StatusBadRequest)
		return
	}

	user,err := service.GetUser(userUUID)
	if err != nil{
		http.Error(w,"User not found",http.StatusNotFound)
		return
	}
	userBytes,err := json.Marshal(user)
	if err != nil{
		http.Error(w,"Failed to marshal user Data",http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)
}


func SignUp(w http.ResponseWriter, r *http.Request){
	var user models.User
	userBytes,err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w,"Failed to read the request body",http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(userBytes,&user)
	if err != nil{
		http.Error(w,"Failed to parse the request body",http.StatusBadRequest)
		return
	}
	info,err := service.AddUser(user)
	if err != nil{
		http.Error(w,err.Error(),http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(info))
}



func UpdateUser(w http.ResponseWriter, r *http.Request){
	userId := r.URL.Query().Get("id")
	if userId == ""{
		http.Error(w, "need userId",http.StatusBadRequest)
		return
	}

	userUUID,err := uuid.Parse(userId)
	if err != nil{
		http.Error(w,"Invalid uuid",http.StatusBadRequest)
	}

	userBytes,err := io.ReadAll(r.Body)
	if err != nil{
		http.Error(w,"Failed while reading the request body",http.StatusInternalServerError)
		return
	}
	var user models.User
	err = json.Unmarshal(userBytes,&user)
	if err!= nil{
		http.Error(w, "Failed to reques body",http.StatusBadRequest)
		return
	}
	user.Id = userUUID
	info,err := service.UpdateUser(user)
	if err != nil{
		http.Error(w,err.Error(),http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(info))
}







func DeleteUser(w http.ResponseWriter, r *http.Request){
	userId := r.URL.Query().Get("id")
	if userId == ""{
		http.Error(w,"need id of User",http.StatusBadRequest)
		return
	}
	userUUID,err := uuid.Parse(userId)
	if err != nil{
		http.Error(w,"Invalid user id",http.StatusBadRequest)
		return
	}

	info,err := service.DeleteUser(userUUID)
	if err != nil{
		http.Error(w,err.Error(),http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(info))
}

