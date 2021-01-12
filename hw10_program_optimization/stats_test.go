// +build !bench

package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDomainStat(t *testing.T) {
	data := `{"Id":1,"Name":"Howard Mendoza","Username":"0Oliver","Email":"aliquid_qui_ea@Browsedrive.gov","Phone":"6-866-899-36-79","Password":"InAQJvsq","Address":"Blackbird Place 25"}
{"Id":2,"Name":"Jesse Vasquez","Username":"qRichardson","Email":"mLynch@broWsecat.com","Phone":"9-373-949-64-00","Password":"SiZLeNSGn","Address":"Fulton Hill 80"}
{"Id":3,"Name":"Clarence Olson","Username":"RachelAdams","Email":"RoseSmith@Browsecat.com","Phone":"988-48-97","Password":"71kuz3gA5w","Address":"Monterey Park 39"}
{"Id":4,"Name":"Gregory Reid","Username":"tButler","Email":"5Moore@Teklist.net","Phone":"520-04-16","Password":"r639qLNu","Address":"Sunfield Park 20"}
{"Id":5,"Name":"Janice Rose","Username":"KeithHart","Email":"nulla@Linktype.com","Phone":"146-91-01","Password":"acSBF5","Address":"Russell Trail 61"}`

	t.Run("find 'com'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "com")
		require.NoError(t, err)
		require.Equal(t, DomainStat{
			"browsecat.com": 2,
			"linktype.com":  1,
		}, result)
	})

	t.Run("find 'gov'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "gov")
		require.NoError(t, err)
		require.Equal(t, DomainStat{"browsedrive.gov": 1}, result)
	})

	t.Run("find 'unknown'", func(t *testing.T) {
		result, err := GetDomainStat(bytes.NewBufferString(data), "unknown")
		require.NoError(t, err)
		require.Equal(t, DomainStat{}, result)
	})
}

func Test_parseEmailDomain(t *testing.T) {
	tests := []struct {
		name        string
		in          []byte
		expectedRes []byte
		expectedErr error
	}{
		{
			"empty input",
			[]byte{},
			nil,
			ErrMalformedData,
		},
		{
			"no { at the beginning of input",
			[]byte(`"Id":1,"Name":"name","Username":"username","Email":"nobody@nowhere.org","Phone":"1234567890","Password":"passwd","Address":"address"}`),
			nil,
			ErrMalformedData,
		},
		{
			"no } at the end of input",
			[]byte(`{"Id":1,"Name":"name","Username":"username","Email":"nobody@nowhere.org","Phone":"1234567890","Password":"passwd","Address":"address"`),
			nil,
			ErrMalformedData,
		},
		{
			"no Email field in input",
			[]byte(`{"Id":1,"Name":"name","Username":"username","Phone":"1234567890","Password":"passwd","Address":"address"}`),
			nil,
			ErrMalformedData,
		},
		{
			"no \" at the end of email value",
			[]byte(`{"Id":1,"Name":"name","Username":"username","Email":"nobody@nowhere.org,"Phone":"1234567890","Password":"passwd","Address":"address"}`),
			nil,
			ErrMalformedData,
		},
		{
			"no @ in email",
			[]byte(`{"Id":1,"Name":"name","Username":"username","Email":"nobody","Phone":"1234567890","Password":"passwd","Address":"address"}`),
			nil,
			ErrMalformedData,
		},
		{
			"empty email domain part",
			[]byte(`{"Id":1,"Name":"name","Username":"username","Email":"nobody@","Phone":"1234567890","Password":"passwd","Address":"address"}`),
			nil,
			ErrMalformedData,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("case [%s]", tt.name), func(t *testing.T) {
			actual, err := parseEmailDomain(tt.in)
			if err == nil {
				assert.Equal(t, tt.expectedRes, actual, "correct result")
			} else {
				assert.True(t, errors.Is(tt.expectedErr, err), "correct error")
			}
		})
	}
}
