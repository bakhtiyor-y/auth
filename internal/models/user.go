package models

import "github.com/golang-jwt/jwt"

type CreateUser struct {
	AuthUser
	User
}

type User struct {
	Id                string
	Role              string
	Name              string
	Surname           string
	Phone             string
	Address           string
	AddressCoordinate Coordinate
}

type AuthUser struct {
	Login    string
	Password string
}

type Coordinate struct {
	X float64
	Y float64
}

type UserInfo struct {
	Login string
	Role  string
}

type CustomClaims struct {
	jwt.StandardClaims
	Login string `json:"username"`
	UserId  string `json:"userId"`
}
