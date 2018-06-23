package main

import (
	"github.com/sclevine/agouti"
	"github.com/prometheus/common/log"
	"time"
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

	if err := page.Navigate("https://www.kakuyasu.co.jp/ec/common/CSfLogin.jsp?transfer=../member/CPmPersonalShop_001.jsp"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	log.Info(page.Title())
	page.Screenshot("log/01.png")

	page.Find("#bodyContainer > form > div.login_left > div > div.heightLine-group99 > div > dl:nth-child(1) > dd > ul > li > input").Fill(id)
	page.Find("#bodyContainer > form > div.login_left > div > div.heightLine-group99 > div > dl:nth-child(2) > dd > ul > li > input").Fill(pass)
	page.Find("#bodyContainer > form > div.login_left > div > div.loginbtn > input").Click()
	time.Sleep(3 * time.Second)
	page.Screenshot("log/02.png")
}
