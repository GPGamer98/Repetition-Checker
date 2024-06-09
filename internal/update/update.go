package update

import (
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const Version = "0.0.2"

func DoSelfUpdate() (bool, string) {
	v := semver.MustParse(Version)

	latest, err := selfupdate.UpdateSelf(v, '')
}