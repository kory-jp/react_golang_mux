package seed

import (
	"fmt"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/infrastructure"
)

func TagsData() (tags domain.Tags) {
	tag1 := domain.Tag{
		Value: "shopping",
		Label: "買い物",
	}

	tag2 := domain.Tag{
		Value: "business",
		Label: "仕事",
	}

	tag3 := domain.Tag{
		Value: "hobby",
		Label: "趣味",
	}

	tag4 := domain.Tag{
		Value: "study",
		Label: "学習",
	}

	tag5 := domain.Tag{
		Value: "payment",
		Label: "支払い",
	}

	tags = append(tags, tag1, tag2, tag3, tag4, tag5)
	return
}

func TagsSeed(con *infrastructure.SqlHandler) (err error) {
	tags := TagsData()
	for _, t := range tags {
		cmd := fmt.Sprintf(`
			insert into
				tags(
					value,
					label,
					created_at
				)
			values ("%s", "%s", "%s")
		`, t.Value, t.Label, time.Now().Format("2006/01/02 15:04:05"))
		_, err = con.Conn.Exec(cmd)
	}
	return
}
