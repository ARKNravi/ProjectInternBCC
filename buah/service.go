package buah

type Service interface {
	FindAll() ([]Buah, error)
	FindByID(ID int) (Buah, error)
	Create(buahRequest BuahRequest) (Buah, error)
	Update(ID int, buahRequest BuahRequest) (Buah, error)
	Delete(ID int) (Buah, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Buah, error) {
	buahs, err := s.repository.FindAll()
	return buahs, err
}

func (s *service) FindByID(ID int) (Buah, error) {
	buah, err := s.repository.FindByID(ID)
	return buah, err
	//return s.repository.FindAll()
}

func (s *service) Create(buahRequest BuahRequest) (Buah, error) {

	buah := Buah{
		Jenis:       buahRequest.Jenis,
		Price:       buahRequest.Price,
		Description: buahRequest.Description,
		Matang:      buahRequest.Matang,
		Discount:    buahRequest.Discount,
	}

	newBuah, err := s.repository.Create(buah)

	return newBuah, err

	//return s.repository.FindAll()
}

func (s *service) Update(ID int, buahRequest BuahRequest) (Buah, error) {

	buah, err := s.repository.FindByID(ID)

	buah.Jenis = buahRequest.Jenis
	buah.Price = buahRequest.Price
	buah.Description = buahRequest.Description
	buah.Matang = buahRequest.Matang
	buah.Discount = buahRequest.Discount

	newBuah, err := s.repository.Update(buah)

	return newBuah, err

	//return s.repository.FindAll()
}

func (s *service) Delete(ID int) (Buah, error) {

	buah, err := s.repository.FindByID(ID)
	newBuah, err := s.repository.Delete(buah)
	return newBuah, err

	//return s.repository.FindAll()
}
