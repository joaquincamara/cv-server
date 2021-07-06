package personalprojects

type Service interface {
	FindAll() ([]*PersonalProjects, error)
	Add(personalProjects *PersonalProjects) error
	Delete(personalProjects *PersonalProjects) error
}

type service struct {
	PersonalProjectsRepo Repository
}

func NewCoolFeatureService(personalProjectsRepo Repository) Service {
	return &service{PersonalProjectsRepo: personalProjectsRepo}
}

func (s *service) FindAll() ([]*PersonalProjects, error) {
	return s.PersonalProjectsRepo.FindAll()
}

func (s *service) Add(personalProjects *PersonalProjects) error {
	return s.PersonalProjectsRepo.Add(personalProjects)
}

func (s *service) Delete(personalProjects *PersonalProjects) error {
	return s.PersonalProjectsRepo.Delete(personalProjects)
}
