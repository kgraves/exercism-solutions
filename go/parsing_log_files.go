package parsinglogfiles

import "fmt"
import "regexp"
import "strings"

var validLogPrefixesRegex *regexp.Regexp = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
var logSeparatorRegex *regexp.Regexp = regexp.MustCompile(`<[~*=-]*>`)
var passwordRegex *regexp.Regexp = regexp.MustCompile(`".*password.*"`)
var lineSuffixRegex *regexp.Regexp = regexp.MustCompile(`end-of-line\d+`)
var usernameRegex *regexp.Regexp = regexp.MustCompile(`User +(\w+)`)

func IsValidLine(text string) bool {
	return validLogPrefixesRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
    return logSeparatorRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0

    for _, line := range lines {
        if passwordRegex.MatchString(strings.ToLower(line)) {
            count++
        }
    }

    return count
}

func RemoveEndOfLineText(text string) string {
	if lineSuffixRegex.MatchString(text) {
        text = lineSuffixRegex.ReplaceAllString(text, "")
    }

    return text
}

func TagWithUserName(lines []string) []string {
    for i, line := range lines {
        res := usernameRegex.FindStringSubmatch(line)

        if len(res) > 0 {
            lines[i] = fmt.Sprintf("[USR] %s %s", res[1], line)
        }
    }

    return lines
}
