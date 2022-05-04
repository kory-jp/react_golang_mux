package seed

import (
	"fmt"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/infrastructure"
)

func TagsData() (tags domain.Tags) {
	tag1 := domain.Tag{
		Value: "meeting",
		Label: "会議",
	}

	tag2 := domain.Tag{
		Value: "document",
		Label: "資料",
	}

	tag3 := domain.Tag{
		Value: "negotiation",
		Label: "商談",
	}

	tag4 := domain.Tag{
		Value: "report",
		Label: "報告",
	}

	tag5 := domain.Tag{
		Value: "input",
		Label: "入力",
	}

	tag6 := domain.Tag{
		Value: "research",
		Label: "調査",
	}

	tag7 := domain.Tag{
		Value: "planning",
		Label: "計画",
	}

	tag8 := domain.Tag{
		Value: "sales",
		Label: "営業",
	}

	tag9 := domain.Tag{
		Value: "contact",
		Label: "連絡",
	}

	tags = append(tags, tag1, tag2, tag3, tag4, tag5, tag6, tag7, tag8, tag9)
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
