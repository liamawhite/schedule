package parser_test

import (
	"testing"
	"time"

	"github.com/liamawhite/schedule/pkg/lexer"
	"github.com/liamawhite/schedule/pkg/parser"
	"github.com/liamawhite/schedule/pkg/date"
	"github.com/stretchr/testify/assert"
)

func TestParseDate_InStatement(t *testing.T) {
	relativeTo, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")

	tests := []struct {
		input   string
		want    *date.Date
		wantErr bool
	}{
		{
			input: "in 6 days",
			want:  &date.Date{Year: 2020, Month: 1, Day: 7},
		},
		{
			input: "in 1 day",
			want:  &date.Date{Year: 2020, Month: 1, Day: 2},
		},
		{
			input: "in 2 weeks",
			want:  &date.Date{Year: 2020, Month: 1, Day: 15},
		},
		{
			input: "in 1 week",
			want:  &date.Date{Year: 2020, Month: 1, Day: 8},
		},
		{
			input: "in 3 months",
			want:  &date.Date{Year: 2020, Month: 4, Day: 1},
		},
		{
			input: "in 1 month",
			want:  &date.Date{Year: 2020, Month: 2, Day: 1},
		},
		{
			input: "in 1 year",
			want:  &date.Date{Year: 2021, Month: 1, Day: 1},
		},
		{
			input: "in 2 years",
			want:  &date.Date{Year: 2022, Month: 1, Day: 1},
		},
		{
			input:   "in days",
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
