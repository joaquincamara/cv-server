package devTechs

type Service interface {
	FindAll() ([]*DevTech, error)
	Add(devTech *DevTech) error
	Delete(id int) error
	Update(devTech *DevTech) error
}

type service struct {
	DevTechRepo Repository
}

func NewDevTechService(devTechRepo Repository) Service {
	return &service{DevTechRepo: devTechRepo}
}

func (s *service) Add(devTech *DevTech) error {
	return s.DevTechRepo.Add(devTech)
}

func (s *service) FindAll() ([]*DevTech, error) {
	return s.DevTechRepo.FindAll()
}

func (s *service) Delete(id int) error {
	return s.DevTechRepo.Delete(id)
}
func (s *service) Update(devTech *DevTech) error {
	return s.DevTechRepo.Update(devTech)
}
