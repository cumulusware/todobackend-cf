package todos

// DataStore provides the interface required to retrieve and save todos.
type DataStore interface {
	Create(*Todo) (string, error)
	GetAll(baseURL string) ([]Todo, error)
	GetByID(id, url string) (Todo, error)
	UpdateByID(id string, todo *Todo) error
	DeleteAll() error
	DeleteByID(id string) error
}
