package go_edit

import (
    "testing"
)

type stringpair struct {
    a, b string
    result int
}

var tests = []stringpair{
    stringpair{"blah", "blah", 0},
    stringpair{"Blah", "BLAH", 6},
    stringpair{"intention", "iitention", 2},
    stringpair{"intention", "iitenaion", 4},
    stringpair{"execution", "intention", 7},
}

func TestMinEditDistance(t *testing.T) {
    for _, pair := range tests {
        val := MinEditDistance(pair.a, pair.b)
        if val != pair.result {
            t.Error(
                "For", pair.a, pair.b,
                "expected", pair.result,
                "got", val,
            )
        }
    }
}
