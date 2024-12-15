package service

import (
	"CONTACTAPP/models"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)


var db *sql.DB
func InitializeDatabase(database *sql.DB){
	db = database
}

func SignIn(username,password string)(models.User,error){
	var user models.User
	query := `SELECT id, username, password FROM users WHERE username = $1 AND password = $2`
	err := db.QueryRow(query,username,password).Scan(&user,&user.UserName,&user.Password)
	if err != nil{
		if err == sql.ErrNoRows{
			return models.User{},errors.New("user not found")
		}
		return models.User{}, err
	}
	return user,nil
}


func AddUser(user models.User)(string,error){
	if checkUserNameExists(user.UserName){
		return "Username already exists",errors.New("username already exists")
	}
	user.Id = uuid.New()
	query := `INSERT INTO users(id,username,password) VALUES ($1,$2,$3)`
	_,err := db.Exec(query,user.Id,user.UserName,user.Password)
	if err != nil{
		return "",err
	}
	return "User added", nil
}

func DeleteUser(id uuid.UUID)(string,error){
	query := `DELETE FROM users WHERE id = $1`
	_,err := db.Exec(query,id)
	if err != nil{
		return "User not found", errors.New("user not found")
	}
	return "user deleted",nil
}

func UpdateUser(user models.User)(string, error){
	query := `UPDATE users SET username = $1, password = $2 WHERE id = $3`
	_,err := db.Exec(query,user.UserName,user.Password,user.Id)
	if err != nil{
		return "user not found, fatal", errors.New("user not found")
	}
	return "user updated",nil
}


func GetUser(id uuid.UUID)(models.User,error){
	var user models.User
	query := `SELECT id,username,password FROM users WHERE id = $1`
	err := db.QueryRow(query,id).Scan(&user.Id,&user.UserName,&user.Password)
	if err != nil{
		if err != sql.ErrNoRows{
			return models.User{}, errors.New("user not found")
		}
		return models.User{} , err
	}
	return user, nil
}


func checkUserNameExists(userName string)bool{
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)`
	err := db.QueryRow(query,userName).Scan(&exists)
	return err == nil && exists
}

func GetUsers() ([]models.User,error){
	query := `SELECT id, username, password FROM users`
	rows,err := db.Query(query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next(){
		var user models.User
		err := rows.Scan(&user.Id,&user.UserName,&user.Password)
		if err != nil{
			return nil, err
		}
		users = append(users, user)
	}
	return users,nil
}

