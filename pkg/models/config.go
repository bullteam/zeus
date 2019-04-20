package models

// Database ...
type Database struct {
	MySQL  *MySQL  `json:"mysql,omitempty"`
}

// MySQL ...
type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Database string `json:"database"`
}
