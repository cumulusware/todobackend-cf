package todos

var todos = []Todo{
	{"First task"},
}

// Todo models a todo for the TodoBackend.
type Todo struct {
	Title string `json:"title"`
}
