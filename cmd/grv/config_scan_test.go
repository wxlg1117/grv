package main

import (
	"errors"
	"strings"
	"testing"
)

func TestScanSingleConfigToken(t *testing.T) {
	var singleTokenTests = []struct {
		input         string
		expectedToken ConfigToken
	}{
		{
			input: "-!\"word1世界",
			expectedToken: ConfigToken{
				tokenType: CtkWord,
				value:     "-!\"word1世界",
				startPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
				endPos: ConfigScannerPos{
					line: 1,
					col:  10,
				},
			},
		},
		{
			input: "\"word \\t\\\"with\\\"\\n spaces\"",
			expectedToken: ConfigToken{
				tokenType: CtkWord,
				value:     "word \t\"with\"\n spaces",
				startPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
				endPos: ConfigScannerPos{
					line: 1,
					col:  26,
				},
			},
		},
		{
			input: " \t\r\v\f",
			expectedToken: ConfigToken{
				tokenType: CtkWhiteSpace,
				value:     " \t\r\v\f",
				startPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
				endPos: ConfigScannerPos{
					line: 1,
					col:  5,
				},
			},
		},
		{
			input: "# Comment",
			expectedToken: ConfigToken{
				tokenType: CtkComment,
				value:     "# Comment",
				startPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
				endPos: ConfigScannerPos{
					line: 1,
					col:  9,
				},
			},
		},
		{
			input: "--option",
			expectedToken: ConfigToken{
				tokenType: CtkOption,
				value:     "--option",
				startPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
				endPos: ConfigScannerPos{
					line: 1,
					col:  8,
				},
			},
		},
		{
			input: "\n",
			expectedToken: ConfigToken{
				tokenType: CtkTerminator,
				value:     "\n",
				startPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
				endPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
			},
		},
		{
			input: "",
			expectedToken: ConfigToken{
				tokenType: CtkEOF,
				startPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
				endPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
			},
		},
		{
			input: "\"Unterminated string",
			expectedToken: ConfigToken{
				tokenType: CtkInvalid,
				value:     "\"Unterminated string",
				startPos: ConfigScannerPos{
					line: 1,
					col:  1,
				},
				endPos: ConfigScannerPos{
					line: 1,
					col:  20,
				},
				err: errors.New("Unterminated string"),
			},
		},
	}

	for _, singleTokenTest := range singleTokenTests {
		scanner := NewConfigScanner(strings.NewReader(singleTokenTest.input))
		token, err := scanner.Scan()

		if err != nil {
			t.Errorf("Scan failed with error %v", err)
		} else if !token.Equal(&singleTokenTest.expectedToken) {
			t.Errorf("ConfigToken does not match expected value. Expected %v, Actual %v", singleTokenTest.expectedToken, *token)
		}
	}
}

func TestScanMultipleConfigTokens(t *testing.T) {
	var multiTokenTests = []struct {
		input          string
		expectedTokens []ConfigToken
	}{
		{
			input: "theme --create \"my theme\"\n",
			expectedTokens: []ConfigToken{
				{
					tokenType: CtkWord,
					value:     "theme",
					startPos: ConfigScannerPos{
						line: 1,
						col:  1,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  5,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  6,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  6,
					},
				},
				{
					tokenType: CtkOption,
					value:     "--create",
					startPos: ConfigScannerPos{
						line: 1,
						col:  7,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  14,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  15,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  15,
					},
				},
				{
					tokenType: CtkWord,
					value:     "my theme",
					startPos: ConfigScannerPos{
						line: 1,
						col:  16,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  25,
					},
				},
				{
					tokenType: CtkTerminator,
					value:     "\n",
					startPos: ConfigScannerPos{
						line: 1,
						col:  26,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  26,
					},
				},
				{
					tokenType: CtkEOF,
					startPos: ConfigScannerPos{
						line: 1,
						col:  26,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  26,
					},
				},
			},
		},
		{
			input: "set theme mytheme\nset\tCommitView.dateformat \"%yyyy-mm-dd HH:MM\"\n",
			expectedTokens: []ConfigToken{
				{
					tokenType: CtkWord,
					value:     "set",
					startPos: ConfigScannerPos{
						line: 1,
						col:  1,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  3,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  4,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  4,
					},
				},
				{
					tokenType: CtkWord,
					value:     "theme",
					startPos: ConfigScannerPos{
						line: 1,
						col:  5,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  9,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  10,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  10,
					},
				},
				{
					tokenType: CtkWord,
					value:     "mytheme",
					startPos: ConfigScannerPos{
						line: 1,
						col:  11,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  17,
					},
				},
				{
					tokenType: CtkTerminator,
					value:     "\n",
					startPos: ConfigScannerPos{
						line: 1,
						col:  18,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  18,
					},
				},
				{
					tokenType: CtkWord,
					value:     "set",
					startPos: ConfigScannerPos{
						line: 2,
						col:  1,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  3,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     "\t",
					startPos: ConfigScannerPos{
						line: 2,
						col:  4,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  4,
					},
				},
				{
					tokenType: CtkWord,
					value:     "CommitView.dateformat",
					startPos: ConfigScannerPos{
						line: 2,
						col:  5,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  25,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 2,
						col:  26,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  26,
					},
				},
				{
					tokenType: CtkWord,
					value:     "%yyyy-mm-dd HH:MM",
					startPos: ConfigScannerPos{
						line: 2,
						col:  27,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  45,
					},
				},
				{
					tokenType: CtkTerminator,
					value:     "\n",
					startPos: ConfigScannerPos{
						line: 2,
						col:  46,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  46,
					},
				},
				{
					tokenType: CtkEOF,
					startPos: ConfigScannerPos{
						line: 2,
						col:  46,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  46,
					},
				},
			},
		},
		{
			input: "theme --create \\\n\tmytheme\n",
			expectedTokens: []ConfigToken{
				{
					tokenType: CtkWord,
					value:     "theme",
					startPos: ConfigScannerPos{
						line: 1,
						col:  1,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  5,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  6,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  6,
					},
				},
				{
					tokenType: CtkOption,
					value:     "--create",
					startPos: ConfigScannerPos{
						line: 1,
						col:  7,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  14,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " \t",
					startPos: ConfigScannerPos{
						line: 1,
						col:  15,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  1,
					},
				},
				{
					tokenType: CtkWord,
					value:     "mytheme",
					startPos: ConfigScannerPos{
						line: 2,
						col:  2,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  8,
					},
				},
				{
					tokenType: CtkTerminator,
					value:     "\n",
					startPos: ConfigScannerPos{
						line: 2,
						col:  9,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  9,
					},
				},
				{
					tokenType: CtkEOF,
					startPos: ConfigScannerPos{
						line: 2,
						col:  9,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  9,
					},
				},
			},
		},
		{
			input: "theme --create \"my theme\nset theme mytheme\n",
			expectedTokens: []ConfigToken{
				{
					tokenType: CtkWord,
					value:     "theme",
					startPos: ConfigScannerPos{
						line: 1,
						col:  1,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  5,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  6,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  6,
					},
				},
				{
					tokenType: CtkOption,
					value:     "--create",
					startPos: ConfigScannerPos{
						line: 1,
						col:  7,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  14,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  15,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  15,
					},
				},
				{
					tokenType: CtkInvalid,
					value:     "\"my theme\nset theme mytheme\n",
					startPos: ConfigScannerPos{
						line: 1,
						col:  16,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  18,
					},
					err: errors.New("Unterminated string"),
				},
				{
					tokenType: CtkEOF,
					startPos: ConfigScannerPos{
						line: 2,
						col:  18,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  18,
					},
				},
			},
		},
		{
			input: "set theme mytheme # Set theme \n # set theme again\nset theme mytheme #EOF",
			expectedTokens: []ConfigToken{
				{
					tokenType: CtkWord,
					value:     "set",
					startPos: ConfigScannerPos{
						line: 1,
						col:  1,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  3,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  4,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  4,
					},
				},
				{
					tokenType: CtkWord,
					value:     "theme",
					startPos: ConfigScannerPos{
						line: 1,
						col:  5,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  9,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  10,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  10,
					},
				},
				{
					tokenType: CtkWord,
					value:     "mytheme",
					startPos: ConfigScannerPos{
						line: 1,
						col:  11,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  17,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  18,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  18,
					},
				},
				{
					tokenType: CtkComment,
					value:     "# Set theme ",
					startPos: ConfigScannerPos{
						line: 1,
						col:  19,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  30,
					},
				},
				{
					tokenType: CtkTerminator,
					value:     "\n",
					startPos: ConfigScannerPos{
						line: 1,
						col:  31,
					},
					endPos: ConfigScannerPos{
						line: 1,
						col:  31,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 2,
						col:  1,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  1,
					},
				},
				{
					tokenType: CtkComment,
					value:     "# set theme again",
					startPos: ConfigScannerPos{
						line: 2,
						col:  2,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  18,
					},
				},
				{
					tokenType: CtkTerminator,
					value:     "\n",
					startPos: ConfigScannerPos{
						line: 2,
						col:  19,
					},
					endPos: ConfigScannerPos{
						line: 2,
						col:  19,
					},
				},
				{
					tokenType: CtkWord,
					value:     "set",
					startPos: ConfigScannerPos{
						line: 3,
						col:  1,
					},
					endPos: ConfigScannerPos{
						line: 3,
						col:  3,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 3,
						col:  4,
					},
					endPos: ConfigScannerPos{
						line: 3,
						col:  4,
					},
				},
				{
					tokenType: CtkWord,
					value:     "theme",
					startPos: ConfigScannerPos{
						line: 3,
						col:  5,
					},
					endPos: ConfigScannerPos{
						line: 3,
						col:  9,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 3,
						col:  10,
					},
					endPos: ConfigScannerPos{
						line: 3,
						col:  10,
					},
				},
				{
					tokenType: CtkWord,
					value:     "mytheme",
					startPos: ConfigScannerPos{
						line: 3,
						col:  11,
					},
					endPos: ConfigScannerPos{
						line: 3,
						col:  17,
					},
				},
				{
					tokenType: CtkWhiteSpace,
					value:     " ",
					startPos: ConfigScannerPos{
						line: 3,
						col:  18,
					},
					endPos: ConfigScannerPos{
						line: 3,
						col:  18,
					},
				},
				{
					tokenType: CtkComment,
					value:     "#EOF",
					startPos: ConfigScannerPos{
						line: 3,
						col:  19,
					},
					endPos: ConfigScannerPos{
						line: 3,
						col:  22,
					},
				},
				{
					tokenType: CtkEOF,
					startPos: ConfigScannerPos{
						line: 3,
						col:  22,
					},
					endPos: ConfigScannerPos{
						line: 3,
						col:  22,
					},
				},
			},
		},
	}

	for _, multiTokenTest := range multiTokenTests {
		scanner := NewConfigScanner(strings.NewReader(multiTokenTest.input))

		for _, expectedToken := range multiTokenTest.expectedTokens {
			token, err := scanner.Scan()

			if err != nil {
				t.Errorf("Scan failed with error %v", err)
			} else if !token.Equal(&expectedToken) {
				t.Errorf("ConfigToken does not match expected value. Expected %v, Actual %v", expectedToken, *token)
			}
		}
	}
}
