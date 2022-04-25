package mysql

var GetTagsState = `
	select
		id,
		value,
		label
	from
		tags
`
