package repository

import "gorm.io/gorm"

type PostRepo struct {
	Database *gorm.DB
}
