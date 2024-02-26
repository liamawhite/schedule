package parser_test

import (
	"testing"
	"time"

	"github.com/liamawhite/schedule/pkg/lexer"
    "github.com/liamawhite/schedule/pkg/parser"
	"github.com/liamawhite/schedule/pkg/date"
	"github.com/stretchr/testify/assert"
)

func TestParseDate_ExpressionStatement(t *testing.T) {
	relativeTo, _ := time.Parse(time.RFC3339, "2020-01-02T00:00:00Z") // thurs

	tests := []struct {
		input   string
		want    *date.Date
		wantErr bool
	}{
		{
			input: "monday",
			want:  &date.Date{Year: 2020, Month: 1, Day: 6},
		},
		{
			input: "thurs",
			want:  &date.Date{Year: 2020, Month: 1, Day: 9},
		},
		{
			input: "sat",
			want:  &date.Date{Year: 2020, Month: 1, Day: 4},
		},
		{
			input: "week",
			want:  &date.Date{Year: 2020, Month: 1, Day: 6},
		},
		{
			input: "month",
			want:  &date.Date{Year: 2020, Month: 2, Day: 1},
		},
		{
			input: "year",
			want:  &date.Date{Year: 2021, Month: 1, Day: 1},
		},
		{
			input: "weekend",
			want:  &date.Date{Year: 2020, Month: 1, Day: 4},
		},
		{
			input: "weekday",
			want:  &date.Date{Year: 2020, Month: 1, Day: 3},
		},
		{
			input: "jan",
			want:  &date.Date{Year: 2021, Month: 1, Day: 1},
		},
		{
			input: "sept",
			want:  &date.Date{Year: 2020, Month: 9, Day: 1},
		},
		{
			input: "25",
			want:  &date.Date{Year: 2020, Month: 1, Day: 25},
		},
		{
			input: "1st",
			want:  &date.Date{Year: 2020, Month: 2, Day: 1},
		},
		{
			input: "2nd",
			want:  &date.Date{Year: 2020, Month: 2, Day: 2},
		},
		{
			input: "1st jan",
			want:  &date.Date{Year: 2021, Month: 1, Day: 1},
		},
		{
			input: "1 jan 2020",
			want:  &date.Date{Year: 2020, Month: 1, Day: 1},
		},
		{
			input: "2nd feb 2020",
			want:  &date.Date{Year: 2020, Month: 2, Day: 2},
		},
		{
			input: "3rd mar",
			want:  &date.Date{Year: 2020, Month: 3, Day: 3},
		},
		{
			input: "april 4",
			want:  &date.Date{Year: 2020, Month: 4, Day: 4},
		},
		{
			input: "may 5 2020",
			want:  &date.Date{Year: 2020, Month: 5, Day: 5},
		},
		{
			input: "jun 2020",
			want:  &date.Date{Year: 2020, Month: 6, Day: 1},
		},
		{
			input: "july 7th",
			want:  &date.Date{Year: 2020, Month: 7, Day: 7},
		},
		{
			input: "Aug 8th 2020",
			want:  &date.Date{Year: 2020, Month: 8, Day: 8},
		},
		{
			input: "2020",
			want:  &date.Date{Year: 2020, Month: 1, Day: 1},
		},
		{
			input:   "notathing",
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

func TestParse_ExpressionStatement(t *testing.T) {
    relativeTo, _ := time.Parse(time.RFC3339, "2020-01-02T00:00:00Z") // thurs

    tests := []struct {
        input   string
        want    []*date.Date
        end     time.Time
        wantErr bool
    }{
        {
            input: "daily",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 1, Day: 3},
                {Year: 2020, Month: 1, Day: 4},
                {Year: 2020, Month: 1, Day: 5},
                {Year: 2020, Month: 1, Day: 6},
                {Year: 2020, Month: 1, Day: 7},
                {Year: 2020, Month: 1, Day: 8},
                {Year: 2020, Month: 1, Day: 9},
            },
            end: time.Date(2020, 1, 9, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "weekly",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 1, Day: 9},
                {Year: 2020, Month: 1, Day: 16},
                {Year: 2020, Month: 1, Day: 23},
                {Year: 2020, Month: 1, Day: 30},
                {Year: 2020, Month: 2, Day: 6},
                {Year: 2020, Month: 2, Day: 13},
            },
            end: time.Date(2020, 2, 13, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "monthly",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 2, Day: 2},
                {Year: 2020, Month: 3, Day: 2},
                {Year: 2020, Month: 4, Day: 2},
                {Year: 2020, Month: 5, Day: 2},
                {Year: 2020, Month: 6, Day: 2},
                {Year: 2020, Month: 7, Day: 2},
                {Year: 2020, Month: 8, Day: 2},
                {Year: 2020, Month: 9, Day: 2},
                {Year: 2020, Month: 10, Day: 2},
                {Year: 2020, Month: 11, Day: 2},
                {Year: 2020, Month: 12, Day: 2},
                {Year: 2021, Month: 1, Day: 2},
                {Year: 2021, Month: 2, Day: 2},
            },  
            end: time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "yearly",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2021, Month: 1, Day: 2},
                {Year: 2022, Month: 1, Day: 2},
                {Year: 2023, Month: 1, Day: 2},
                {Year: 2024, Month: 1, Day: 2},
                {Year: 2025, Month: 1, Day: 2},
                {Year: 2026, Month: 1, Day: 2},       
            },
            end: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            p := parser.New(lexer.New(tt.input))
            got, err := p.Parse(relativeTo)

            if !tt.wantErr && err != nil {
                t.Errorf("got error: %+v", err)
            }

            if tt.wantErr && err == nil {
                t.Errorf("expected an error but got none")
            }
            
            assert.Equal(t, tt.want, got.Between(relativeTo, tt.end, true))
        })
    }
}
