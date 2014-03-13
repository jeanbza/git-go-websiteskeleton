package common

import (
    "bytes"
)

func StrCat(str1 string, str2 string) (string) {
    var catString bytes.Buffer

    catString.WriteString(str1)
    catString.WriteString(str2)

    return catString.String()
}