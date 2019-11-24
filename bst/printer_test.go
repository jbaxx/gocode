package bsarbol

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintDataBytes(t *testing.T) {

	cases := []struct {
		name string
		data int
		max  int
		want string
	}{
		{
			name: "Data: 2 runes, max: 2 runes",
			data: 12,
			max:  12,
			want: "[12]",
		},
		{
			name: "Data: 1 runes, max: 2 runes",
			data: 7,
			max:  12,
			want: "[ 7]",
		},
		{
			name: "Data: 1 runes, max: 3 runes",
			data: 3,
			max:  124,
			want: "[  3]",
		},
		{
			name: "Data: 1 runes, max: 6 runes",
			data: 3,
			max:  124497,
			want: "[     3]",
		},
		{
			name: "Data: 2 runes, max: 6 runes",
			data: 32,
			max:  124497,
			want: "[    32]",
		},
	}

	buffer := bytes.Buffer{}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {

			printData(&buffer, test.data, test.max)
			got := buffer.String()
			got = strings.TrimSuffix(got, "\n")
			defer buffer.Reset()

			if got != test.want {
				t.Errorf("got %s; want %s", got, test.want)
			}

		})
	}

	t.Run("Expect error: Data: 3 runes, max: 2 runes", func(t *testing.T) {

		data := 323
		max := 12

		err := printData(&buffer, data, max)
		got := buffer.String()
		got = strings.TrimSuffix(got, "\n")
		defer buffer.Reset()

		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if e, ok := err.(*ErrMaxRuneCount); !ok {
			t.Errorf("expected error type: *ErrMaxRuneCount; got instead %#v AND %#v", e, err)
		}
	})

}

func BenchmarkPrintDataBytesMath(b *testing.B) {
	builder := strings.Builder{}

	data := 32
	max := 124497

	for i := 0; i < b.N; i++ {
		printData(&builder, data, max)
		builder.Reset()
	}

}
