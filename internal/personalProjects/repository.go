package personalprojects

type Repository interface {
	FindAll() ([]*PersonalProjects, error)
	Add(personalProjects *PersonalProjects) error
	Delete(personalProjects *PersonalProjects) error
}
