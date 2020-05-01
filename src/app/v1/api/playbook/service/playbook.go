package service

import (
	"github.com/sofyan48/remora/src/app/v1/utility/ansible"
)

// PlaybookService ...
type PlaybookService struct {
	Ansible ansible.AnsbileLibraryInterface
}

// PlaybookServiceHandler ...
func PlaybookServiceHandler() *PlaybookService {
	return &PlaybookService{
		Ansible: ansible.AnsbileLibraryHandler(),
	}
}

// PlaybookServiceInterface ...
type PlaybookServiceInterface interface {
	PlaybookService(conn, inventory string) error
}

// PlaybookService ...
func (service *PlaybookService) PlaybookService(conn, inventory string) error {
	connection := service.Ansible.SetConn(conn)
	invent := service.Ansible.SetInventory(inventory)

	err := service.Ansible.Execute(connection, invent, "playbook/example/example.yml", "Runn Example")

	return err
}
