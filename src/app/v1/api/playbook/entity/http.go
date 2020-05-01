package entity

// PlaybookParameters  ...
type PlaybookParameters struct {
	Inventory  string `json:"inventory" form:"inventory"`
	Connection string `json:"connection" form:"connection"`
	Apps       string `json:"apps" form:"apps"`
	Report     string `json:"report" form:"report"`
}
