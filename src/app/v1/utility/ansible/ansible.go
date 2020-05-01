package ansible

import (
	ansibler "github.com/apenella/go-ansible"
)

// AnsbileLibrary ...
type AnsbileLibrary struct{}

// AnsbileLibraryHandler ...
func AnsbileLibraryHandler() *AnsbileLibrary {
	return &AnsbileLibrary{}
}

// AnsbileLibraryInterface ...
type AnsbileLibraryInterface interface {
}

// SetConn ...
func (handler *AnsbileLibrary) SetConn(conn string) *ansibler.AnsiblePlaybookConnectionOptions {
	return &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: conn,
	}
}

// SetInventory ...
func (handler *AnsbileLibrary) SetInventory(inventory string) *ansibler.AnsiblePlaybookOptions {
	return &ansibler.AnsiblePlaybookOptions{
		Inventory: inventory + ",",
	}
}

// Execute ...
func (handler *AnsbileLibrary) Execute(conn *ansibler.AnsiblePlaybookConnectionOptions,
	inventory *ansibler.AnsiblePlaybookOptions, path, prefix string) error {
	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          path,
		ConnectionOptions: conn,
		Options:           inventory,
		ExecPrefix:        prefix,
	}
	err := playbook.Run()
	if err != nil {
		return err
	}
	return nil
}
