package config

import (
	"gorm.io/gorm"
)

// DB
var (
	db *gorm.DB
	logger *Logger
)

func Init() error{
	return nil
}

func GetLogger(p string) *Logger {
	// New Logger Initialize
	logger = NewLogger(p)
	return logger
}