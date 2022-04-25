package mysql

var FindByEmailState = `
	select
		*
	from
		users
	where
	 email = ?
`

var FindByIdState = `
	select
		*
	from
		users
	where
	 id = ?
`
