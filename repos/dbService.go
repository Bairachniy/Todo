package repos


type Repo interface {
	GetAll() ([]string,error)
	Create(todo string) error
}
