package dto

type StudentUpdateDto struct {
	MajorID        int  `json:"major_id,omitempty"`
	IsTermAccepted bool `json:"is_term_accepted,omitempty"`
}
