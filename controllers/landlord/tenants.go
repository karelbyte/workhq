package landlordcontrollers

import (
	"elpuertodigital/workhq/models"
	landlordmodels "elpuertodigital/workhq/models/landlord"
	"encoding/json"
	_ "fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func TenantIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	tenant := landlordmodels.Tenant{}

	tenants := tenant.All()

	tenantCollection := landlordmodels.TenantResourceCollection(tenants)

	data, _ := json.Marshal(tenantCollection)

	w.Write(data)

}


func TenantStore(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	tenant := landlordmodels.Tenant{}

	decodeErr := json.NewDecoder(r.Body).Decode(&tenant)

	if decodeErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		error := models.ManagerError(decodeErr)
		w.Write(error)
		return
	}

	tenant.Store()

	resource := landlordmodels.TenantResource(tenant)

	res, _ := json.Marshal(resource)

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func TenantShow(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	parameterId := p.ByName("id")

	tenantId, err := uuid.Parse(parameterId)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		error := models.ManagerError(err)
		w.Write(error)
		return
	}

	tenant := landlordmodels.Tenant{
		Id: tenantId,
	}

	tenant.Find()

	resource := landlordmodels.TenantResource(tenant)

	res, _ := json.Marshal(resource)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func TenantUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	parameterId := p.ByName("id")

	tenantId, err := uuid.Parse(parameterId)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		error := models.ManagerError(err)
		w.Write(error)
		return
	}

	tenant := landlordmodels.Tenant{
		Id: tenantId,
	}

	decodeErr := json.NewDecoder(r.Body).Decode(&tenant)

	if decodeErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		error := models.ManagerError(decodeErr)
		w.Write(error)
		return
	}
	
	tenant.Update()

	resource := landlordmodels.TenantResource(tenant)

	res, _ := json.Marshal(resource)

	w.Write(res)

}

