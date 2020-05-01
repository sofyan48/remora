package service

import (
	"os"

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
	PlaybookService(app, conn, inventory string) error
}

// PlaybookService ...
func (service *PlaybookService) PlaybookService(app, conn, inventory string) error {
	connection := service.Ansible.SetConn(conn)
	invent := service.Ansible.SetInventory(inventory)
	playbookStorage := os.Getenv("PLAYBOOK_PATH")
	playbookPath := playbookStorage + "/" + app + "/main.yml"
	err := service.Ansible.Execute(connection, invent, playbookPath, "")
	return err
}
