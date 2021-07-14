package coolFeatures

type Service interface {
	FindAll() ([]*Coolfeatures, error)
	Add(coolfeatures *Coolfeatures) error
	Delete(id int) error
	Update(coolfeatures *Coolfeatures) error
}

type service struct {
	CoolfeaturesRepo Repository
}

func NewCoolFeaturesService(coolfeaturesRepo Repository) Service {
	return &service{CoolfeaturesRepo: coolfeaturesRepo}
}

func (s *service) FindAll() ([]*Coolfeatures, error) {
	return s.CoolfeaturesRepo.FindAll()
}

func (s *service) Add(coolfeatures *Coolfeatures) error {
	return s.CoolfeaturesRepo.Add(coolfeatures)
}

func (s *service) Delete(id int) error {
	return s.CoolfeaturesRepo.Delete(id)
}

func (s *service) Update(Coolfeatures *Coolfeatures) error {
	return s.CoolfeaturesRepo.Update(Coolfeatures)
}
