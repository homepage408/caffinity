package utils

import (
	"database/sql"
	"fmt"
)

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

func ParseStringToInt32(s string, defaultValue int) int32 {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
		return int32(defaultValue)
	}
	return int32(result)
}
