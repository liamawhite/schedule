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
				{Type: token.INT, Literal: "1"},
				{Type: token.ST, Literal: "st"},
				{Type: token.OF, Literal: "of"},
				{Type: token.THE, Literal: "the"},
				{Type: token.MONTH, Literal: "month"},
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
				{Type: token.MONDAY, Literal: "monday"},
				{Type: token.TUESDAY, Literal: "tuesday"},
				{Type: token.WEDNESDAY, Literal: "wednesday"},
				{Type: token.THURSDAY, Literal: "thursday"},
				{Type: token.FRIDAY, Literal: "friday"},
				{Type: token.SATURDAY, Literal: "saturday"},
				{Type: token.SUNDAY, Literal: "sunday"},
                {Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "mon tues tue weds wed thu thur thurs fri sat sun",
			want: []token.Token{
				{Type: token.MONDAY, Literal: "mon"},
				{Type: token.TUESDAY, Literal: "tues"},
				{Type: token.TUESDAY, Literal: "tue"},
				{Type: token.WEDNESDAY, Literal: "weds"},
				{Type: token.WEDNESDAY, Literal: "wed"},
				{Type: token.THURSDAY, Literal: "thu"},
				{Type: token.THURSDAY, Literal: "thur"},
				{Type: token.THURSDAY, Literal: "thurs"},
				{Type: token.FRIDAY, Literal: "fri"},
				{Type: token.SATURDAY, Literal: "sat"},
				{Type: token.SUNDAY, Literal: "sun"},
                {Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "january february march april may june july august september october november december",
			want: []token.Token{
				{Type: token.JANUARY, Literal: "january"},
				{Type: token.FEBRUARY, Literal: "february"},
				{Type: token.MARCH, Literal: "march"},
				{Type: token.APRIL, Literal: "april"},
				{Type: token.MAY, Literal: "may"},
				{Type: token.JUNE, Literal: "june"},
				{Type: token.JULY, Literal: "july"},
				{Type: token.AUGUST, Literal: "august"},
				{Type: token.SEPTEMBER, Literal: "september"},
				{Type: token.OCTOBER, Literal: "october"},
				{Type: token.NOVEMBER, Literal: "november"},
				{Type: token.DECEMBER, Literal: "december"},
                {Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "jan feb mar apr jun jul aug sep sept oct nov dec",
			want: []token.Token{
				{Type: token.JANUARY, Literal: "jan"},
				{Type: token.FEBRUARY, Literal: "feb"},
				{Type: token.MARCH, Literal: "mar"},
				{Type: token.APRIL, Literal: "apr"},
				{Type: token.JUNE, Literal: "jun"},
				{Type: token.JULY, Literal: "jul"},
				{Type: token.AUGUST, Literal: "aug"},
				{Type: token.SEPTEMBER, Literal: "sep"},
				{Type: token.SEPTEMBER, Literal: "sept"},
				{Type: token.OCTOBER, Literal: "oct"},
				{Type: token.NOVEMBER, Literal: "nov"},
				{Type: token.DECEMBER, Literal: "dec"},
                {Type: token.EOF, Literal: ""},
			},
		},
		{
			input: "1st 2Nd 3rD 4th 5th 6th 7th 8th 9th 10th 11th 12th 13th 14th 15th 16th 17th 18th 19th 20th 21st 22nd 23rd 24th 25th 26th 27th 28th 29th 30th 31st",
			want: []token.Token{
				{Type: token.INT, Literal: "1"},
				{Type: token.ST, Literal: "st"},
				{Type: token.INT, Literal: "2"},
				{Type: token.ND, Literal: "Nd"},
				{Type: token.INT, Literal: "3"},
				{Type: token.RD, Literal: "rD"},
				{Type: token.INT, Literal: "4"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "5"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "6"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "7"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "8"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "9"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "10"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "11"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "12"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "13"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "14"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "15"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "16"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "17"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "18"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "19"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "20"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "21"},
				{Type: token.ST, Literal: "st"},
				{Type: token.INT, Literal: "22"},
				{Type: token.ND, Literal: "nd"},
				{Type: token.INT, Literal: "23"},
				{Type: token.RD, Literal: "rd"},
				{Type: token.INT, Literal: "24"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "25"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "26"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "27"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "28"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "29"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "30"},
				{Type: token.TH, Literal: "th"},
				{Type: token.INT, Literal: "31"},
				{Type: token.ST, Literal: "st"},
                {Type: token.EOF, Literal: ""},
			},
		},
	}
	for _, tt := range tests {
		l := New(tt.input)
		for i, want := range tt.want {
			got := l.NextToken()
			if got != want {
				t.Errorf("test %d: got %+v, want %+v", i, got, want)
			}
		}
	}
}
