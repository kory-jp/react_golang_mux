package database

type TodTagRelationsRepository struct {
	SqlHandler
}

var CreateTodoTagRelationsState = `
	insert into
		todo_tag_relations(
			todo_id,
			tag_id
		)
	value (?, ?)
`

func (repo *TodTagRelationsRepository) Store(todoId int64, tagIds []int) (err error) {
	for _, v := range tagIds {
		_, err = repo.Execute(CreateTodoTagRelationsState, todoId, v)
		if err != nil {
			// fmt.Println(err)
			// log.Println(err)
			return err
		}
	}
	return err
}
