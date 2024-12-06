package responses

import (
	"time"
)

type ResponseUser struct {
	Id           uint64    `json:"id"`
	Name         string    `json:"name"`
	SolidBalance uint64    `json:"solid_balance"`
	StockBalance uint64    `json:"stock_balance"`
	IsBlocked    bool      `json:"is_blocked"`
	LastFarming  time.Time `json:"last_farming"`
	CreatedAt    time.Time `json:"created_at"`
}

type SingUpResponseOK struct {
	User ResponseUser `json:"user"`
}

type ErrorResponse struct {
	ErrorName string `json:"error_name"`
	Error     string `json:"error"`
}
type SingUpResponseError struct {
	ErrorResponse
}

type SingInResponseError struct {
	ErrorResponse
}

type FarmResponseOK struct {
	User   ResponseUser `json:"user"`
	Amount uint64       `json:"amount"`
}

type FarmResponseError struct {
	User ResponseUser `json:"user"`
	ErrorResponse
}

type BuyStocksResponseOK struct {
	User ResponseUser `json:"user"`
}

type BuyStocksResponseError struct {
	User ResponseUser `json:"user"`
	ErrorResponse
}

type UpdateNameResponseOK struct {
	User ResponseUser `json:"user"`
}

type UpdateNameResponseError struct {
	User ResponseUser `json:"user"`
	ErrorResponse
}

type UpdatePasswordResponseOK struct {
	User ResponseUser `json:"user"`
}

type UpdatePasswordResponseError struct {
	User ResponseUser `json:"user"`
	ErrorResponse
}

type BlockResponseOK struct {
	User ResponseUser `json:"user"`
}

type BlockResponseError struct {
	User ResponseUser `json:"user"`
	ErrorResponse
}

type UnblockResponseOK struct {
	User ResponseUser `json:"user"`
}

type UnlockResponseError struct {
	User ResponseUser `json:"user"`
	ErrorResponse
}