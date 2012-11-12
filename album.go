package lastfm

type albumMethods struct {
	Client *Client
}

func (q *albumMethods) Method(suffix string) string {
	return "album." + suffix
}

type affiliation struct {
	SupplierName string `json:"supplierName"`
	Price        struct {
		Currency  string `json:"currency:`
		Amount    string `json:"amount"`
		Formatted string `json:"formatted"`
	} `json:"price"`
	BuyLink      string `json:"buyLink"`
	SupplierIcon string `json:"supplierIcon"`
	IsSearch     string `json:"isSearch"`
}

type BuyLinksResponse struct {
	Affiliations struct {
		Physicals struct {
			Affiliation []affiliation `json:"affiliation"`
		} `json:"physicals"`
		Downloads struct {
			Affiliation []affiliation `json:"affiliation"`
		} `json:"downloads"`
	} `json:"affiliations"`
}

func (r *BuyLinksResponse) Physicals() []affiliation {
	return r.Affiliations.Physicals.Affiliation
}

func (r *BuyLinksResponse) Downloads() []affiliation {
	return r.Affiliations.Downloads.Affiliation
}

func (a *albumMethods) GetBuyLinks(artist, album string, autocorrect bool, country string) (*BuyLinksResponse, error) {
	resp := BuyLinksResponse{}
	var autocorrectString string
	if autocorrect {
		autocorrectString = "1"
	} else {
		autocorrectString = "0"
	}
	query := map[string]string{
		"method":      a.Method("getbuylinks"),
		"artist":      artist,
		"album":       album,
		"autocorrect": autocorrectString,
		"country":     country,
	}
	err := a.Client.Execute(query, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *albumMethods) GeyBuyLinksByMBID(mbid string, country string) {

}
