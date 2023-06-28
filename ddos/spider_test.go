package ddos

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DdosTestSuite struct {
	suite.Suite
}

func TestDdosTestSuite(t *testing.T) {
	suite.Run(t, new(DdosTestSuite))
}

func (suite *DdosTestSuite) TestGetUserAgent_ShouldReturnNonEmptyString() {
	userAgent := getUserAgent()
	suite.NotEmpty(userAgent, "Expected a non-empty string")
}

func (suite *DdosTestSuite) TestGetOSForPlatform_ShouldReturnExpectedValues() {
	testCases := []struct {
		platform string
	}{
		{"Macintosh"},
		{"Windows"},
		{"X11"},
		{"Random"},
	}

	for _, tc := range testCases {
		os := getOSForPlatform(tc.platform)
		if tc.platform == "Random" {
			suite.Empty(os, "Expected an empty string for unknown platform")
		} else {
			suite.NotEmpty(os, "Expected a non-empty string for platform %s", tc.platform)
		}
	}
}

func (suite *DdosTestSuite) TestBuildChromeUserAgent_ShouldReturnChromeUserAgentStringInExpectedFormat() {
	userAgent := buildChromeUserAgent("Mac OS X")
	suite.Contains(userAgent, "Mozilla/5.0 (Mac OS X) AppleWebKit/")
	suite.Contains(userAgent, "Chrome/")
	suite.Contains(userAgent, "Safari/")
}

func (suite *DdosTestSuite) TestBuildIEUserAgent_ShouldReturnWindowsUserAgentStringInExpectedFormat() {
	userAgent := buildIEUserAgent("Windows NT 10.0")
	suite.Contains(userAgent, "Mozilla/5.0 (compatible; MSIE")
	suite.Contains(userAgent, "Windows NT 10.0")
	suite.Contains(userAgent, "Trident/")
}
