package lexer

import (
	"testing"

	"github.com/liamawhite/schedule/pkg/token"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		input string
		want  []token.Token
	}{
		{
			input: "every 1st of the month",
			want: []token.Token{
				{Type: token.EVERY, Literal: "every"},
				{Type: token.NUMBER, Literal: "1"},
                {Type: token.ST, Literal: "st"},
				{Type: token.OF, Literal: "of"},
				{Type: token.THE, Literal: "the"},
				{Type: token.MONTH, Literal: "month"},
				{Type: token.EOF, Literal: ""},
			},
		},
        {
            input: "daily weekly monthly yearly",
            want: []token.Token{
                {Type: token.DAILY, Literal: "daily"},
                {Type: token.WEEKLY, Literal: "weekly"},
                {Type: token.MONTHLY, Literal: "monthly"},
                {Type: token.YEARLY, Literal: "yearly"},
                {Type: token.EOF, Literal: ""},
            },
        },
		{
			input: "somerandomword",
			want: []token.Token{
				{Type: token.ILLEGAL, Literal: "somerandomword"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: ",/",
			want: []token.Token{
				{Type: token.COMMA, Literal: ","},
				{Type: token.SLASH, Literal: "/"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "DaY",
			want: []token.Token{
				{Type: token.DAY, Literal: "DaY"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "every in next OTHER oN Of ThE",
			want: []token.Token{
				{Type: token.EVERY, Literal: "every"},
				{Type: token.IN, Literal: "in"},
				{Type: token.NEXT, Literal: "next"},
				{Type: token.OTHER, Literal: "OTHER"},
				{Type: token.ON, Literal: "oN"},
				{Type: token.OF, Literal: "Of"},
				{Type: token.THE, Literal: "ThE"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "day days week weeks month months year years",
			want: []token.Token{
				{Type: token.DAY, Literal: "day"},
				{Type: token.DAY, Literal: "days"},
				{Type: token.WEEK, Literal: "week"},
				{Type: token.WEEK, Literal: "weeks"},
				{Type: token.MONTH, Literal: "month"},
				{Type: token.MONTH, Literal: "months"},
				{Type: token.YEAR, Literal: "year"},
				{Type: token.YEAR, Literal: "years"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "monday tuesday wednesday thursday friday saturday sunday",
			want: []token.Token{
				{Type: token.DAYOFWEEK, Literal: "monday"},
				{Type: token.DAYOFWEEK, Literal: "tuesday"},
				{Type: token.DAYOFWEEK, Literal: "wednesday"},
				{Type: token.DAYOFWEEK, Literal: "thursday"},
				{Type: token.DAYOFWEEK, Literal: "friday"},
				{Type: token.DAYOFWEEK, Literal: "saturday"},
				{Type: token.DAYOFWEEK, Literal: "sunday"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "mon tues tue weds wed thu thur thurs fri sat sun",
			want: []token.Token{
				{Type: token.DAYOFWEEK, Literal: "mon"},
				{Type: token.DAYOFWEEK, Literal: "tues"},
				{Type: token.DAYOFWEEK, Literal: "tue"},
				{Type: token.DAYOFWEEK, Literal: "weds"},
				{Type: token.DAYOFWEEK, Literal: "wed"},
				{Type: token.DAYOFWEEK, Literal: "thu"},
				{Type: token.DAYOFWEEK, Literal: "thur"},
				{Type: token.DAYOFWEEK, Literal: "thurs"},
				{Type: token.DAYOFWEEK, Literal: "fri"},
				{Type: token.DAYOFWEEK, Literal: "sat"},
				{Type: token.DAYOFWEEK, Literal: "sun"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "january february march april may june july august september october november december",
			want: []token.Token{
				{Type: token.MONTHOFYEAR, Literal: "january"},
				{Type: token.MONTHOFYEAR, Literal: "february"},
				{Type: token.MONTHOFYEAR, Literal: "march"},
				{Type: token.MONTHOFYEAR, Literal: "april"},
				{Type: token.MONTHOFYEAR, Literal: "may"},
				{Type: token.MONTHOFYEAR, Literal: "june"},
				{Type: token.MONTHOFYEAR, Literal: "july"},
				{Type: token.MONTHOFYEAR, Literal: "august"},
				{Type: token.MONTHOFYEAR, Literal: "september"},
				{Type: token.MONTHOFYEAR, Literal: "october"},
				{Type: token.MONTHOFYEAR, Literal: "november"},
				{Type: token.MONTHOFYEAR, Literal: "december"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "jan feb mar apr jun jul aug sep sept oct nov dec",
			want: []token.Token{
				{Type: token.MONTHOFYEAR, Literal: "jan"},
				{Type: token.MONTHOFYEAR, Literal: "feb"},
				{Type: token.MONTHOFYEAR, Literal: "mar"},
				{Type: token.MONTHOFYEAR, Literal: "apr"},
				{Type: token.MONTHOFYEAR, Literal: "jun"},
				{Type: token.MONTHOFYEAR, Literal: "jul"},
				{Type: token.MONTHOFYEAR, Literal: "aug"},
				{Type: token.MONTHOFYEAR, Literal: "sep"},
				{Type: token.MONTHOFYEAR, Literal: "sept"},
				{Type: token.MONTHOFYEAR, Literal: "oct"},
				{Type: token.MONTHOFYEAR, Literal: "nov"},
				{Type: token.MONTHOFYEAR, Literal: "dec"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "1st 2Nd 3rD 4th 5th 6th 7th 8th 9th 10th 11th 12th 13th 14th 15th 16th 17th 18th 19th 20th 21st 22nd 23rd 24th 25th 26th 27th 28th 29th 30th 31st",
			want: []token.Token{
				{Type: token.NUMBER, Literal: "1"},
                {Type: token.ST, Literal: "st"},
				{Type: token.NUMBER, Literal: "2"},
                {Type: token.ND, Literal: "Nd"},
				{Type: token.NUMBER, Literal: "3"},
                {Type: token.RD, Literal: "rD"},
				{Type: token.NUMBER, Literal: "4"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "5"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "6"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "7"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "8"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "9"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "10"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "11"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "12"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "13"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "14"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "15"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "16"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "17"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "18"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "19"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "20"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "21"},
                {Type: token.ST, Literal: "st"},
				{Type: token.NUMBER, Literal: "22"},
                {Type: token.ND, Literal: "nd"},
				{Type: token.NUMBER, Literal: "23"},
                {Type: token.RD, Literal: "rd"},
				{Type: token.NUMBER, Literal: "24"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "25"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "26"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "27"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "28"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "29"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "30"},
                {Type: token.TH, Literal: "th"},
				{Type: token.NUMBER, Literal: "31"},
                {Type: token.ST, Literal: "st"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "weekday weekend",
			want: []token.Token{
				{Type: token.SECTIONOFWEEK, Literal: "weekday"},
				{Type: token.SECTIONOFWEEK, Literal: "weekend"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "thurs",
			want: []token.Token{
				{Type: token.DAYOFWEEK, Literal: "thurs"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			l := New(tt.input)
			for i, want := range tt.want {
				got := l.NextToken()
				if got != want {
					t.Errorf("test %d: got %+v, want %+v", i, got, want)
				}
			}
		})
	}
}
