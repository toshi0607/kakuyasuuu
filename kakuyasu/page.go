package kakuyasu

import (
	"github.com/sclevine/agouti"
)

// ログイン
type LoginPage struct {
	agouti.Page
}

func (l LoginPage) loginButton() *agouti.Selection {
	return l.Find("#bodyContainer > form > div.login_left > div > div.loginbtn > input")
}

func (l LoginPage) IDInput() *agouti.Selection {
	return l.Find("#bodyContainer > form > div.login_left > div > div.heightLine-group99 > div > dl:nth-child(1) > dd > ul > li > input")
}

func (l LoginPage) pwInput() *agouti.Selection {
	return l.Find("#bodyContainer > form > div.login_left > div > div.heightLine-group99 > div > dl:nth-child(2) > dd > ul > li > input")
}

// マイページ
type MayPage struct {
	agouti.Page
}

func (m MayPage) presetOrderButton() *agouti.Selection {
	return m.FindByID("mylink5")
}

// 定番注文
type PresetOrderPage struct {
	agouti.Page
}

func (p PresetOrderPage) selectPresetButton() *agouti.Selection {
	return p.Find("#bodyContainer > table > tbody > tr:nth-child(9) > td:nth-child(7) > input:nth-child(1)")
}

// カート
type CartPage struct {
	agouti.Page
}

func (c CartPage) conformItemsButton() *agouti.Selection  {
	return c.FindByID("goOrder")
}

// STEP1.お届け先の指定
type AddressPage struct {
	agouti.Page
}

func (a AddressPage) confirmAddressButton() *agouti.Selection {
	return a.FindByID("setDelDay")
}

// STEP2.お届け方法・希望日の指定
type DatePage struct {
	agouti.Page
}

func (d DatePage) confirmDateButton() *agouti.Selection {
	return d.FindByID("imgGoNext")
}

// STEP3.お届け時間・お支払方法の指定
type DateTimePage struct {
	agouti.Page
}

func (d DateTimePage) dateTimeOption() *agouti.Selection {
	return d.Find("#timeAndPay > div > dl.payInfo_tb2.dlvtime > dd > select")
}

func (d DateTimePage) confirmDateTimeButton() *agouti.Selection {
	return d.FindByID("goOrderConfirmFrom1hour04")
}

// STEP4.ご注文内容の確認
type OrderPage struct {
	agouti.Page
}

func (o OrderPage) orderButton() *agouti.Selection {
	return o.FindByID("goOrder")
}
