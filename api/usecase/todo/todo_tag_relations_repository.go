package usecase

type TodoTagRelationsRepository interface {
	Store(int64, []int) error
}
