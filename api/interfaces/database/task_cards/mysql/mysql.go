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
			isFinished
		)
	value (?, ?, ?, ?, ?, ?, ?)
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

var ShowTaskCardState = `
	select
		*
	from
		task_cards
	where
		id = ?
	and
		user_id = ?
`

var UpdateTaskCardState = `
	update
		task_cards
	set
		title = ?,
		purpose = ?,
		content = ?,
		memo = ?
	where
		id = ?
	and
		user_id = ?
`

var ChangeBoolState = `
		update
			task_cards
		set
			isFinished = ?
		where
			id = ?
		and
			user_id = ?
`

var DeleteTaskCardState = `
		delete from
			task_cards
		where
			id = ?
		and
			user_id = ?
`
