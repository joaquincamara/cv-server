package experience

type Service interface {
	FindAll() ([]*Experience, error)
	Add(experience *Experience) error
	Delete(id int) error
	Update(experience *Experience) error
}

type service struct {
	experienceRepo Repository
}

func NewExperienceService(experience Repository) Service {
	return &service{experienceRepo: experience}
}

func (s *service) FindAll() ([]*Experience, error) {
	return s.experienceRepo.FindAll()
}

func (s *service) Add(experience *Experience) error {
	return s.experienceRepo.Add(experience)
}

func (s *service) Delete(id int) error {
	return s.experienceRepo.Delete(id)
}

func (s *service) Update(experience *Experience) error {
	return s.experienceRepo.Update(experience)
}
