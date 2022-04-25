package mysql

var CreateTaskCardState = `
	insert into
		task_cards(
			user_id,
			todo_id,
			title,
			purpose,
			content,
			memo,
			isFinished,
			created_at
		)
	value (?, ?, ?, ?, ?, ?, ?, ?)
`
