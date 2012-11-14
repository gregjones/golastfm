package lastfm

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct{}

var (
	_          = Suite(&S{})
	testKey    = "d0a8c6b594b43669503d9f51aaabea22"
	testClient *Client
)

func (s *S) TestInvalidMethodError(c *C) {
	wrapper := struct{ Foo string }{}
	query := map[string]string{
		"method": "foo",
	}
	// tc := NewClient(testKey, "")
	err := testClient.Execute(query, &wrapper)
	c.Assert(err, NotNil)
}

func init() {
	testClient = NewClient(testKey, "")
}
