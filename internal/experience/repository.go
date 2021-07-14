package experience

type Repository interface {
	FindAll() ([]*Experience, error)
	Add(experience *Experience) error
	Delete(id int) error
	Update(experience *Experience) error
}
