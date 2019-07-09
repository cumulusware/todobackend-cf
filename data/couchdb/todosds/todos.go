package todosds

import (
	"context"
	"fmt"
	"time"

	"github.com/cumulusware/todobackend-cf/resources/todos"
	"github.com/flimzy/kivik"
)

// DataStore implements the DataStore interface for todos.
type DataStore struct {
	ctx context.Context
	DB  *kivik.DB
}

// NewDataStore creates a new DataStore.
func NewDataStore(ctx context.Context, c *kivik.Client) (*DataStore, error) {
	var ds DataStore
	dbName := "todos"
	// Check to see if the todos database already exists.
	dbExists, err := c.DBExists(context.TODO(), dbName, nil)
	if err != nil {
		return &ds, fmt.Errorf("error determining if %s db exists: %s", dbName, err)
	}
	// If the todos database doesn't exist, create it.
	if !dbExists {
		err = c.CreateDB(ctx, dbName, nil)
		if err != nil {
			return &ds, fmt.Errorf("error creating %s database: %s", dbName, err)
		}
		// TODO: Create the design documents for various queries.
	}
	db, err := c.DB(ctx, dbName, nil)
	if err != nil {
		return &ds, fmt.Errorf("error getting %s db handle: %s", dbName, err)
	}

	ds = DataStore{ctx, db}
	return &ds, nil
}

type todoDoc struct {
	ID        string `json:"_id,omitempty"`
	Rev       string `json:"_rev,omitempty"`
	Deleted   bool   `json:"_deleted,omitempty"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Create stores a new todo in the DataStore.
func (ds *DataStore) Create(todo *todos.Todo) (string, error) {
	docID, _, err := ds.DB.CreateDoc(ds.ctx, todo, nil)
	return docID, err
}

// GetAll returns all todos found in the DataStore.
func (ds *DataStore) GetAll(baseURL string) ([]todos.Todo, error) {
	var todos []todos.Todo

	// Get all the docs from the todos database.
	rows, err := ds.DB.AllDocs(ds.ctx, kivik.Options{"include_docs": true})
	if err != nil {
		return todos, fmt.Errorf("error getting all docs: %s", err)
	}

	// Loop through each row and create a todo from the doc, which is added to
	// the list of todos.
	for rows.Next() {
		var doc todoDoc
		if err := rows.ScanDoc(&doc); err != nil {
			return todos, fmt.Errorf("error scanning doc: %s", err)
		}
		todo := convertDocToTodo(doc)
		todos = append(todos, todo)
	}

	return todos, nil
}

// GetByID returns one todo found in the DataStore.
func (ds *DataStore) GetByID(id, url string) (todos.Todo, error) {
	var todo todos.Todo
	row, err := ds.DB.Get(ds.ctx, id, nil)
	if err != nil {
		return todo, fmt.Errorf("error getting doc with ID %s: %s", id, err)
	}
	var doc todoDoc
	if err := row.ScanDoc(&doc); err != nil {
		return todo, fmt.Errorf("error scanning doc: %s", err)
	}
	todo = convertDocToTodo(doc)
	todo.URL = url

	return todo, nil
}

// DeleteAll deletes all todos in the DataStore.
func (ds *DataStore) DeleteAll() error {

	// Get all docs.
	var docs []todoDoc
	rows, err := ds.DB.AllDocs(ds.ctx, kivik.Options{"include_docs": true})
	if err != nil {
		return fmt.Errorf("error getting all docs: %s", err)
	}

	// Iterate through each doc and set to deleted.
	for rows.Next() {
		var d todoDoc
		if err := rows.ScanDoc(&d); err != nil {
			return fmt.Errorf("error scanning doc: %s", err)
		}
		d.Deleted = true
		docs = append(docs, d)
	}

	// Bulk update all docs to be deleted.
	time.Sleep(300 * time.Millisecond) // Added for IBM Cloud rate limit on lite plan.
	_, err = ds.DB.BulkDocs(ds.ctx, docs, nil)
	return err
}

func convertDocToTodo(doc todoDoc) todos.Todo {
	return todos.Todo{
		Title:     doc.Title,
		Completed: doc.Completed,
	}
}
