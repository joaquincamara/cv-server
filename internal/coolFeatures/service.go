package coolfeatures

type Service interface {
	FindAll() ([]*Coolfeatures, error)
	Add(coolfeatures *Coolfeatures) error
	Delete(coolfeatures *Coolfeatures) error
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

func (s *service) Delete(coolfeatures *Coolfeatures) error {
	return s.CoolfeaturesRepo.Delete(coolfeatures)
}
