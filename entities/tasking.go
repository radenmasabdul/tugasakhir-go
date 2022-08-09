package entities

type Tasking struct {
	Id         int64
	Task       string `validate:"required"`
	Assigne    string `validate:"required"`
	DeadLine   string `validate:"required" label:"Deadline"`
	IsDone     bool
	UpdateDate string
}
