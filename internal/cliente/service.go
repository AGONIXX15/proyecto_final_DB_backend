package cliente

import "errors"


type ClienteService struct {
	Repo ClienteRepository
}


func NewClienteService(repo ClienteRepository) *ClienteService {
	return &ClienteService{Repo: repo}
}


func (s *ClienteService) CreateCliente(c *Cliente) error {
		if c.NombreCompleto == "" {
		return errors.New("error nombre viene vacio")
	}

	return s.Repo.Create(c)
}

func (s *ClienteService) UpdateCliente(c *Cliente) error {
	return s.Repo.Update(c)
}



func (s *ClienteService) DeleteCliente(documento uint) error {
	return s.Repo.Delete(documento)
}

func (s *ClienteService) GetAllClientes()  ([]Cliente, error) {
	return s.Repo.FindAll()
}

func (s *ClienteService) GetCliente(documento uint) (*Cliente,error) {
	return s.Repo.FindByID(documento)
}
