package lastfm

import (
	"encoding/xml"
	"fmt"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct{}

var _ = Suite(&S{})

func (s *S) TestBuyLinksUnmarshal(c *C) {
	respString := `<?xml version="1.0" encoding="utf-8"?>
<lfm status="ok"><affiliations><physicals><affiliation><supplierName>Amazon</supplierName><price><currency>GBP</currency><amount>6.10</amount><formatted>£6.10</formatted></price><buyLink>http://www.last.fm/affiliate/byid/8/3418994/1/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22</buyLink><supplierIcon>http://cdn.last.fm/favicons/1/amazon.gif</supplierIcon><isSearch>0</isSearch></affiliation><affiliation><supplierName>eBay</supplierName><buyLink>http://www.last.fm/affiliate/byid/8/3418994/90/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22</buyLink><supplierIcon>http://cdn.last.fm/favicons/ebay.gif</supplierIcon><isSearch>1</isSearch></affiliation></physicals><downloads><affiliation><supplierName>Amazon MP3</supplierName><price><currency>GBP</currency><amount>7.90</amount><formatted>£7.90</formatted></price><buyLink>http://www.last.fm/affiliate/byid/8/3418994/44/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22</buyLink><supplierIcon>http://cdn.last.fm/favicons/amazon-mp3-16x16-a.gif</supplierIcon><isSearch>0</isSearch></affiliation><affiliation><supplierName>7digital</supplierName><price><currency>GBP</currency><amount>7.99</amount><formatted>£7.99</formatted></price><buyLink>http://www.last.fm/affiliate/byid/8/3418994/13/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22</buyLink><supplierIcon>http://cdn.last.fm/favicons/7digital.gif</supplierIcon><isSearch>0</isSearch></affiliation><affiliation><supplierName>iTunes</supplierName><price><currency>GBP</currency><amount>7.99</amount><formatted>£7.99</formatted></price><buyLink>http://www.last.fm/affiliate/byid/8/3418994/24/ws.album.buylinks.d0a8c6b594b43669503d9f51aaabea22</buyLink><supplierIcon>http://cdn.last.fm/favicons/itunesbadge.gif</supplierIcon><isSearch>0</isSearch></affiliation></downloads></affiliations></lfm>`
	v := BuyLinks{}
	err := xml.Unmarshal([]byte(respString), &struct {
		Affiliations *BuyLinks `xml:"affiliations"`
	}{Affiliations: &v})
	c.Assert(err, IsNil)
	c.Assert(len(v.Physicals), Equals, 2)
	c.Assert(v.Physicals[0].SupplierName, Equals, "Amazon")
}

func (s *S) TestBuyLinksE2E(c *C) {
	client := NewClient("d0a8c6b594b43669503d9f51aaabea22", "")
	buyLinks, err := client.Album().GetBuyLinks("Radiohead", "OK Computer", false, "GB")
	c.Assert(err, IsNil)
	c.Assert(buyLinks.Physicals, Not(HasLen), 0)
}

var _ = fmt.Println

func (s *S) TestGetInfoUnmarshal(c *C) {
	respString := `<?xml version="1.0" encoding="utf-8"?>
<lfm status="ok"><album><name>Believe</name><artist>Cher</artist><id>2026126</id><mbid>250e1aa0-fbb9-4f15-8321-3550b6c742ac</mbid><url>http://www.last.fm/music/Cher/Believe</url><releasedate>    5 Jul 2005, 00:00</releasedate><image size="small">http://userserve-ak.last.fm/serve/34s/72903330.png</image><image size="medium">http://userserve-ak.last.fm/serve/64s/72903330.png</image><image size="large">http://userserve-ak.last.fm/serve/174s/72903330.png</image><image size="extralarge">http://userserve-ak.last.fm/serve/300x300/72903330.png</image><image size="mega">http://userserve-ak.last.fm/serve/_/72903330/Believe.png</image><listeners>205287</listeners><playcount>1067073</playcount><tracks><track rank="1"><name>Believe</name><duration>222</duration><mbid>028523f5-23b3-4910-adc1-46d932e2fb55</mbid><url>http://www.last.fm/music/Cher/_/Believe</url><streamable fulltrack="0">1</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track><track rank="2"><name>The Power</name><duration>236</duration><mbid>173cf503-cc44-4291-ab91-2286aafe6efa</mbid><url>http://www.last.fm/music/Cher/_/The+Power</url><streamable fulltrack="0">0</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track><track rank="3"><name>Runaway</name><duration>286</duration><mbid>379f760d-1f29-4317-ab04-06a8218a874d</mbid><url>http://www.last.fm/music/Cher/_/Runaway</url><streamable fulltrack="0">0</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track><track rank="4"><name>All or Nothing</name><duration>237</duration><mbid>12e9720f-c14f-45f1-8c47-de8a3a08f4c0</mbid><url>http://www.last.fm/music/Cher/_/All+or+Nothing</url><streamable fulltrack="0">0</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track><track rank="5"><name>Strong Enough</name><duration>223</duration><mbid>03bc1da8-3b25-4060-b843-1fbf60c046a8</mbid><url>http://www.last.fm/music/Cher/_/Strong+Enough</url><streamable fulltrack="0">1</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track><track rank="6"><name>Dov'è L'amore</name><duration>258</duration><mbid>08d88e27-6245-4684-aada-c78a8eb996be</mbid><url>http://www.last.fm/music/Cher/_/Dov%27%C3%A8+L%27amore</url><streamable fulltrack="0">0</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track><track rank="7"><name>Takin' Back My Heart</name><duration>272</duration><mbid>07a38e80-ba81-494a-a61a-e8d81a40413e</mbid><url>http://www.last.fm/music/Cher/_/Takin%27+Back+My+Heart</url><streamable fulltrack="0">0</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track><track rank="8"><name>Taxi Taxi</name><duration>304</duration><mbid>66f526c9-b135-4458-86cf-77065ce8f0aa</mbid><url>http://www.last.fm/music/Cher/_/Taxi+Taxi</url><streamable fulltrack="0">0</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track><track rank="9"><name>Love Is the Groove</name><duration>271</duration><mbid>832f8f9a-95e4-476b-b108-14dec1dc84ba</mbid><url>http://www.last.fm/music/Cher/_/Love+Is+the+Groove</url><streamable fulltrack="0">0</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track><track rank="10"><name>We All Sleep Alone</name><duration>233</duration><mbid>0188050c-401a-4633-a593-ba137390e9c4</mbid><url>http://www.last.fm/music/Cher/_/We+All+Sleep+Alone</url><streamable fulltrack="0">1</streamable><artist><name>Cher</name><mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid><url>http://www.last.fm/music/Cher</url></artist></track></tracks><toptags><tag><name>sourabh</name><url>http://www.last.fm/tag/sourabh</url></tag><tag><name>albums</name><url>http://www.last.fm/tag/albums</url></tag><tag><name>pop</name><url>http://www.last.fm/tag/pop</url></tag><tag><name>90s</name><url>http://www.last.fm/tag/90s</url></tag><tag><name>dance</name><url>http://www.last.fm/tag/dance</url></tag></toptags><wiki><published>Sat, 6 Mar 2010 16:48:03 +0000</published><summary><![CDATA[Believe is the twenty-third studio album by American  singer-actress Cher, released on November 10, 1998 by Warner Bros. Records. The RIAA certified it Quadruple Platinum on December 23, 1999, recognizing four million shipments in the United States; Worldwide, the album has sold more than 20 million copies, making it the biggest-selling album of her career. In 1999 the album received three Grammy Awards nominations including &quot;Record of the Year&quot;, &quot;Best Pop Album&quot; and winning &quot;Best Dance Recording&quot; for the single &quot;Believe&quot;. ]]></summary><content><![CDATA[REPLACED CONTENT]]></content></wiki></album></lfm>
`
	v := AlbumInfo{}
	err := xml.Unmarshal([]byte(respString), &struct {
		Album *AlbumInfo `xml:"album"`
	}{Album: &v})
	c.Assert(err, IsNil)
	c.Check(v.Name, Equals, "Believe")
	c.Check(v.Artist, Equals, "Cher")
	c.Check(v.Images, HasLen, 5)
	c.Check(v.Images[3].Url, Equals, "http://userserve-ak.last.fm/serve/300x300/72903330.png")
	c.Check(v.Tracks, HasLen, 10)
	c.Check(v.Tracks[2].Name, Equals, "Runaway")
}

func (s *S) TestGetInfoE2E(c *C) {
	client := NewClient("d0a8c6b594b43669503d9f51aaabea22", "")
	albumInfo, err := client.Album().GetInfo("Radiohead", "OK Computer", false, "", "")
	c.Assert(err, IsNil)
	c.Assert(albumInfo.Name, Equals, "OK Computer")
	c.Assert(albumInfo.Tracks, HasLen, 12)
}

func (s *S) TestGetShoutsUnmarshal(c *C) {
	respString := `<?xml version="1.0" encoding="utf-8"?>
<lfm status="ok"><shouts artist="Cher" album="Believe" page="1" perPage="50" totalPages="1" total="36"><shout><body>♥ ♥ ♥</body><author>Sweetboy_88</author><date>Sun, 28 Oct 2012 19:30:33</date></shout><shout><body>COMEBACK DO SÉCULO! &lt;3</body><author>igorgeneration</author><date>Wed, 10 Oct 2012 00:45:03</date></shout><shout><body>Taxi Taxi &lt;3</body><author>digiroro</author><date>Sun, 16 Sep 2012 17:44:54</date></shout><shout><body>Your the best cher the one and only and u always will be i love your music !!!!!</body><author>chenoarose60</author><date>Tue, 26 Jun 2012 23:36:06</date></shout><shout><body>&lt;3</body><author>FreeMHZ</author><date>Fri, 23 Mar 2012 02:14:48</date></shout><shout><body>&lt;3</body><author>MisterBoombox</author><date>Sun, 26 Feb 2012 16:36:55</date></shout><shout><body>м м)</body><author>EricValeev</author><date>Mon, 26 Sep 2011 10:56:03</date></shout><shout><body>Still sounds fresh</body><author>lacerdathiago</author><date>Sat, 3 Sep 2011 12:45:55</date></shout><shout><body>We All Sleep Alone ♥</body><author>MisterBoombox</author><date>Sun, 28 Aug 2011 15:12:59</date></shout><shout><body>Believe &amp; Strong Enough ♥ OMG! perfect songs</body><author>JFERREIRA87</author><date>Fri, 24 Jun 2011 00:53:11</date></shout><shout><body>divina.</body><author>Cvilla69</author><date>Thu, 10 Feb 2011 19:39:12</date></shout><shout><body>yeah it really rocks [2]</body><author>top20fanatico</author><date>Wed, 15 Dec 2010 21:38:31</date></shout><shout><body>I remember this album from my childhood, it was on tape and was played 24/7 by my parents in home as well as in car :D all of us really loved it. and I still think it's very good although my music taste is very different now.</body><author>IrminkaT</author><date>Sat, 24 Jul 2010 17:37:07</date></shout><shout><body>one of the best pop albums.</body><author>gokhanesque</author><date>Mon, 5 Jul 2010 21:14:48</date></shout><shout><body>Believe ♥</body><author>marcio11</author><date>Sun, 13 Jun 2010 21:39:19</date></shout><shout><body>mi mucica preferedia</body><author>marina0179</author><date>Sun, 16 May 2010 01:22:12</date></shout><shout><body>Depois de tantos anos de carreira, é desnecessário dizer que o que quer que a diva Cher faça, ela o fará bem. Desde do pop otimista da dupla com Sony, passando por rainha das pistas da disco no final dos setenta, até o seu lado hard rock produzidos e compostos por Bon Jovi e Desmond Child. O single hit “Believe” chegou fuzilando as rádios com seus efeitos vocais futurísticos que são usados até os dias de hoje sem nenhuma timidez. O ábum ainda conta com o mega sucesso “Strong Enough” super dançante com cara de Barry White do futuro. Sua voz totalmente diferente continua a surpreender qualquer amador desconfiado e está muito bem encaixada nessa nova jornada dance de sua carreira. “The Power”, “Takin’ Back My Heart”  e “Runaway” poderiam ter se tornados grandes sucessos também, mas ainda empolgam os fãs desse trabalho. O álbum ainda conta com a boa versão dance de “We All Sleep Alone” de seu álbum “Cher” de 87.</body><author>ClaudioGurgel</author><date>Tue, 23 Mar 2010 20:02:54</date></shout><shout><body>I Love Strong Enough</body><author>rofeki</author><date>Sat, 13 Mar 2010 02:51:27</date></shout><shout><body>Strong enough .. gotta love it :)</body><author>Tracelace</author><date>Fri, 12 Mar 2010 21:52:41</date></shout><shout><body>Strong Enough&lt;3</body><author>blackout_rem</author><date>Fri, 6 Nov 2009 19:27:04</date></shout><shout><body>Runaway, Taxi Taxi, All or nothing &lt;3&lt;3&lt;3</body><author>arch_enemy_1987</author><date>Tue, 15 Sep 2009 19:13:27</date></shout><shout><body>Me encanta la voz de esta mujer.
Y como actriz también me gusta.
Beleve es espectacular.....
</body><author>PELIRROJAROJA</author><date>Fri, 16 Jan 2009 02:23:04</date></shout><shout><body>believe album is not bad but not.com.mercial,livin' proof,it's a man's world,stars,i'd rather believe in you,prisoner,cher,heart of stone,love hurts albums are better :D hihi what else :D</body><author>CherFanGay</author><date>Fri, 21 Nov 2008 09:57:19</date></shout><shout><body>yeah it really rocks</body><author>blond-ambition</author><date>Thu, 18 Sep 2008 17:11:37</date></shout><shout><body>i love taxi taxi:)</body><author>CherFanGay</author><date>Tue, 29 Jan 2008 13:59:32</date></shout><shout><body>did cher die?</body><author>rgurney</author><date>Mon, 16 Jul 2007 05:04:56</date></shout><shout><body>I love Taxi Taxi so much!!x</body><author>Maxes</author><date>Sat, 26 May 2007 23:53:57</date></shout><shout><body>Her face creeps me out. But I love her music. </body><author>TMNWY</author><date>Sun, 20 May 2007 21:24:52</date></shout><shout><body>strong voice,great vocalist as if u need telling.kaz07</body><author>kaz07</author><date>Mon, 30 Apr 2007 02:17:31</date></shout><shout><body>No comments for Believe? Nice album, but Living Proof is better ;&gt;</body><author>IParalyze</author><date>Tue, 30 Jan 2007 14:21:10</date></shout></shouts></lfm>`
	shouts := []Shout{}
	err := xml.Unmarshal([]byte(respString), &struct {
		Shouts *[]Shout `xml:"shouts>shout"`
	}{Shouts: &shouts})
	c.Assert(err, IsNil)
	c.Check(shouts, HasLen, 30)
}

func (s *S) TestGetShoutsE2E(c *C) {
	client := NewClient("d0a8c6b594b43669503d9f51aaabea22", "")
	shouts, err := client.Album().GetShouts("Radiohead", "OK Computer", false, 1)
	c.Assert(err, IsNil)
	c.Assert(shouts, HasLen, 50)
}
