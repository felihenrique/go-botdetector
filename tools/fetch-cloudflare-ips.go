package tools

import (
	"strings"

	httputils "github.com/felihenrique/go-botdetector/pkg/http-utils"
)

const CLOUDFLARE_IPS_URL = "https://www.cloudflare.com/ips-v4"

func FetchCloudflareIps() ([]string, error) {
	data, err := httputils.GetData(CLOUDFLARE_IPS_URL)

	if err != nil {
		return nil, err
	}

	ips := strings.Split(string(data), "\n")

	return ips, nil
}
