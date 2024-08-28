package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID          string `json:"id" validate:"len:36"`
		Name        string
		Age         int      `validate:"min:18|max:50"`
		Email       string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role        UserRole `validate:"in:admin,stuff"`
		Phones      []string `validate:"len:11"`
		meta        json.RawMessage
		Application App `validate:"nested"`
	}

	App struct {
		Version   string `validate:"len:5"`
		UserToken Token  `validate:"nested"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
		Response  Response `validate:"nested"`
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "123456789012345678901234567890123456",
				Name:   "qwert",
				Age:    12,
				Email:  "qwert",
				Role:   "qwert",
				Phones: []string{"123456", "12345678999"},
				meta:   []byte{},
				Application: App{
					Version: "12345678",
					UserToken: Token{
						Header:    []byte{},
						Payload:   []byte{},
						Signature: []byte{},
						Response: Response{
							Code: 100,
						},
					},
				},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Age",
					Err:   ErrNumberMin,
				},
				ValidationError{
					Field: "Phones",
					Err:   ErrStringLength,
				},
				ValidationError{
					Field: "Email",
					Err:   ErrStringRegexp,
				},
				ValidationError{
					Field: "Role",
					Err:   ErrStringIn,
				},
				ValidationError{
					Field: "Version",
					Err:   ErrStringLength,
				},
				ValidationError{
					Field: "Code",
					Err:   ErrNumberIn,
				},
			},
		},
		// ...
		// Place your code here.
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			// Place your code here.
			err := Validate(tt.in)
			var valErr ValidationErrors
			require.ErrorAs(t, err, &valErr)

			for _, expectEr := range strings.Split(tt.expectedErr.Error(), "\n") {
				require.ErrorContains(t, err, expectEr)
			}
		})
	}
}
