package kakuyasu

import (
	"github.com/sclevine/agouti"
	"github.com/prometheus/common/log"
)

var page *agouti.Page

type Crawler struct {
	ID           string
	Pass         string
	DeliveryTime string
}

func initDriver() agouti.WebDriver {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("failed to start driver:%v", err)
	}
	p, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("failed to open page:%v", err)
	}
	page = p
	return *driver
}

func (s Crawler) Execute() {
	d := initDriver()
	defer d.Stop()

	start()
	login(s.ID, s.Pass)
	selectPresetOrder()
	addToCart()
	confirmItems()
	relogin(s.ID, s.Pass)
	confirmAddress()
	confirmDate()
	confirmDateTime(s.DeliveryTime)
	//order()
}

func handlePage() {

}

func start() {
	if err := page.Navigate("https://www.kakuyasu.co.jp/ec/common/CSfLogin.jsp?transfer=../member/CPmPersonalShop_001.jsp"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
}

func login(id, pass string) {
	p := LoginPage{*page}
	p.IDInput().Fill(id)
	p.pwInput().Fill(pass)
	p.loginButton().Click()
}

func selectPresetOrder() {
	p := MayPage{*page}
	p.presetOrderButton().Click()
}

func addToCart() {
	p := PresetOrderPage{*page}
	p.selectPresetButton().Click()
}

func confirmItems() {
	p := CartPage{*page}
	p.conformItemsButton().Click()
}

func relogin(id, pass string) {
	login(id, pass)
}

func confirmAddress() {
	p := AddressPage{*page}
	p.confirmAddressButton().Click()
}

func confirmDate() {
	p := DatePage{*page}
	p.confirmDateButton().Click()
}

func confirmDateTime(deliveryTime string) {
	p := DateTimePage{*page}
	p.dateTimeOption().Select(deliveryTime)
	p.confirmDateTimeButton().Click()
}

func order() {
	p := OrderPage{*page}
	p.orderButton().Click()
}
