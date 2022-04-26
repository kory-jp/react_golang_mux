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

var SumTaskCardItemsState = `
		select count(*) from
			task_cards
		where
			user_id = ?
		and
			todo_id = ?
`

var GetTaskCardsState = `
		select
			*
		from
			task_cards
		where
			user_id = ?
		and
			todo_id = ?
		order by
			id desc
		limit 5
		offset ?
`
