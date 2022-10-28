package response

import (
	"strings"
)

func IsSQLError(e error) bool {
	if strings.Contains(e.Error(), "SQLSTATE") {
		return true
	}
	return false
}

func GetSQLErrorCode(e error) string {
	str := e.Error()
	i := strings.Index(str, "(")
	j := strings.Index(str, ")")
	// start := str[i+1:]
	strSqlErr := str[i+1 : j]
	words := strings.Fields(strSqlErr)
	return words[1]
}
