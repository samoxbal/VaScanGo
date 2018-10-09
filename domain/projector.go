package domain

type Projector interface {
	Project(readModel interface{}) error
}