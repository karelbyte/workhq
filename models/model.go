package models

import "encoding/json"

type IdModel struct {
	Id string `json:"id"`
}

type Moment struct {
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt *string   `json:"-"`
}

type Err struct {
  Status  string  `json:"status"`
  Message string `json:"message"`
}

type Model interface {
  All()
  Store()
  List()
  Update()
  Find()
  Delete()
}

func ManagerError(err error) []byte{
  error := Err{
    Status: "error",
    Message: err.Error(),
  }
  res, _ := json.Marshal(error)
  return res
}
