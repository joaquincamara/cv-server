package personalProjects

type Repository interface {
	FindAll() ([]*PersonalProjects, error)
	Add(personalProjects *PersonalProjects) error
	Delete(id int) error
	Update(personalProjects *PersonalProjects) error
}
