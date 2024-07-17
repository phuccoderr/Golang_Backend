package database

type AbstractRepository[T any] interface {
	Get(id int) (*T, error)
	List(offset int, limit int) ([]*T, error)
	Create(t *T) error
	Update(t *T) error
}
