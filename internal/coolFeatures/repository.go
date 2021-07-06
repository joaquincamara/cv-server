package coolfeatures

type Repository interface {
	FindAll() ([]*Coolfeatures, error)
	Add(coolfeatures *Coolfeatures) error
	Delete(coolfeatures *Coolfeatures) error
}
