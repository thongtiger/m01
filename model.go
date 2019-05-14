package main

type ProfileForm struct {
	Username     string `json:"username" validate:"required"`
	EmailAddress string `json:"email" validate:"required,email"`
	Firstname    string `json:"firstname" validate:"required"`
	Lastname     string `json:"lastname" validate:"required"`
	Gender       string `json:"sex" validate:"required,len=1"`
	CountryCode  string `json:"country" validate:"required,len=2"`
	DateOfBirth  int    `json:"birth_day" validate:"required,numeric"`
	MonthOfBirth int    `json:"birth_month" validate:"required,numeric"`
	YearOfBirth  int    `json:"birth_year" validate:"required,numeric"`
}
