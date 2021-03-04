package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

// An environment consists of name and a URL of the JSON data.
type environment struct {
	name string
	url  string
}

// CLI Args.
var (
	app       = kingpin.New("compare_core_versions", "A script that periodically compares the core_version of two Horizon environments.")
	frequency = app.Flag("frequency", "How often to compare versions (in seconds).").Short('t').Default("10").Int()
	name1     = app.Flag("name1", "Name of the first environment to compare").Required().String()
	url1      = app.Flag("url1", "URL of the first environment to compare").Required().String()
	name2     = app.Flag("name2", "Name of the second environment to compare").Required().String()
	url2      = app.Flag("url2", "URL of the second environment to compare").Required().String()
)

func main() {
	// Parse CLI flags.
	kingpin.MustParse(app.Parse(os.Args[1:]))

	env1 := environment{
		name: *name1,
		url:  *url1,
	}
	env2 := environment{
		name: *name2,
		url:  *url2,
	}

	for {
		compareCoreVersions(env1, env2)
		time.Sleep(time.Duration(*frequency) * time.Second)
	}
}

// Compares core versions and alerts if they are different.
func compareCoreVersions(env1, env2 environment) {

	// Download and parse JSON files.
	h1, h2 := getHorizonData(env1.url), getHorizonData(env2.url)

	// Extract SemVer and Hash Version from core_version string.
	h1SemVer, h2SemVer := h1.CoreSemVer(), h2.CoreSemVer()
	h1Hash, h2Hash := h1.CoreVersionHash(), h2.CoreVersionHash()

	semVerDiffers := h1SemVer != h2SemVer
	hashDiffers := h1Hash != h2Hash

	// Handle all cases of core_version differing.
	if semVerDiffers || hashDiffers {
		alert := fmt.Sprintf("[ALERT] core_version is different on %s and %s. %s (%s) != %s (%s).", env1.name, env2.name, h1SemVer, h1Hash, h2SemVer, h2Hash)
		if !semVerDiffers && hashDiffers {
			alert += " Warning: version numbers are identical despite version hashes differing."
		} else if !hashDiffers && semVerDiffers {
			alert += " Warning: version hashes are identical despite version numbers differing."
		}
		log.Println(alert)
	} else {
		log.Println("[INFO] core_version is the same.")
	}
}

func getHorizonData(URL string) HorizonData {
	horizon, err := getJsonFile(URL)
	if err != nil {
		log.Fatal("[ERROR] ", err)
	}
	return horizon
}

// Downloads file from URL and parses it as JSON. Returns it as HorizonData.
func getJsonFile(URL string) (HorizonData, error) {

	response, err := http.Get(URL)
	if err != nil {
		return HorizonData{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return HorizonData{}, errors.New("Received status code: " + strconv.Itoa(response.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return HorizonData{}, err
	}

	horizon := HorizonData{}
	if err := json.Unmarshal(bodyBytes, &horizon); err != nil {
		return HorizonData{}, err
	}

	return horizon, nil
}
