package repository

import "projects/mock"

type CategoryRepository interface {
	FindById(id string) *mock.Category
}
