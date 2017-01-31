package element_scraper

import (
  "testing"
  "strings"
)

const WebPage = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:fb="http://ogp.me/ns/fb#">
<head>
  <title>Restaurant</title>
  <meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
</head>

<body>
  <div id="pagewrapper">
    <div id="main" class="clearfix">
      <div id="content_wrapper">
        <div id='content'>
          <div class="column right portlets">
            <div class= 'columnContent'>
              <div class='portlet'>
                <div class='lunch'>
                  <h2>Lounas, Jyväskylä: </h2>
                  <table class='todayLunch'>
                    <tr><td class='section'  >Torstai:</td><td></td></tr>
                    <tr><td class='dish'  >Chicken Vindaloo</td><td class='price'>9.00&euro;</td></tr>
                    <tr class='row2'><td class='desc'>Kanaa ja perunoita tulisessa kastikkeessa</td><td></td></tr>
                    <tr><td class='dish'  >Lamb Korma</td><td class='price'>9.00&euro;</td></tr>
                    <tr class='row2'><td class='desc'>Lammasta kermakastikkeessa</td><td></td></tr>
                    <tr><td class='dish'  >Palak panir</td><td class='price'>9.00&euro;</td></tr>
                    <tr class='row2'><td class='desc'>Pinaattia ja juustoa kermaisessa tomaattikastikkeessa</td><td></td></tr>
                  </table></div></div>
                </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</body>
</html>
`

func TestFindElement(t *testing.T) {
  reader := strings.NewReader(WebPage)
  elementScraper, err := NewElementScraperFromReader(reader)
  element := elementScraper.Find(".todayLunch") // "table"
  if err != nil {
    t.Error("Error:", err)
  } else {
    if !strings.Contains(element, "Chicken Vindaloo") {
      t.Error("Element content is not correct")
    }
    if strings.Contains(element, "Jyväskylä") {
      t.Error("Too much data included")
    }
  }
}
