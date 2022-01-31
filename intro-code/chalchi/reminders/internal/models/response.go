package models

type ResponseOutput struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    *Output  `json:"data"`
	Errors  []string `json:"errors"`
}

type ResponseOutputs struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    *[]Output `json:"data"`
	Errors  []string  `json:"errors"`
}

type ResponseUser struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    *User    `json:"data"`
	Errors  []string `json:"errors"`
}

type ResponseUsers struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    *[]User  `json:"data"`
	Errors  []string `json:"errors"`
}

type ResponseSchedule struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    *Schedule `json:"data"`
	Errors  []string  `json:"errors"`
}

type ResponseSchedules struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *[]Schedule `json:"data"`
	Errors  []string    `json:"errors"`
}
