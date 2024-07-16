package database

type AbstractRepository[T any] interface {
	Get(id int) (*T, error)
	List() ([]*T, error)
	Create(t *T) error
	Update(t *T) error
}
