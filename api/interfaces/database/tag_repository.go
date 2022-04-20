package database

import (
	"fmt"
	"log"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

type TagRepository struct {
	SqlHandler
}

var GetTagsState = `
	select
		*
	from
		tags
`

func (repo *TagRepository) FindAll() (tags domain.Tags, err error) {
	rows, err := repo.Query(GetTagsState)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag domain.Tag
		err = rows.Scan(
			&tag.ID,
			&tag.Value,
			&tag.Label,
		)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
			return nil, err
		}
		tags = append(tags, tag)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		log.Panicln(err)
		return nil, err
	}
	rows.Close()
	return tags, nil
}
