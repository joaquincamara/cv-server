package devTechs

type Repository interface {
	FindAll() ([]*DevTech, error)
	Add(devTech *DevTech) error
	Delete(id int) error
	Update(devTech *DevTech) error
}
