package entity

type User struct {
	ID        int    `json: "ID,omitempty"`
	Name_task string `json: "name,omitempty"`
	Ager      int    `json: "ager,omitempty"`
	Group     string `json: "group,omitempty"`
	Priority  int    `json: "priority,omitempty"`
}
