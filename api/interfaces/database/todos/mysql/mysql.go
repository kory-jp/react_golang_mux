package mysql

// --- テストで利用するためクエリ部分を書き出して定義 ---
// --- todo新規作成のクエリ ---
var CreateTodoState = `
	insert into
		todos(
			user_id,
			title,
			content,
			image_path,
			isFinished,
			importance,
			urgency
		)
	value (?, ?, ?, ?, ?, ?, ?)
`

// --- 作成されたtodoの総数 ---
var SumTodoItemsState = `
	select count(*) from
		todos
	where
		user_id = ?
`

// --- Todo一覧取得 ---
var GetTodosState = `
	select
		t.id,
		t.user_id,
		t.title,
		t.image_path,
		t.isFinished,
		t.created_at,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.user_id = ?
	group by
		t.id
	order by
		id desc
	limit 6
	offset ?
`

// --- Todo詳細取得 ---
var ShowTodoState = `
	select
		t.*,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.id = ?
	and
		t.user_id = ?
	group by
		t.id
`

// --- Tag検索 ---
var FindByAllConditionSumTodoItemsState = `
	select count(*) from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		tg.id = ?
	and
		t.user_id = ?	
	and
		t.importance = ?
	and
		t.urgency = ?
`

var FindByTagIdImpScoreSumTodoItemsState = `
	select count(*) from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		tg.id = ?
	and
		t.user_id = ?
	and
		t.importance = ?
`

var FindByTagIdUrgScoreSumTodoItemsState = `
	select count(*) from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		tg.id = ?
	and
		t.user_id = ?
	and
		t.urgency = ?
`

var FindByImpScoreUrgScoreSumTodoItemsState = `
	select count(*) from
		todos as t
	where
		t.user_id = ?
	and
		t.importance = ?
	and
		t.urgency = ?
`

var FindByTagIdSumTodoItemsState = `
	select count(*) from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		tg.id = ?
	and
		t.user_id = ?
`

var FindByImpScoreSumTodoItemsState = `
	select count(*) from
		todos as t
	where
		t.user_id = ?
	and
		t.importance = ?
`

var FindByUrgScoreSumTodoItemsState = `
	select count(*) from
		todos as t
	where
		t.user_id = ?
	and
		t.urgency = ?
`

// -----

var FindByAllConditionTodosState = `
	select
		t.id,
		t.user_id,
		t.title,
		t.image_path,
		t.isFinished,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.id in (
			select
				ttr.todo_id
			from
				todo_tag_relations as ttr
			left join
				tags as tg
			on
				ttr.tag_id = tg.id
			where
				tg.id = ?
		)
	and
		t.user_id = ?
	and
		t.importance = ?
	and
		t.urgency = ?
	group by
		t.id
	order by
		id desc
	limit 5
	offset ?
`

var FindByTagIdImpScoreTodosState = `
	select
		t.id,
		t.user_id,
		t.title,
		t.image_path,
		t.isFinished,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.id in (
			select
				ttr.todo_id
			from
				todo_tag_relations as ttr
			left join
				tags as tg
			on
				ttr.tag_id = tg.id
			where
				tg.id = ?
		)
	and
		t.user_id = ?
	and
		t.importance = ?
	group by
		t.id
	order by
		id desc
	limit 5
	offset ?
`

var FindByTagIdUrgScoreTodosState = `
	select
		t.id,
		t.user_id,
		t.title,
		t.image_path,
		t.isFinished,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.id in (
			select
				ttr.todo_id
			from
				todo_tag_relations as ttr
			left join
				tags as tg
			on
				ttr.tag_id = tg.id
			where
				tg.id = ?
		)
	and
		t.user_id = ?
	and
		t.urgency = ?
	group by
		t.id
	order by
		id desc
	limit 5
	offset ?
`

var FindByImpScoreUrgScoreTodosState = `
	select
		t.id,
		t.user_id,
		t.title,
		t.image_path,
		t.isFinished,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.user_id = ?
	and
		t.importance = ?
	and
		t.urgency = ?
	group by
		t.id
	order by
		id desc
	limit 5
	offset ?
`

var FindByTagIdTodosState = `
	select
		t.id,
		t.user_id,
		t.title,
		t.image_path,
		t.isFinished,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.id in (
			select
				ttr.todo_id
			from
				todo_tag_relations as ttr
			left join
				tags as tg
			on
				ttr.tag_id = tg.id
			where
				tg.id = ?
		)
	and
		t.user_id = ?
	group by
		t.id
	order by
		id desc
	limit 5
	offset ?
`

var FindByImpScoreTodosState = `
	select
		t.id,
		t.user_id,
		t.title,
		t.image_path,
		t.isFinished,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.user_id = ?
	and
		t.importance = ?
	group by
		t.id
	order by
		id desc
	limit 5
	offset ?
`

var FindByUrgScoreTodosState = `
	select
		t.id,
		t.user_id,
		t.title,
		t.image_path,
		t.isFinished,
		group_concat(tg.id),
		group_concat(tg.value),
		group_concat(tg.label)
	from
		todos as t
	left join
		todo_tag_relations as ttr
	on
		t.id = ttr.todo_id
	left join
		tags as tg
	on
		ttr.tag_id = tg.id
	where
		t.user_id = ?
	and
		t.urgency = ?
	group by
		t.id
	order by
		id desc
	limit 5
	offset ?
`

// --- Todo更新 ---
var UpdateTodoState = `
	update
		todos
	set
		title = ?,
		content = ?,
		image_path = ?,
		importance = ?,
		urgency = ?
	where
		id = ?
	and
		user_id = ?
`

// --- isFinishedの真偽値を変更 ---
var ChangeBoolState = `
	update
		todos
	set
		isFinished = ?
	where
		id = ?
	and
		user_id = ?
`

// --- Todo削除 ---
var DeleteTodoState = `
	delete from
		todos
	where
		id = ?
	and
		user_id = ?
`
