package main

import (
	"github.com/sclevine/agouti"
	"github.com/prometheus/common/log"
	"os"
)

func main() {
	id := os.Getenv("KAKUYASU_ID")
	pass := os.Getenv("KAKUYASU_PASS")

	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	// ログイン
	if err := page.Navigate("https://www.kakuyasu.co.jp/ec/common/CSfLogin.jsp?transfer=../member/CPmPersonalShop_001.jsp"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	log.Info(page.Title())
	page.Screenshot("log/01.png")

	page.Find("#bodyContainer > form > div.login_left > div > div.heightLine-group99 > div > dl:nth-child(1) > dd > ul > li > input").Fill(id)
	page.Find("#bodyContainer > form > div.login_left > div > div.heightLine-group99 > div > dl:nth-child(2) > dd > ul > li > input").Fill(pass)
	page.Find("#bodyContainer > form > div.login_left > div > div.loginbtn > input").Click()
	page.Screenshot("log/02.png")

	// 定番注文
	page.FindByID("mylink5").Click()
	page.Screenshot("log/03.png")

	// カートに追加
	page.Find("#bodyContainer > table > tbody > tr:nth-child(9) > td:nth-child(7) > input:nth-child(1)").Click()
	page.Screenshot("log/04.png")

	// ご注文手続きへ進む
	page.FindByID("goOrder").Click()
	page.Screenshot("log/05.png")

	// ログイン情報再入力
	page.Find("#bodyContainer > form > div.login_left > div > div.heightLine-group99 > div > dl:nth-child(1) > dd > ul > li > input").Fill(id)
	page.Find("#bodyContainer > form > div.login_left > div > div.heightLine-group99 > div > dl:nth-child(2) > dd > ul > li > input").Fill(pass)
	page.Find("#bodyContainer > form > div.login_left > div > div.loginbtn > input").Click()
	page.Screenshot("log/06.png")

	// お届け方法・希望日を指定する
	price, _ := page.Find("#COdorInfo_tb2 > div.step1Area01 > div:nth-child(6) > dl.end.COdorInfo_tb_order04.txt_bold > dd").Text()
	if price != "2,551円" {
		log.Fatalf("price changed to %s", price)
	}
	// ご注文手続きへ進む
	page.FindByID("setDelDay").Click()
	page.Screenshot("log/07.png")

	// お届け時間・お支払い方法を指定する
	page.FindByID("imgGoNext").Click()
	page.Screenshot("log/08.png")

	// ご注文の確認へ進む
	page.Find("#timeAndPay > div > dl.payInfo_tb2.dlvtime > dd > select").Select("2000")
	page.Screenshot("log/09.png")
	page.FindByID("goOrderConfirmFrom1hour04").Click()
	page.Screenshot("log/10.png")

	// 注文の確定
	price, _ = page.FindByClass("txb").Text()
	if price != "2,551円" {
		log.Fatalf("price changed to %s", price)
	}
	//page.FindByID("goOrder").Click()
	page.Screenshot("log/11.png")
}
