package repository

import "crud_mysql_api/infras"

type Repository interface {
	UserRepository
	ProductRepository
	VariantRepository
}

type RepositoryImpl struct {
	DB *infras.Conn
}

func ProvideRepo(db *infras.Conn) *RepositoryImpl {
	return &RepositoryImpl{
		DB: db,
	}
}
