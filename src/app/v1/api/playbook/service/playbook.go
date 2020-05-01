package service

import (
	"errors"
	"os"

	"github.com/sofyan48/remora/src/app/v1/api/playbook/entity"
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
	PlaybookService(params *entity.PlaybookParameters) (string, error)
}

// PlaybookService ...
func (service *PlaybookService) PlaybookService(params *entity.PlaybookParameters) (string, error) {
	connection := service.Ansible.SetConn(params.Connection)
	invent := service.Ansible.SetInventory(params.Inventory)
	playbookStorage := os.Getenv("PLAYBOOK_PATH")
	playbookPath := playbookStorage + "/" + params.Apps + "/main.yml"
	if playbookStorage == "" {
		return "nil", errors.New("Setup playbook Path Storage")
	}
	if params.Report == "1" {
		return "Playbook Processing", service.Ansible.Execute(connection, invent, playbookPath, "")
	}
	go service.Ansible.Execute(connection, invent, playbookPath, "")
	return "Playbook Processing", nil
}
