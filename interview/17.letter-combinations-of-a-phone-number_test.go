package interview

import (
    "testing"
)

func TestLetterCombinations(t *testing.T) {
    tests := []struct {
        name   string
        digits string
        want   []string
    }{
        {
            name:   "标准23",
            digits: "23",
            want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
        },
        {
            name:   "空字符串",
            digits: "",
            want:   nil,
        },
        {
            name:   "单个数字",
            digits: "2",
            want:   []string{"a", "b", "c"},
        },
        {
            name:   "三个数字",
            digits: "234",
            want: []string{
                "adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
                "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
                "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
            },
        },
        {
            name:   "含7和9的4个字母",
            digits: "79",
            want: []string{"pw", "px", "py", "pz", "qw", "qx", "qy", "qz", "rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := letterCombinations(tt.digits)
            if !strEq(got, tt.want) {
                t.Errorf("letterCombinations(%q) = %v, want %v", tt.digits, got, tt.want)
            }
        })
    }
}

func strSliceEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    m := make(map[string]bool)
    for _, v := range a {
        m[v] = true
    }
    for _, v := range b {
        if !m[v] {
            return false
        }
    }
    return true
}
