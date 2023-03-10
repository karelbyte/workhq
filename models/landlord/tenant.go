package landlordmodels

import (
	"elpuertodigital/workhq/dbg"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type Uanony struct{}

type Tenant struct {
	Id        uuid.UUID `json:"id"`
	Domain    string    `json:"fullname"`
	Db        string    `json:"email"`
	Status    uint16    `json:"status"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt *string   `json:"-"`
}

func (t *Tenant) All() []Tenant {

	selectQuery := "SELECT * FROM Tenants;"

	res, err := dbg.DB().Query(selectQuery)

	if err != nil {
		panic(err.Error())
	}

	var Tenants []Tenant

	for res.Next() {

		Tenant := Tenant{}

		err := res.Scan(&Tenant.Id, &Tenant.Domain, &Tenant.Db, &Tenant.Status, &Tenant.CreatedAt, &Tenant.UpdatedAt, &Tenant.DeletedAt)

		if err != nil {
			panic(err.Error())
		}

		Tenants = append(Tenants, Tenant)
	}

	return Tenants
}

func (t *Tenant) Store() *Tenant {

	id := uuid.New()

	insertQuery := fmt.Sprint("INSERT INTO Tenants(id, domain, db, status) VALUES ('",
		id.String(), "','", t.Domain, "','", t.Db, "','", t.Status, "',", t.Status, ");")

	_, err := dbg.DB().Exec(insertQuery)

	if err != nil {
		panic(err.Error())
	}

	t.Id = id

	t.Find()

	return t

}

func (t *Tenant) Find() *Tenant {

	findQuery := `SELECT * FROM Tenants WHERE id = '` + t.Id.String() + "';"

	res := dbg.DB().QueryRow(findQuery)

	err := res.Scan(&t.Id, &t.Domain, &t.Db, &t.Status, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt)

	if err != nil {
		panic(err.Error())
	}

	return t

}

func (t *Tenant) Update() *Tenant {

	srtId := t.Id.String()

	updatedQuery := `UPDATE Tenants SET`

	if t.Domain != "" {
		updatedQuery += ` domain='` + t.Domain + `'`
	}

	if t.Db != "" {
		updatedQuery += `, db='` + t.Db + `'`
	}

	if t.Status != 0 {
		updatedQuery += `, status=` + strconv.FormatInt(int64(t.Status), 10)
	}

	updatedQuery += ` WHERE id = '` + srtId + `';`

	_, err := dbg.DB().Exec(updatedQuery)

	if err != nil {
		panic(err.Error())
	}

	t.Find()

	return t

}

func (t *Tenant) Delete() {

}

func TenantResource(Tenant Tenant) interface{} {
	TenantResourcer := struct {
		Id        uuid.UUID `json:"id"`
		Domain    string    `json:"domain"`
		Db        string    `json:"db"`
		Status    uint16    `json:"status"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
	}{
		Id:        Tenant.Id,
		Domain:    Tenant.Domain,
		Db:        Tenant.Db,
		Status:    Tenant.Status,
		CreatedAt: Tenant.CreatedAt,
		UpdatedAt: Tenant.UpdatedAt,
	}

	return TenantResourcer
}

func TenantResourceCollection(Tenants []Tenant) []interface{} {

	var resources []interface{}

	for _, Tenant := range Tenants {
		resources = append(resources, TenantResource(Tenant))
	}

	return resources
}
