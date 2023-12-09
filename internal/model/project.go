package model

type Project struct {
	IDUser               string `json:"id_user"`
	ID                   string `json:"id_project"`
	Name                 string `json:"name"`
	DescriptionOfProject string `json:"description_of_project"`
}

type ProjectUser struct {
	IDUser string `json:"id_user"`
	ID     string `json:"id_project"`
}

type Note struct {
	IDProject         string `json:"id_project"`
	ID                string `json:"id_note"`
	DescriptionOfNote string `json:"description_of_note"`
	Completeness      bool   `json:"completeness"`
}
