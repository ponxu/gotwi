package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/tweets/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetCountsTweetsCountsRecent_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.TweetCountsTweetsCountsRecentResponse
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.TweetCountsTweetsCountsRecentResponse{
				Errors: []resources.PartialError{
					{Title: "test partical error"},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.TweetCountsTweetsCountsRecentResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.TweetCountsTweetsCountsRecentResponse{
				Errors: []resources.PartialError{}},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}
