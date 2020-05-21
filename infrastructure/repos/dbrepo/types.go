package dbrepo

type Todo struct {
	ID   int    `db:"id"`
	Name string `db:"todo"`
}
