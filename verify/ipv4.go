package verify

import (
	"strconv"

	"github.com/sjlleo/netflix-verify/util"
)

type IPv4Verifier struct {
	Config
	IP              string
	unblockStatus   int
	unblockTestChan chan UnblockTestResult
}

func (v *IPv4Verifier) Execute() *VerifyResponse {
	var err error
	var response VerifyResponse
	v.unblockStatus = AreaUnavailable
	response.Type = 1

	if v.IP, err = util.DnsResolver(4); err != nil {
		response.StatusCode = NetworkUnrachable
		return &response
	}

	v.unblockTestChan = make(chan UnblockTestResult)
	defer close(v.unblockTestChan)

		res := <-v.unblockTestChan
		if res.CountryCode != "" {
			v.unblockStatus = CustomMovieUnblock
			response.CountryCode = res.CountryCode
			response.CountryName = util.CountryCodeToCountryName(res.CountryCode)
		} else {
			v.unblockStatus = CustomMovieBlock
		}
	}
	return &response
}

func (v *IPv4Verifier) upgradeStatus(status int) {
	if v.unblockStatus < status {
		v.unblockStatus = status
	}
}
