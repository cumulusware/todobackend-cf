package todos

// DataStore provides the interface required to retrieve and save todos.
type DataStore interface {
	Create(*Todo) (string, error)
	GetAll(baseURL string) ([]Todo, error)
	GetByID(id, url string) (Todo, error)
	DeleteAll() error
}
