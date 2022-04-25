package mysql

var CreateTodoTagRelationsState = `
	insert into
		todo_tag_relations(
			todo_id,
			tag_id,
			created_at
		)
	value (?, ?, ?)
`

var DeleteTodoTagRelationsState = `
		delete from
			todo_tag_relations
		where
			todo_tag_relations.todo_id = ?
`
