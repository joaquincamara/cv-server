package coolfeatures

type Repository interface {
	FindAll() ([]*Coolfeatures, error)
	Add(coolfeatures *Coolfeatures) error
	Delete(id int) error
	Update(coolfeatures *Coolfeatures) error
}
