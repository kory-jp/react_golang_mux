package seed

import (
	"fmt"
	"strconv"

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

	tc6 := domain.TaskCard{
		UserID:  1,
		TodoID:  6,
		Title:   "今四半期の見直し",
		Purpose: "翌四半期の営業計画を策定するため、基準となる数値を確認",
		Content: "企画書、契約書、草案等の資料集め  \n 数値の集計、計算",
		Memo:    "岡野に資料集めの手伝いを依頼",
	}

	tc7 := domain.TaskCard{
		UserID:  1,
		TodoID:  7,
		Title:   "当日のイベントのスケジュールを計画",
		Purpose: "当日の時間配分を確定させ、プレゼンの内容を取捨選択する",
		Content: "先方の佐々木様とタイムスケジュールを相談",
		Memo:    "相談日時:水曜日の17時",
	}

	tc8 := domain.TaskCard{
		UserID:  1,
		TodoID:  8,
		Title:   "木曜日までに第二営業課からフィードバックシートを回収",
		Purpose: "週明けの月曜日の部署会議において、今後の活動方針を共有するため、事前に問題点の洗い出しが必要なので木曜日を締め切りで回収",
		Content: "部署全体にフィードバックシートの提出締め切りをメールにて周知する",
		Memo:    "",
	}

	tc9 := domain.TaskCard{
		UserID:  1,
		TodoID:  9,
		Title:   "SNS,Googleトレンドなどから神奈川県に関連した検索ワードを抽出",
		Purpose: "検索ワードから地域の傾向を把握する",
		Content: "Twitter,Facebook,Instagram,Googleトレンドの検索量からキーワードを抽出",
		Memo:    "検索期間は過去四年間に設定 \n 検索ワードは飲食、娯楽、芸術に絞る",
	}

	tc10 := domain.TaskCard{
		UserID:  1,
		TodoID:  10,
		Title:   "領収書を整理",
		Purpose: "経理部へ精算申請する際に領収書も添付する必要がある",
		Content: "領収書を整理して、勘定科目ごとに仕分けまでして提出",
		Memo:    "提出時に先々月の提出忘れの領収書も申請可能か確認",
	}

	tc11 := domain.TaskCard{
		UserID:  1,
		TodoID:  11,
		Title:   "レジュメを作成",
		Purpose: "効果的な会議運営、意見の活性化のためレジュメを作成して事前に配布する",
		Content: "レジュメに以下を記載 \n 会議のスケジュール \n 活動実績の分析 \n 来月以降の目標",
		Memo:    "レジュメは藤田部長へ一度確認をお願いする",
	}

	taskCards = append(
		taskCards,
		tc1,
		tc2,
		tc3,
		tc4,
		tc5,
		tc6,
		tc7,
		tc8,
		tc9,
		tc10,
		tc11,
	)
	return
}

func TaskCardsSeed(con *infrastructure.SqlHandler) (err error) {
	taskCards := TaskCardsDate()
	for _, t := range taskCards {
		cmd := fmt.Sprintf(`
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
		values (%s, %s, "%s", "%s", "%s", "%s", %s)
		 `, strconv.Itoa(t.UserID), strconv.Itoa(t.TodoID), t.Title, t.Purpose, t.Content, t.Memo, "0")
		_, err = con.Conn.Exec(cmd)
	}
	return
}
