package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ReadConfTest struct {
	name          string
	fileName      string
	expectedToken string
	shouldError   bool
}

func TestReadTokenFromFile(t *testing.T) {
	testCases := []ReadConfTest{
		{
			name:          "Successful Read",
			fileName:      "testConf.yml",
			expectedToken: "ThisIsATest",
			shouldError:   false,
		},
		{
			name:          "File Not Found",
			fileName:      "IdontExist.yankee",
			expectedToken: "This should fail",
			shouldError:   true,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual, err := ReadTokenFromFile(test.fileName)
			if test.shouldError {
				assert.Nil(t, actual)
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, test.expectedToken, actual.AuthToken)
			}
		})
	}

}
