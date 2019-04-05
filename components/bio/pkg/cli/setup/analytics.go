package setup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/biome-sh/biome-go/components/bio/pkg/ui"

	homedir "github.com/mitchellh/go-homedir"
)

func setupAnalytics() {
	analyticsHelpText()
	if ui.PromptYesNo("Enable analytics?") {
		createTouchFile(true)
	} else {
		createTouchFile(false)
		fmt.Println("  Okay, maybe another time.")
	}
}

func analyticsHelpText() {
	fmt.Println()
	ui.Heading("Analytics")
	fmt.Println(`
  The "bio" command-line tool will optionally send anonymous usage data to 
  Biome's Google Analytics account. This is a strictly opt-in activity 
  and no tracking will occur unless you respond affirmatively to the 
  question below.`)
	fmt.Println(`
  We collect this data to help improve Biome's user experience. For 
  example, we would like to know the category of tasks users are 
  performing, and which ones they are having trouble with (e.g. mistyping 
  command line arguments).`)
	fmt.Println(`
  To see what kinds of data are sent and how they are anonymized, please 
  read more about our analytics here: 
  https://www.habitat.sh/docs/about-analytics/`)
	fmt.Println()
}

func createTouchFile(optIn bool) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	basePath := filepath.Join(home, ".bio", "cache", "analytics")
	var newFile *os.File
	optOutFile := filepath.Join(basePath, "OPTED_OUT")
	optInFile := filepath.Join(basePath, "OPTED_IN")
	if optIn {
		ui.Begin("Opting in to analytics")
		ui.Status(ui.Deleting, optOutFile)
		os.Remove(optOutFile)
		ui.Status(ui.Creating, optInFile)
		newFile, err = os.Create(optInFile)
		ui.End("Analytics opted in, thank you!")
	} else {
		ui.Begin("Opting out of analytics")
		ui.Status(ui.Deleting, optInFile)
		os.Remove(optInFile)
		ui.Status(ui.Creating, optOutFile)
		newFile, err = os.Create(optOutFile)
		ui.End("Analytics opted out, we salute you just the same!")
	}
	if err != nil {
		fmt.Println("Error:", err)
	}
	newFile.Close()
}
