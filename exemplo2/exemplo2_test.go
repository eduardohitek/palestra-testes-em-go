package exemplo2

import "testing"

func TestIsOdd(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		exp      bool
		isErrExp bool
		errExp   error
	}{
		{name: "when odd and positive", input: 2, exp: true},
		{name: "when even and positive", input: 3, exp: false},
		{name: "when input is zero", input: 0, exp: false, isErrExp: true, errExp: ErrInputIsZero},
		{name: "when negative", input: -2, isErrExp: true, errExp: ErrInputCantBeNegative},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := isOdd(tc.input)
			if tc.isErrExp && err == nil {
				t.Errorf("expected to return an error, got nil")
			}
			if tc.isErrExp && err != tc.errExp {
				t.Errorf("expected the errr to be %q, got %q", tc.errExp, err)
			}
			if result != tc.exp {
				t.Errorf("expected to be %t, but found %t", tc.exp, result)
			}
		})
	}

}
