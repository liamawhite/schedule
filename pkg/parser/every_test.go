package parser_test

import (
	"testing"
	"time"

	"github.com/liamawhite/schedule/pkg/date"
	"github.com/liamawhite/schedule/pkg/lexer"
	"github.com/liamawhite/schedule/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestParse_EveryStatements(t *testing.T) {
    relativeTo, _ := time.Parse(time.RFC3339, "2020-01-02T00:00:00Z") // thurs

    tests := []struct {
        input   string
        want    []*date.Date
        end     time.Time
        wantErr bool
    }{
        {
            input: "every day",
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
            input: "every week",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 1, Day: 9},
                {Year: 2020, Month: 1, Day: 16},
                {Year: 2020, Month: 1, Day: 23},
                {Year: 2020, Month: 1, Day: 30},
                {Year: 2020, Month: 2, Day: 6},
                {Year: 2020, Month: 2, Day: 13},
                {Year: 2020, Month: 2, Day: 20},
            },
            end: time.Date(2020, 2, 20, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every month",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 2, Day: 2},
                {Year: 2020, Month: 3, Day: 2},
                {Year: 2020, Month: 4, Day: 2},
                {Year: 2020, Month: 5, Day: 2},
                {Year: 2020, Month: 6, Day: 2},
                {Year: 2020, Month: 7, Day: 2},
                {Year: 2020, Month: 8, Day: 2},
            },
            end: time.Date(2020, 8, 2, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every year",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2021, Month: 1, Day: 2},
                {Year: 2022, Month: 1, Day: 2},
                {Year: 2023, Month: 1, Day: 2},
                {Year: 2024, Month: 1, Day: 2},
                {Year: 2025, Month: 1, Day: 2},
                {Year: 2026, Month: 1, Day: 2},
                {Year: 2027, Month: 1, Day: 2},
            },  
            end: time.Date(2027, 1, 2, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every weekday",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 1, Day: 3},
                {Year: 2020, Month: 1, Day: 6},
                {Year: 2020, Month: 1, Day: 7},
                {Year: 2020, Month: 1, Day: 8},
                {Year: 2020, Month: 1, Day: 9},
                {Year: 2020, Month: 1, Day: 10},
                {Year: 2020, Month: 1, Day: 13},
            },
            end: time.Date(2020, 1, 13, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every weekend",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 4},
                {Year: 2020, Month: 1, Day: 11},
                {Year: 2020, Month: 1, Day: 18},
                {Year: 2020, Month: 1, Day: 25},
                {Year: 2020, Month: 2, Day: 1},
                {Year: 2020, Month: 2, Day: 8},
                {Year: 2020, Month: 2, Day: 15},
                {Year: 2020, Month: 2, Day: 22},
            },
            end: time.Date(2020, 2, 22, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every mon, tues, wed",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 6},
                {Year: 2020, Month: 1, Day: 7},
                {Year: 2020, Month: 1, Day: 8},
                {Year: 2020, Month: 1, Day: 13},
                {Year: 2020, Month: 1, Day: 14},
                {Year: 2020, Month: 1, Day: 15},
                {Year: 2020, Month: 1, Day: 20},
                {Year: 2020, Month: 1, Day: 21},
            },
            end: time.Date(2020, 1, 21, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every tues, thurs",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 1, Day: 7},
                {Year: 2020, Month: 1, Day: 9},
                {Year: 2020, Month: 1, Day: 14},
                {Year: 2020, Month: 1, Day: 16},
                {Year: 2020, Month: 1, Day: 21},
                {Year: 2020, Month: 1, Day: 23},
                {Year: 2020, Month: 1, Day: 28},
                {Year: 2020, Month: 1, Day: 30},
            },
            end: time.Date(2020, 1, 30, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every sat/sun",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 4},
                {Year: 2020, Month: 1, Day: 5},
                {Year: 2020, Month: 1, Day: 11},
                {Year: 2020, Month: 1, Day: 12},
                {Year: 2020, Month: 1, Day: 18},
                {Year: 2020, Month: 1, Day: 19},
                {Year: 2020, Month: 1, Day: 25},
                {Year: 2020, Month: 1, Day: 26},
            },
            end: time.Date(2020, 1, 26, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every 6 days",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 1, Day: 8},
                {Year: 2020, Month: 1, Day: 14},
                {Year: 2020, Month: 1, Day: 20},
                {Year: 2020, Month: 1, Day: 26},
                {Year: 2020, Month: 2, Day: 1},
                {Year: 2020, Month: 2, Day: 7},
                {Year: 2020, Month: 2, Day: 13},
            },
            end: time.Date(2020, 2, 13, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every 2 weeks",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 1, Day: 16},
                {Year: 2020, Month: 1, Day: 30},
                {Year: 2020, Month: 2, Day: 13},
                {Year: 2020, Month: 2, Day: 27},
                {Year: 2020, Month: 3, Day: 12},
                {Year: 2020, Month: 3, Day: 26},
                {Year: 2020, Month: 4, Day: 9},
            },  
            end: time.Date(2020, 4, 9, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every 3 months",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2020, Month: 4, Day: 2},
                {Year: 2020, Month: 7, Day: 2},
                {Year: 2020, Month: 10, Day: 2},
                {Year: 2021, Month: 1, Day: 2},
                {Year: 2021, Month: 4, Day: 2},
                {Year: 2021, Month: 7, Day: 2},
                {Year: 2021, Month: 10, Day: 2},
            },  
            end: time.Date(2021, 10, 2, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every 2 years",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 2},
                {Year: 2022, Month: 1, Day: 2},
                {Year: 2024, Month: 1, Day: 2},
                {Year: 2026, Month: 1, Day: 2},
                {Year: 2028, Month: 1, Day: 2},
                {Year: 2030, Month: 1, Day: 2},
                {Year: 2032, Month: 1, Day: 2},
                {Year: 2034, Month: 1, Day: 2},
            },
            end: time.Date(2034, 1, 2, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every jan mar sept",
            want: []*date.Date{
                {Year: 2020, Month: 3, Day: 1},
                {Year: 2020, Month: 9, Day: 1},
                {Year: 2021, Month: 1, Day: 1},
                {Year: 2021, Month: 3, Day: 1},
                {Year: 2021, Month: 9, Day: 1},
                {Year: 2022, Month: 1, Day: 1},
                {Year: 2022, Month: 3, Day: 1},
                {Year: 2022, Month: 9, Day: 1},
            },
            end: time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every 1st, 15th",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 15},
                {Year: 2020, Month: 2, Day: 1},
                {Year: 2020, Month: 2, Day: 15},
                {Year: 2020, Month: 3, Day: 1},
                {Year: 2020, Month: 3, Day: 15},
                {Year: 2020, Month: 4, Day: 1},
                {Year: 2020, Month: 4, Day: 15},
                {Year: 2020, Month: 5, Day: 1},
            },
            end: time.Date(2020, 5, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            input: "every 1st, 15th jan, dec",
            want: []*date.Date{
                {Year: 2020, Month: 1, Day: 15},
                {Year: 2020, Month: 12, Day: 1},
                {Year: 2020, Month: 12, Day: 15},
                {Year: 2021, Month: 1, Day: 1},
                {Year: 2021, Month: 1, Day: 15},
                {Year: 2021, Month: 12, Day: 1},
                {Year: 2021, Month: 12, Day: 15},
                {Year: 2022, Month: 1, Day: 1},
            },  
            end: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
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
