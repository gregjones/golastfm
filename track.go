package lastfm

type TrackMethods struct {
	Client *Client
}

func (t TrackMethods) method(suffix string) string {
	return "track." + suffix
}

func (t TrackMethods) getInfo(query map[string]string) (*TrackInfo, error) {
	wrapper := struct {
		Track *TrackInfo `xml:"track"`
	}{}
	query["method"] = t.method("getinfo")
	err := t.Client.Execute(query, &wrapper)
	if err != nil {
		return nil, err
	}
	return wrapper.Track, nil
}

func (t TrackMethods) GetInfo(track, artist string, autocorrect bool, username string) (*TrackInfo, error) {
	query := map[string]string{
		"track":       track,
		"artist":      artist,
		"autocorrect": boolToString(autocorrect),
		"username":    username,
	}
	return t.getInfo(query)
}

func (t TrackMethods) GetInfoByMBID(mbid, username string) (*TrackInfo, error) {
	query := map[string]string{
		"mbid":     mbid,
		"username": username,
	}
	return t.getInfo(query)
}
