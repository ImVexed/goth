package azureadv2_test

import (
	"testing"

	"github.com/ImVexed/goth"
	"github.com/ImVexed/goth/providers/azureadv2"
	"github.com/stretchr/testify/assert"
)

func Test_Implements_Session(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &azureadv2.Session{}

	a.Implements((*goth.Session)(nil), s)
}

func Test_GetAuthURL(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &azureadv2.Session{}

	_, err := s.GetAuthURL()
	a.Error(err)

	s.AuthURL = "/foo"

	url, _ := s.GetAuthURL()
	a.Equal(url, "/foo")
}

func Test_ToJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &azureadv2.Session{}

	data := s.Marshal()
	a.Equal(`{"au":"","at":"","rt":"","exp":"0001-01-01T00:00:00Z"}`, data)
}

func Test_String(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &azureadv2.Session{}

	a.Equal(s.String(), s.Marshal())
}
