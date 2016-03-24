// Package go-edit provides the algorithms needed to compute the minimum edit distance and
// weigthed edit distance between two strings.
package go_edit

import (
    "math"
)

// MinEditDistance will return the minimum edit distance between two strings i, j
func MinEditDistance(str1, str2 string) int {
    // Have to initalize things like C...
    var d = make([][]int, len(str1))

    for x := range d {
        d[x] = make([]int, len(str2))
    }

    for i := 0; i < len(str1); i++ {
        for j := 0; j < len(str2); j++ {
            // Check all edge cases to calculate the next distance
            if i == 0 && j == 0 {
                if str1[i] != str2[j] {
                    d[i][j] = 1
                } else {
                    d[i][j] = 0
                }
            } else if i == 0 && j > 0 {
                d[i][j] = d[i][j-1] + 1
            } else if i > 0 && j == 0 {
                d[i][j] = d[i-1][j] + 1
            } else if i > 0 && j > 0 {
                z := 0
                if str1[i] != str2[j] {
                    z = 2
                }
                d[i][j] = min([]int{d[i-1][j] + 1, d[i][j-1] + 1, d[i-1][j-1] + z})
            }
        }
    }

    return d[len(str1) - 1][len(str2) - 1]
}

// Internal min function that will find the minimum value of an array of type int
func min(arr []int) int {
    min := int(math.MaxInt64)
    for _, val := range arr {
        if val < min {
            min = val
        }
    }

    return min
}
