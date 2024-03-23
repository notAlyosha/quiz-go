package entity

type Checkable interface {
	Check() bool
}

type Saveable interface {
	Save()
}

type Question interface {
	Saveable
	Checkable
}
