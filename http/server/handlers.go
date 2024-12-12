package server

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/VandiKond/StocksBack/config/requests"
	"github.com/VandiKond/StocksBack/config/responses"
	"github.com/VandiKond/StocksBack/config/user_cfg"
	"github.com/VandiKond/StocksBack/pkg/user_service"
	"github.com/VandiKond/vanerrors"
)

// The errors
const (
	WrongMethod = "wrong method"
	InvalidBody = "invalid body"
)

// User to response user
func ToResponseUser(usr user_cfg.User) responses.ResponseUser {
	return responses.ResponseUser{
		Id:           usr.Id,
		Name:         usr.Name,
		SolidBalance: usr.SolidBalance,
		StockBalance: usr.StockBalance,
		IsBlocked:    usr.IsBlocked,
		LastFarming:  usr.LastFarming,
		CreatedAt:    usr.CreatedAt,
	}
}

// Vanerror to response error
func ToErrorResponse(err error) responses.ErrorResponse {
	return responses.ErrorResponse{
		ErrorName: vanerrors.GetName(err),
		Error:     vanerrors.GetMessage(err),
	}
}

// Main page
func (h *Handler) MainHandler(w http.ResponseWriter, r *http.Request) {
	// Gets all pages
	pages := []string{}
	for fn := range h.funcs {
		pages = append(pages, fn)
	}
	slices.Sort[[]string, string](pages)
	// Sends data
	json.NewEncoder(w).Encode(responses.MainResponse{
		Pages: pages,
	})
}

// It creates a new user
func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	// Gets body
	req := requests.SignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		// Creates an error
		resp := vanerrors.NewSimple(InvalidBody)

		// Writes data
		responses.SignUpResponseError{
			ErrorResponse: ToErrorResponse(resp),
		}.
			SendJson(w, http.StatusBadRequest)

		return
	}

	// Signs up
	usr, err := req.SignUp(h.db)

	if err != nil {

		// Checks error variants
		var status = http.StatusBadRequest
		if user_service.IsServerError(err) {

			status = http.StatusInternalServerError
		}

		// Writes data
		responses.SignUpResponseError{
			ErrorResponse: ToErrorResponse(err),
		}.
			SendJson(w, status)

		h.logger.Warnf("unable to Sign up, reason: %v", err)
		return
	}

	// Converts user
	resp := ToResponseUser(*usr)

	// Sends data
	json.NewEncoder(w).Encode(responses.SignUpResponseOK{
		User: resp,
	})

	h.logger.Printf("Sign up: %v", *usr)
}

// Farms
func (h *Handler) FarmHandler(w http.ResponseWriter, r *http.Request, u user_cfg.User) {
	// Farming
	amount, usr, err := user_service.Farm(u.Id, h.db)

	if err != nil {
		// Checks error variants
		var status = http.StatusBadRequest
		if user_service.IsServerError(err) {

			status = http.StatusInternalServerError
		}
		// Writes data
		responses.SignUpResponseError{
			ErrorResponse: ToErrorResponse(err),
		}.
			SendJson(w, status)

		h.logger.Warnf("%v unable to farm, reason: %v", u, err)

		return
	}

	// Converts user
	resp := ToResponseUser(*usr)

	// Sends data
	json.NewEncoder(w).Encode(responses.FarmResponseOK{
		User:   resp,
		Amount: amount,
	})

	h.logger.Printf("farm (%d) : %v", amount, *usr)
}

// Buy stocks
func (h *Handler) BuyStocksHandler(w http.ResponseWriter, r *http.Request, u user_cfg.User) {
	// Gets body
	var req requests.BuyStocksRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		// Creates an error
		resp := vanerrors.NewSimple(InvalidBody)

		// Writes data
		responses.BuyStocksResponseError{
			ErrorResponse: ToErrorResponse(resp),
		}.
			SendJson(w, http.StatusBadRequest)

		return
	}

	// Buying stocks
	usr, err := user_service.BuyStocks(u.Id, req.Num, h.db)

	if err != nil {
		// Checks error variants
		var status = http.StatusBadRequest
		if user_service.IsServerError(err) {

			status = http.StatusInternalServerError
		}
		// Writes data
		responses.BuyStocksResponseError{
			ErrorResponse: ToErrorResponse(err),
		}.
			SendJson(w, status)

		h.logger.Warnf("%v unable to buy stocks, reason: %v", u, err)

		return
	}

	// Converts user
	resp := ToResponseUser(*usr)

	// Sends data
	json.NewEncoder(w).Encode(responses.BuyStocksResponseOK{
		User: resp,
	})

	h.logger.Printf("buy stocks (%d) : %v", req.Num, *usr)
}

// Update name
func (h *Handler) UpdateNameHandler(w http.ResponseWriter, r *http.Request, u user_cfg.User) {
	// Gets body
	var req requests.UpdateNameRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		// Creates an error
		resp := vanerrors.NewSimple(InvalidBody)

		// Writes data
		responses.UpdateNameResponseError{
			ErrorResponse: ToErrorResponse(resp),
		}.
			SendJson(w, http.StatusBadRequest)

		return
	}

	usr, err := user_service.UpdateName(u.Id, req.Name, h.db)

	if err != nil {
		// Checks error variants
		var status = http.StatusBadRequest
		if user_service.IsServerError(err) {

			status = http.StatusInternalServerError
		}
		// Writes data
		responses.UpdateNameResponseError{
			ErrorResponse: ToErrorResponse(err),
		}.
			SendJson(w, status)

		h.logger.Warnf("%v unable to update name, reason: %v", u, err)

		return
	}

	// Converts user
	resp := ToResponseUser(*usr)

	// Sends data
	json.NewEncoder(w).Encode(responses.UpdateNameResponseOK{
		User: resp,
	})

	h.logger.Printf("update name (was %s) : %v", u.Name, *usr)

}

// Update password
func (h *Handler) UpdatePasswordHandler(w http.ResponseWriter, r *http.Request, u user_cfg.User) {
	// Gets body
	var req requests.UpdatePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		// Creates an error
		resp := vanerrors.NewSimple(InvalidBody)

		// Writes data
		responses.UpdatePasswordResponseError{
			ErrorResponse: ToErrorResponse(resp),
		}.
			SendJson(w, http.StatusBadRequest)
		return
	}

	usr, err := user_service.UpdatePassword(u.Id, req.Password, h.db)

	if err != nil {
		// Checks error variants
		var status = http.StatusBadRequest
		if user_service.IsServerError(err) {

			status = http.StatusInternalServerError
		}
		// Writes data
		responses.UpdatePasswordResponseError{
			ErrorResponse: ToErrorResponse(err),
		}.
			SendJson(w, status)

		h.logger.Warnf("%v unable to update password, reason: %v", u, err)

		return
	}

	// Converts user
	resp := ToResponseUser(*usr)

	// Sends data
	json.NewEncoder(w).Encode(responses.UpdatePasswordResponseOK{
		User: resp,
	})

	h.logger.Printf("update password: %v", u.Name, *usr)
}

// Block user
func (h *Handler) BlockHandler(w http.ResponseWriter, r *http.Request, u user_cfg.User) {
	// Checking the key header (without it not allowed)
	key := r.Header.Get("Key")

	if key == "" {
		// Creates an error
		resp := vanerrors.NewSimple(InvalidHeader)

		// Writes data
		responses.BlockResponseError{
			ErrorResponse: ToErrorResponse(resp),
		}.
			SendJson(w, http.StatusForbidden)

		return
	}

	// Gets body
	var req requests.BlockRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		// Creates an error
		resp := vanerrors.NewSimple(InvalidBody)

		// Writes data
		responses.BlockResponseError{
			ErrorResponse: ToErrorResponse(resp),
		}.
			SendJson(w, http.StatusBadRequest)

		return
	}

	usr, err := user_service.Block(u.Id, h.db)

	if err != nil {
		// Checks error variants
		var status = http.StatusBadRequest
		if user_service.IsServerError(err) {

			status = http.StatusInternalServerError
		}
		// Writes data
		responses.BlockResponseError{
			ErrorResponse: ToErrorResponse(err),
		}.
			SendJson(w, status)

		h.logger.Warnf("%v unable to block, reason: %v", u, err)

		return
	}

	// Converts user
	resp := ToResponseUser(*usr)

	// Sends data
	json.NewEncoder(w).Encode(responses.BlockResponseOK{
		User: resp,
	})

	h.logger.Printf("block: %v", *usr)
}

// Unlock user
func (h *Handler) UnblockHandler(w http.ResponseWriter, r *http.Request, u user_cfg.User) {
	// Checking the key header (without it not allowed)
	key := r.Header.Get("Key")

	if key == "" {
		// Creates an error
		resp := vanerrors.NewSimple(InvalidHeader)

		// Writes data
		responses.BlockResponseError{
			ErrorResponse: ToErrorResponse(resp),
		}.
			SendJson(w, http.StatusForbidden)

		return
	}

	// Gets body
	var req requests.UnblockRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		// Creates an error
		resp := vanerrors.NewSimple(InvalidBody)

		// Writes data
		responses.UnblockResponseError{
			ErrorResponse: ToErrorResponse(resp),
		}.
			SendJson(w, http.StatusBadRequest)
		return
	}

	usr, err := user_service.Unblock(u.Id, h.db)

	if err != nil {
		// Checks error variants
		var status = http.StatusBadRequest
		if user_service.IsServerError(err) {

			status = http.StatusInternalServerError
		}
		// Writes data
		responses.UnblockResponseError{
			ErrorResponse: ToErrorResponse(err),
		}.
			SendJson(w, status)

		h.logger.Warnf("%v unable to unblock, reason: %v", u, err)

		return
	}

	// Converts user
	resp := ToResponseUser(*usr)

	// Sends data
	json.NewEncoder(w).Encode(responses.UnblockResponseOK{
		User: resp,
	})

	h.logger.Printf("unblock: %v", *usr)
}

// Get's user
func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	// Gets body
	var req requests.GetRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		// Creates an error
		resp := vanerrors.NewSimple(InvalidBody)

		// Writes data
		responses.UpdateNameResponseError{
			ErrorResponse: ToErrorResponse(resp),
		}.
			SendJson(w, http.StatusBadRequest)

		return
	}

	usr, err := user_service.Get(req.Id, h.db)

	if err != nil {
		// Checks error variants
		var status = http.StatusBadRequest
		if user_service.IsServerError(err) {

			status = http.StatusInternalServerError
		}
		// Writes data
		responses.UpdateNameResponseError{
			ErrorResponse: ToErrorResponse(err),
		}.
			SendJson(w, status)

		h.logger.Warnf("user %d not got, reason:", req.Id, err)

		return
	}

	// Converts user
	resp := ToResponseUser(*usr)

	// Sends data
	json.NewEncoder(w).Encode(responses.UpdateNameResponseOK{
		User: resp,
	})

	h.logger.Printf("sended user: %v", *usr)
}
