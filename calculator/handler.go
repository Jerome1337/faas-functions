package function

import (
	"fmt"
	"strconv"
)

func Handle(req []byte) string {
	num, _ := strconv.Atoi(string(req))

	return fmt.Sprintf("The half of: %d is %d", num, num/2)
}
