package dbrepo

type Todo struct {
	ID   string `db:"id"`
	Name string `db:"todo"`
}
