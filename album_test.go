package lastfm

import (
	"encoding/json"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct{}

var _ = Suite(&S{})

func (s *S) TestBuyLinksUnmarshal(c *C) {
	respString := `{"affiliations":{"physicals":{"affiliation":[{"supplierName":"Amazon","price":{"currency":"GBP","amount":"6.31","formatted":"\u00a36.31"},"buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/1\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/1\/amazon.gif","isSearch":"0"},{"supplierName":"eBay","buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/90\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/ebay.gif","isSearch":"1"}]},"downloads":{"affiliation":[{"supplierName":"Amazon MP3","price":{"currency":"GBP","amount":"7.90","formatted":"\u00a37.90"},"buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/44\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/amazon-mp3-16x16-a.gif","isSearch":"0"},{"supplierName":"7digital","price":{"currency":"GBP","amount":"7.99","formatted":"\u00a37.99"},"buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/13\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/7digital.gif","isSearch":"0"},{"supplierName":"iTunes","price":{"currency":"GBP","amount":"7.99","formatted":"\u00a37.99"},"buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/24\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/itunesbadge.gif","isSearch":"0"}]}}}`
	v := BuyLinksResponse{}
	err := json.Unmarshal([]byte(respString), &v)
	c.Assert(err, IsNil)
	c.Assert(len(v.Affiliations.Physicals.Affiliation), Equals, 2)
	c.Assert(v.Affiliations.Physicals.Affiliation[0].SupplierName, Equals, "Amazon")
}

func (s *S) TestBuyLinksShortcuts(c *C) {
	respString := `{"affiliations":{"physicals":{"affiliation":[{"supplierName":"Amazon","price":{"currency":"GBP","amount":"6.31","formatted":"\u00a36.31"},"buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/1\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/1\/amazon.gif","isSearch":"0"},{"supplierName":"eBay","buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/90\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/ebay.gif","isSearch":"1"}]},"downloads":{"affiliation":[{"supplierName":"Amazon MP3","price":{"currency":"GBP","amount":"7.90","formatted":"\u00a37.90"},"buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/44\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/amazon-mp3-16x16-a.gif","isSearch":"0"},{"supplierName":"7digital","price":{"currency":"GBP","amount":"7.99","formatted":"\u00a37.99"},"buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/13\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/7digital.gif","isSearch":"0"},{"supplierName":"iTunes","price":{"currency":"GBP","amount":"7.99","formatted":"\u00a37.99"},"buyLink":"http:\/\/www.last.fm\/affiliate\/byid\/8\/3418994\/24\/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22","supplierIcon":"http:\/\/cdn.last.fm\/favicons\/itunesbadge.gif","isSearch":"0"}]}}}`
	v := BuyLinksResponse{}
	err := json.Unmarshal([]byte(respString), &v)
	c.Check(err, IsNil)

	c.Assert(len(v.Physicals()), Equals, 2)
	c.Assert(v.Physicals()[0].SupplierName, Equals, "Amazon")

	c.Assert(len(v.Downloads()), Equals, 3)
	c.Assert(v.Downloads()[1].SupplierName, Equals, "7digital")
}

func (s *S) TestBuyLinksE2E(c *C) {
	client := NewClient("d0a8c6b594b43669503d9f51aaabea22", "")
	buyLinks, err := client.Album().GetBuyLinks("Radiohead", "OK Computer", false, "GB")
	c.Assert(err, IsNil)
	c.Assert(buyLinks.Physicals(), Not(HasLen), 0)
}
