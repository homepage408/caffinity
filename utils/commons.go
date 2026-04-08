package utils

import "database/sql"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(status int, message string, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

type Empty struct{}

func NullStringToPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func NullFloat64ToPtr(nf sql.NullFloat64) *float64 {
	if !nf.Valid {
		return nil
	}
	return &nf.Float64
}

func NullInt32ToPtr(ni sql.NullInt32) *int32 {
	if !ni.Valid {
		return nil
	}
	return &ni.Int32
}

func NullTimeToPtr(nt sql.NullTime) *string {
	if !nt.Valid {
		return nil
	}
	str := nt.Time.Format("2006-01-02 15:04:05")
	return &str
}
