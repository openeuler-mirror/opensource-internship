package models

type Condition struct {
	Id          int    `json:"id"`
	LicenseName string `json:"license_name,omitempty"`
}
