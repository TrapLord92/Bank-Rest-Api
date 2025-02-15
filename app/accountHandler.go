package app

import (
	"encoding/json"
	"net/http"

	"github.com/TrapLord92/Simples-Go-Bank-Rest-Api/dto"
	"github.com/TrapLord92/Simples-Go-Bank-Rest-Api/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

// swagger:operation GET /customers/{customer_id}/account Accounts New
//
// Create a new account for customer.
//
// # This creates a new account for the given customer
//
// ---
// produces:
// - application/json
// parameters:
//   - name: request
//     in: body
//     description: tags to filter by
//     required: true
//     schema:
//     "$ref": "#/definitions/newAccountRequest"
//
// responses:
//
//	'201':
//	  description: gets all customers from database
//	  schema:
//	    "$ref": "#/definitions/newAccountResponse"
func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

// /customers/2000/accounts/90720
func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		//build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, appError := h.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}

}
