package repository

import "go-unit-test/entity"

type CategoryRepository interface {
	// Jika ada datanya di db, return category, jika tidak ada nil
	FindById(Id string) *entity.Category
}
