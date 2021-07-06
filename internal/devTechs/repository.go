package devTechs

type Repository interface {
	//	FindAll() ([]*DevTech, error)
	Add(devTech *DevTech) error
	Delete(devTech *DevTech) error
}
