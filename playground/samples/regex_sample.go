package samples

import (
	"fmt"
	"regexp"
	"strings"
)

func RegexSampleString() {
	containse := strings.Contains("愛はあるんか", "あい")
	fmt.Println(containse)

	match, _ := regexp.MatchString("あるんか", "愛はあるんか")
	fmt.Println(match)

	r := regexp.MustCompile("[a-zA-Z0-9]")
	ms := r.MatchString("愛はあるんか")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(愛|は|あるんか)")
	r2r := r2.FindString("/愛")
	fmt.Println(r2r)

	r3r := r2.FindStringSubmatch("/愛")
	fmt.Println(r3r)
}
