package experience

type Repository interface {
	FindAll() ([]*Experience, error)
	Add(experience *Experience) error
	Delete(experience *Experience) error
}
