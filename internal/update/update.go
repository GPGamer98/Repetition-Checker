package update

import (
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const Version = "0.0.3"

func DoSelfUpdate() bool {
	v := semver.MustParse(Version)

	latest, err := selfupdate.UpdateSelf(v, "GPGamer98/Repetition-Checker")
	log.Println(latest)
	if err != nil {
		// log.Println("Update failed:", err)
		return false
	}
	if latest.Version.Equals(v) {
		// log.Println("Already updated")
		return true
	} else {
		// log.Println("Updated to version", latest.Version)
		// log.Println("Release notes:", latest.ReleaseNotes)
		return true
	}
}

func CheckUpdate() (bool, string, string) {
	latest, found, err := selfupdate.DetectLatest("GPGamer98/Repetition-Checker")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return false, "", ""
	}

	v := semver.MustParse(Version)
	if !found || latest.Version.LTE(v) {
		log.Println(found)
		log.Println(latest.Version)
		// log.Println("Current version is the latest")
		return false, "", v.String()
	}

	return true, latest.Version.String(), v.String()
}