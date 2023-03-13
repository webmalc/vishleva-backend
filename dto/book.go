package dto

import "time"

// Book is the book request.
type Book struct {
	Name       string    `json:"name"`
	Comment    string    `json:"comment"  binding:"required,min=5"`
	Begin      time.Time `json:"begin" time_format:"2018-09-22T19:42:31+03:00"`
	End        time.Time `json:"end" time_format:"2018-09-22T19:42:31+03:00"`
	Phone      string    `json:"phone" binding:"required,numeric,len=11,startswith=7"`
	ClientName string    `json:"client_name"`
	Email      string    `json:"email,omitempty" binding:"omitempty,email"`
}
