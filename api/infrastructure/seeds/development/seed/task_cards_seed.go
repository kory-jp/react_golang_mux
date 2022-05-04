package seed

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kory-jp/react_golang_mux/api/domain"
	"github.com/kory-jp/react_golang_mux/api/infrastructure"
)

func TaskCardsDate() (taskCards domain.TaskCards) {
	tc1 := domain.TaskCard{
		UserID:  1,
		TodoID:  1,
		Title:   "過去の春季のトレンドを調査",
		Purpose: "過去のトレンドを調査することで流行の傾向を把握",
		Content: "各調査会社のデータを収集",
		Memo:    "多角的に捉えるために参照する調査会社は5社以上を目標",
	}

	tc2 := domain.TaskCard{
		UserID:  1,
		TodoID:  2,
		Title:   "契約を経理部に送信",
		Purpose: "請求書作成の資料のため、受注に関する詳細を提供する必要がある",
		Content: "今週末までに経理部に契約書、原材料費等の見積書をメールにて送信",
		Memo:    "",
	}

	tc3 := domain.TaskCard{
		UserID:  1,
		TodoID:  3,
		Title:   "メールの仕分け",
		Purpose: "優先すべき事項や、抜けがゆるされない事項を判断するため",
		Content: "全てのメールを開封して、種類ごとにフォルダ分類をする",
		Memo:    "数分で処理が可能なメールはその場で返信する",
	}

	tc4 := domain.TaskCard{
		UserID:  1,
		TodoID:  4,
		Title:   "エントリーシートを確認",
		Purpose: "面接時の質問を考える資料としてESを確認",
		Content: "経歴に関する質問を5つ、志望動機に絡めた質問を5つ用意",
		Memo:    "ES以外にもポートフォリオがあれば、そちらも確認",
	}

	tc5 := domain.TaskCard{
		UserID:  1,
		TodoID:  5,
		Title:   "実績整理",
		Purpose: "上司が実績を確認し易いように入力するデータを選定する",
		Content: "契約書、草案などから数値を抽出、計算する",
		Memo:    "",
	}

	taskCards = append(
		taskCards,
		tc1,
		tc2,
		tc3,
		tc4,
		tc5,
	)
	return
}

func TaskCardsSeed(con *infrastructure.SqlHandler) (err error) {
	taskCards := TaskCardsDate()
	for _, t := range taskCards {
		cmd := fmt.Sprintf(`
			insert into
			todos(
				user_id,
				todo_id,
				title,
				purpose,
				content,
				memo,
				isFinished,
				created_at
			)
		values (%s, "%s", "%s", "%s", %s, "%s", "%s", "%s")
		 `, strconv.Itoa(t.UserID), strconv.Itoa(t.TodoID), t.Title, t.Purpose, t.Content, t.Memo, "0", time.Now().Format("2006/01/02 15:04:05"))
		_, err = con.Conn.Exec(cmd)
	}
	return
}
