package main

import "testing"

func TestUnpackString(t *testing.T) {
	validData := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"":         "",
	}

	t.Run("valid_data", func(t *testing.T) {

		for inpStr, expectedStr := range validData {
			t.Run(inpStr, func(t *testing.T) {
				t.Parallel()
				gotStr, err := unpackString(inpStr)

				if err != nil {
					t.Errorf("Getted error from unpackString func: %s", err.Error())
				}

				if expectedStr != gotStr {
					t.Errorf("Incorrect result: \nexpected:%s\ngot:%s", expectedStr, gotStr)
				}

			})
		}

	})

	t.Run("incorrect_data", func(t *testing.T) {
		gotStr, err := unpackString("45")

		if err == nil {
			t.Errorf("unknown error with incorrect string")
		}

		if gotStr != "" {
			t.Errorf("Incorrect result: \nexpected:''\ngot:%s", gotStr)
		}

	})

}
