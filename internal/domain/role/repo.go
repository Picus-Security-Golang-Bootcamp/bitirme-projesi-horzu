package role

import (
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) Migration() {
	r.db.AutoMigrate(&Role{})
}

func (r *RoleRepository) Create(role *Role) error {

	result := r.db.Create(&role)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *RoleRepository) InserSampleData() error {
	roles := []Role{
		{
			Id:   1,
			Role: "admin",
		},
		{
			Id:   2,
			Role: "user",
		},
	}
	for _, role := range roles {
		result := r.db.Create(&role)

		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
