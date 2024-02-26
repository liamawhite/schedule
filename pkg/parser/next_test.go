package parser_test

import (
	"testing"
	"time"

	"github.com/liamawhite/schedule/pkg/lexer"
	"github.com/liamawhite/schedule/pkg/parser"
	"github.com/liamawhite/schedule/pkg/date"
	"github.com/stretchr/testify/assert"
)

func TestParseDate_NextStatements(t *testing.T) {
	relativeTo, _ := time.Parse(time.RFC3339, "2020-01-02T00:00:00Z") // thurs

	tests := []struct {
		input   string
		want    *date.Date
		wantErr bool
	}{
		{
			input: "next monday",
			want:  &date.Date{Year: 2020, Month: 1, Day: 6},
		},
		{
			input: "next thurs",
			want:  &date.Date{Year: 2020, Month: 1, Day: 9},
		},
		{
			input: "next sat",
			want:  &date.Date{Year: 2020, Month: 1, Day: 4},
		},
		{
			input: "next week",
			want:  &date.Date{Year: 2020, Month: 1, Day: 6},
		},
		{
			input: "next month",
			want:  &date.Date{Year: 2020, Month: 2, Day: 1},
		},
		{
			input: "next jan",
			want:  &date.Date{Year: 2021, Month: 1, Day: 1},
		},
		{
			input: "next sept",
			want:  &date.Date{Year: 2020, Month: 9, Day: 1},
		},
		{
			input: "next 25",
			want:  &date.Date{Year: 2020, Month: 1, Day: 25},
		},
		{
			input: "next 1st",
			want:  &date.Date{Year: 2020, Month: 2, Day: 1},
		},
		{
			input: "next 2nd",
			want:  &date.Date{Year: 2020, Month: 2, Day: 2},
		},
		{
			input: "next year",
			want:  &date.Date{Year: 2021, Month: 1, Day: 1},
		},
		{
			input: "next weekend",
			want:  &date.Date{Year: 2020, Month: 1, Day: 4},
		},
		{
			input: "next weekday",
			want:  &date.Date{Year: 2020, Month: 1, Day: 3},
		},
		{
			input:   "next notathing",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := parser.New(lexer.New(tt.input))
			got, err := p.ParseDate(relativeTo)

			assert.Equal(t, tt.want, got)
			if !tt.wantErr && err != nil {
				t.Errorf("got error: %+v", err)
			}

			if tt.wantErr && err == nil {
				t.Errorf("expected an error but got none")
			}
		})
	}
}
