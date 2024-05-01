package main

import (
	"log"
	"testing"
)

func Test_unpack(t *testing.T) {
	testTable := []struct {
		input  string
		expect string
	}{
		{
			input:  "a4bc2d5e",
			expect: "aaaabccddddde",
		},
		{
			input:  "abcd",
			expect: "abcd",
		},
		{
			input:  "45",
			expect: "",
		},
		{
			input:  "",
			expect: "",
		},
		{
			input:  `qwe\4\5`,
			expect: `qwe45`,
		},
		{
			input:  `qwe\45`,
			expect: `qwe44444`,
		}, {
			input:  `qwe\\5`,
			expect: `qwe\\\\\`,
		},
	}

	for _, testCase := range testTable {
		output, err := unpack(testCase.input)
		if err != nil {
			log.Println(err)
		}
		if testCase.expect != output {
			t.Errorf("Incorrect result. Expect %s, got %s",
				testCase.expect,
				output)
		}
	}
}
