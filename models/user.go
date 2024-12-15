package models

import "github.com/google/uuid"

type User struct {
    UserName string    `json:"username"`
    Password string    `json:"password"` 
    Id       uuid.UUID `json:"id"`       
}
