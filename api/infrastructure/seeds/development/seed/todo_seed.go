package seed

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kory-jp/react_golang_mux/api/infrastructure"

	"github.com/kory-jp/react_golang_mux/api/domain"
)

func TodosDate() (todos domain.Todos) {
	todo1 := domain.Todo{
		UserID:    1,
		Title:     "買い物",
		Content:   "野菜、果物、魚",
		ImagePath: "",
	}

	todo2 := domain.Todo{
		UserID:    1,
		Title:     "メール送信",
		Content:   "田中さんに明日までに契約書を送信",
		ImagePath: "",
	}

	todo3 := domain.Todo{
		UserID:    1,
		Title:     "資料作成",
		Content:   "テスト株式会社へプレゼンする資料を火曜日までに作成",
		ImagePath: "",
	}

	todo4 := domain.Todo{
		UserID:    1,
		Title:     "面談",
		Content:   "佐藤部長との面談が明日の9時から",
		ImagePath: "",
	}

	todo5 := domain.Todo{
		UserID:    1,
		Title:     "Amazonで米を購入",
		Content:   "合わせて醤油も購入",
		ImagePath: "",
	}

	todo6 := domain.Todo{
		UserID:    1,
		Title:     "銀行口座を開設",
		Content:   "給与振込用の口座をテスト銀行で開設",
		ImagePath: "",
	}

	todo7 := domain.Todo{
		UserID:    1,
		Title:     "旅行計画",
		Content:   "吉村と来月の旅行を調整",
		ImagePath: "",
	}

	todo8 := domain.Todo{
		UserID:    1,
		Title:     "結婚式の招待状を返信",
		Content:   "5月28日までに出席の旨の返信",
		ImagePath: "",
	}

	todo9 := domain.Todo{
		UserID:    1,
		Title:     "住民票をを入手する",
		Content:   "",
		ImagePath: "",
	}

	todo10 := domain.Todo{
		UserID:    1,
		Title:     "引っ越し準備",
		Content:   "荷物の梱包",
		ImagePath: "",
	}

	todo11 := domain.Todo{
		UserID:    1,
		Title:     "家賃の支払い",
		Content:   "",
		ImagePath: "",
	}

	todos = append(todos, todo1, todo2, todo3, todo4, todo5, todo6, todo7, todo8, todo9, todo10, todo11)
	return
}

func TodosSeed(con *infrastructure.SqlHandler) (err error) {
	todos := TodosDate()
	for _, t := range todos {
		cmd := fmt.Sprintf(`
			insert into
			todos(
				user_id,
				title,
				content,
				image_path,
				isFinished,
				created_at
			)
		values (%s, "%s", "%s", "%s", %s, "%s")
		 `, strconv.Itoa(t.UserID), t.Title, t.Content, t.ImagePath, "0", time.Now().Format("2006/01/02 15:04:05"))
		_, err = con.Conn.Exec(cmd)
	}
	return
}
