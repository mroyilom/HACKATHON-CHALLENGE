None selected 


Skip to content
Using Gmail with screen readers
12 of 10,868
base project
Inbox

Innocent Ekwueme
Attachments
Fri 27 Mar, 16:45 (3 days ago)
to me


 3 attachment
  •  Scanned by Gmail
package main

import (
	"strconv"
)

func BaseConverter(word string, base int) (int64, error) {

	num, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
func dex(n int64, base int) (string, error) {
	return strconv.FormatInt(n, base), nil
}
