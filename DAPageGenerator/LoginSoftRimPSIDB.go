package main

import (
	"database/sql"
	"fmt"

	"github.com/go-ini/ini"
)

func LoginSoftRimPSIDB() (*sql.DB, error) {
	cfg, err := ini.Load("ini/db.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
	}

	user, pwd := cfg.Section("login").Key("user").String(), cfg.Section("login").Key("password").String()
	return sql.Open("mysql", user+":"+pwd+"@/SoftRimPSI?charset=utf8mb4")
}
