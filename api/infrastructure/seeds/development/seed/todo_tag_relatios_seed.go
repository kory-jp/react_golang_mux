package seed

import (
	"fmt"
	"strconv"

	"github.com/kory-jp/react_golang_mux/api/domain"

	"github.com/kory-jp/react_golang_mux/api/infrastructure"
)

func TodoTagRelationsDate() (todoTagRelations domain.TodoTagRelations) {
	ttr1 := domain.TodoTagRelation{
		TodoID: 1,
		TagID:  2,
	}

	ttr2 := domain.TodoTagRelation{
		TodoID: 2,
		TagID:  2,
	}

	ttr3 := domain.TodoTagRelation{
		TodoID: 3,
		TagID:  9,
	}

	ttr4 := domain.TodoTagRelation{
		TodoID: 5,
		TagID:  5,
	}

	ttr5 := domain.TodoTagRelation{
		TodoID: 6,
		TagID:  7,
	}

	ttr6 := domain.TodoTagRelation{
		TodoID: 7,
		TagID:  7,
	}

	ttr7 := domain.TodoTagRelation{
		TodoID: 8,
		TagID:  6,
	}

	ttr8 := domain.TodoTagRelation{
		TodoID: 9,
		TagID:  6,
	}

	ttr9 := domain.TodoTagRelation{
		TodoID: 10,
		TagID:  4,
	}

	ttr10 := domain.TodoTagRelation{
		TodoID: 11,
		TagID:  1,
	}

	todoTagRelations = append(todoTagRelations, ttr1, ttr2, ttr3, ttr4, ttr5, ttr6, ttr7, ttr8, ttr9, ttr10)
	return
}

func TodoTagRelationsSeed(con *infrastructure.SqlHandler) (err error) {
	todoTagRelations := TodoTagRelationsDate()
	for _, t := range todoTagRelations {
		cmd := fmt.Sprintf(`
			insert into
			todo_tag_relations(
				todo_id,
				tag_id
			)
		values (%s, "%s")
		 `, strconv.Itoa(t.TodoID), strconv.Itoa(t.TagID))
		_, err = con.Conn.Exec(cmd)
	}
	return
}
