package vsid

import (
	"strings"
	"github.com/pborman/uuid"
	"fmt"
)

func VsId() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}
