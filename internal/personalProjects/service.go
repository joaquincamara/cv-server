package personalProjects

type Service interface {
	FindAll() ([]*PersonalProjects, error)
	Add(personalProjects *PersonalProjects) error
	Delete(id int) error
	Update(personalProjects *PersonalProjects) error
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

func (s *service) Delete(id int) error {
	return s.PersonalProjectsRepo.Delete(id)
}

func (s *service) Update(personalProjects *PersonalProjects) error {
	return s.PersonalProjectsRepo.Update(personalProjects)
}
