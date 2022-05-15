package utils

import (
	"fmt"

	"github.com/tai9/cargo-nft-be/constants"
)

func ParseApiUrl(url string) string {
	return fmt.Sprintf("%v%v", constants.BASE_URL, url)
}
