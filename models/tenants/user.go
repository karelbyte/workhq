package models

import (
	"elpuertodigital/workhq/dbg"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Uanony struct{}

type User struct {
	Id        uuid.UUID `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Status    uint16    `json:"status"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt *string   `json:"-"`
}

func (u *User) All() []User {

	selectQuery := "SELECT * FROM users;"

	res, err := dbg.DB().Query(selectQuery)

	if err != nil {
		panic(err.Error())
	}

	var users []User

	for res.Next() {

		user := User{}

		err := res.Scan(&user.Id, &user.FullName, &user.Email, &user.Address, &user.Phone,
			&user.Password, &user.Token, &user.Status, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users
}

func (u *User) Store() *User {

	id := uuid.New()

	hashPass := ""

	if u.Password != "" {
		hashPass, _ = HashPassword(u.Password)
	}

	insertQuery := fmt.Sprint("INSERT INTO users(id, fullname, address, email, phone, password, token, status) VALUES ('",
		id.String(), "','", u.FullName, "','", u.Address, "','", u.Email, "','", u.Phone, "','", hashPass, "','", u.Token, "',", u.Status, ");")

	_, err := dbg.DB().Exec(insertQuery)

	if err != nil {
		panic(err.Error())
	}

	u.Id = id

	u.Find()

	return u

}

func (u *User) Find() *User {

	findQuery := `SELECT * FROM users WHERE id = '` + u.Id.String() + "';"

	res := dbg.DB().QueryRow(findQuery)

	err := res.Scan(&u.Id, &u.FullName, &u.Email, &u.Address, &u.Phone,
		&u.Password, &u.Token, &u.Status, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)

	if err != nil {
		panic(err.Error())
	}

	return u

}

func (u *User) Update() *User {

	srtId := u.Id.String()

	updatedQuery := `UPDATE users SET`

	if u.FullName != "" {
		updatedQuery += ` fullname='` + u.FullName + `'`
	}

	if u.Address != "" {
		updatedQuery += `, address='` + u.Address + `'`
	}

	if u.Email != "" {
		updatedQuery += `, email='` + u.Email + `'`
	}

	if u.Phone != "" {
		updatedQuery += `, phone='` + u.Phone + `'`
	}

	if u.Status != 0 {
		updatedQuery += `, status=` + strconv.FormatInt(int64(u.Status), 10)
	}

	updatedQuery += ` WHERE id = '` + srtId + `';`

	_, err := dbg.DB().Exec(updatedQuery)

	if err != nil {
		panic(err.Error())
	}

	u.Find()

	return u

}

func (u *User) Delete() {

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UserResource(user User) interface{} {
	userResourcer := struct {
		Id        uuid.UUID `json:"id"`
		FullName  string    `json:"fullname"`
		Email     string    `json:"email"`
		Address   string    `json:"address"`
		Phone     string    `json:"phone"`
		Status    uint16    `json:"status"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
	}{
		Id:        user.Id,
		FullName:  user.FullName,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResourcer
}

func UserResourceCollection(users []User) []interface{} {

	var resources []interface{}

	for _, user := range users {
		resources = append(resources, UserResource(user))
	}

	return resources
}
