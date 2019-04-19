package models

import (
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	FamilyName string `json:"familyName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	City       string `json:"city"`
	Photo      []byte `json:"userPhoto"`

	BirthDate         time.Time `json:"birthdate"`
	LastSeen          time.Time `json:"lastSeen"`
	LastProfileUpdate time.Time `json:"lastProfileUpdate"`

	Interests []Interest `json:"interests"`

	Friends  []int `json:"friendsID"`
	Programs []int `json:"programsID"`
	Posts    []int `json:"postsID"`
}
