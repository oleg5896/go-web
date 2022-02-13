package repository

type WebItem interface {
}

type WebList interface {
}

type Repository struct {
	WebItem
	WebList
}

func newRepository() *Repository {
	return &Repository{}
}
