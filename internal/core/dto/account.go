package dto

import "github.com/Long-Plan/longplan-api/internal/core/model"

type Account struct {
	UserData    model.Account  `json:"userData"`
	StudentData *model.Student `json:"studentData"`
}
