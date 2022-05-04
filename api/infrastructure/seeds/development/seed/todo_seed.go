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
		UserID:     1,
		Title:      "商品の企画書を作成",
		Content:    "来春の新商品の企画書を作成 締切日:来週末 二日前に一度、確認をお願いする",
		ImagePath:  "",
		Importance: 1,
		Urgency:    2,
	}

	todo2 := domain.Todo{
		UserID:     1,
		Title:      "請求書作成依頼",
		Content:    "受注が成功した鈴木株式会社への請求書を経理部の加藤さんに作成依頼 期限:6/15",
		ImagePath:  "",
		Importance: 2,
		Urgency:    2,
	}

	todo3 := domain.Todo{
		UserID:     1,
		Title:      "メール返信",
		Content:    "未返信のメールを返信",
		ImagePath:  "",
		Importance: 2,
		Urgency:    1,
	}

	todo4 := domain.Todo{
		UserID:     1,
		Title:      "就活生の面接",
		Content:    "2時面接の対応 再来週の水曜日 就活生: 鈴木太郎様",
		ImagePath:  "",
		Importance: 2,
		Urgency:    2,
	}

	todo5 := domain.Todo{
		UserID:     1,
		Title:      "営業実績の入力",
		Content:    "社内ツールに営業活動実績を入力  進捗具合 成約率 不安点 対応策",
		ImagePath:  "",
		Importance: 1,
		Urgency:    2,
	}

	todo6 := domain.Todo{
		UserID:     1,
		Title:      "翌四半期の営業計画",
		Content:    "目標のコンタクト数 制約数 成約率 活動地域の選定 来月のミーティングで発表",
		ImagePath:  "",
		Importance: 2,
		Urgency:    2,
	}

	todo7 := domain.Todo{
		UserID:     1,
		Title:      "東京フーズ株式会社との合同イベント企画",
		Content:    "会議場所:新宿オフィス 日時:5/30 先方:緒方部長 同伴者:吉田係長",
		ImagePath:  "",
		Importance: 1,
		Urgency:    1,
	}

	todo8 := domain.Todo{
		UserID:     1,
		Title:      "業務改善に向けた調査",
		Content:    "第二営業課の従業員から受けたフィードバックから、業務の効率化を図る",
		ImagePath:  "",
		Importance: 2,
		Urgency:    2,
	}

	todo9 := domain.Todo{
		UserID:     1,
		Title:      "神奈川商事へのプレゼンのため、地域トレンドの調査",
		Content:    "プレゼン日時:7/24 地域に根付いた活動をされている神奈川商事への弊社商品をアピールするために地域のトレンドを調査",
		ImagePath:  "",
		Importance: 1,
		Urgency:    2,
	}

	todo10 := domain.Todo{
		UserID:     1,
		Title:      "先月の出張の精算処理",
		Content:    "先日の11~14日で訪問した徳島出張の旅費精算を申請",
		ImagePath:  "",
		Importance: 2,
		Urgency:    2,
	}

	todo11 := domain.Todo{
		UserID:     1,
		Title:      "月末会議の準備",
		Content:    "二日後の社内ミーティングの準備 報告内容 今月の活動実績 来月の活動予定 業務効率化の提案  上記の内容を発表するための台本準備",
		ImagePath:  "",
		Importance: 1,
		Urgency:    1,
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
				importance,
				urgency,
				created_at
			)
		values (%s, "%s", "%s", "%s", %s, "%s", "%s", "%s")
		 `, strconv.Itoa(t.UserID), t.Title, t.Content, t.ImagePath, "0", strconv.Itoa(t.Importance), strconv.Itoa(t.Urgency), time.Now().Format("2006/01/02 15:04:05"))
		_, err = con.Conn.Exec(cmd)
	}
	return
}
