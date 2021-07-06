package aboutme

type Repository interface {
	Add(aboutMe *AboutMe) error
	FindAll() ([]*AboutMe, error)
	Update(aboutMe *AboutMe) error
	Delete(aboutMe *AboutMe) error
}
