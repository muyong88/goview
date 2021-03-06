package functional

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/sayems/golang.webdriver/selenium/pages"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tebeka/selenium"
)

var browser selenium.WebDriver
var page pages.Page

func Test_can_start_a_table_and_see_it_later(t *testing.T) {

	var err error

	// set browser as chrome
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})

	// remote to selenium server
	if browser, err = selenium.NewRemote(caps, ""); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}

	Convey("Test should start a table and see it later", t, func() {

		Convey(` Edith want to look the telemetry data from spacecraft,
			she has heard about a cool new online goview app.
			She goes to check out its homepage `, func() {
			err = browser.Get("http://guest:guest@127.0.0.1:8080/")
			So(err == nil, ShouldBeTrue)
		})

		// use page to interact with browser
		page = pages.Page{Driver: browser}

		Convey("She noticed the title mention 'goview'", func() {
			title, _ := browser.Title()
			So(strings.Contains(title, "goview"), ShouldBeTrue)
		})

		Convey("She is invited to select a data table from a tree to view", func() {
			tree := page.FindElementByLinkText("WYG")
			tree.Click()
			item, err := tree.FindElement("xpath", "//span[contains(text(),'PK-CEH2.xml')]")
			item.Click()
			So(err == nil, ShouldBeTrue)
		})

		Convey("She selects the PK-CEH2.xml table", func() {
			table := page.FindElementByXpath("//div[@class='layui-tab' and @lay-filter='param-tab']")
			So(table, ShouldNotBeNil)
			name, _ := table.Text()
			So(name, ShouldContainSubstring, "PK-CEH2.xml")
		})

		Convey("When she hits enter, the page updates, and now the page shows a table named gcyctd", func() {})

		Convey("Just this time, gcyctd data is sent to the table, she sees the data varing in 2 fps", func() {
			p := simu_init_kafka()
			defer p.Close()

			table := page.FindElementByXpath("//div[@class='layui-tab' and @lay-filter='param-tab']")
			frame, _ := table.FindElement("xpath", "//iframe[contains(@src, 'PK-CEH2.xml')]")
			So(frame, ShouldNotBeNil)

			// total 10 seconds
			for i := 0; i < 100; i++ {
				simu_send_kafka(p, i)
				time.Sleep(time.Millisecond * 100)
			}
		})

		Convey(`Edith wonders whether the site will remember her table.Then she sees that the site has generated
			a unique URL for her -- there is some explanatory text to that effect. `, func() {})

		Convey("She visits that URL - her table is still there", func() {})

		Convey("Satisfied, she goes back to sleep", func() {})
	})

	browser.Quit()
}
