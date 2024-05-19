package interact

import (
	"fmt"
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
	"gmap/internal/interact/animate"
	"time"
)

const appName = "hpzm"
const appDescription = "a High Precision Zone Mapper"
const appVersion = "v0.0.1"

var (
	grumbleApp = grumble.New(&grumble.Config{
		Name:                  appName,
		Description:           appDescription,
		PromptColor:           color.New(color.FgGreen, color.Bold),
		HelpHeadlineColor:     color.New(color.FgGreen),
		HelpHeadlineUnderline: true,
		HelpSubCommands:       true,
	})
)

func init() {
	animate.ShowHiddenLogo()
	time.Sleep(200 * time.Millisecond)
	fmt.Println(color.WhiteString("\nWelcome to:\n"))
	grumbleApp.SetPrintASCIILogo(func(a *grumble.App) {
		/*
			 __
			/\ \
			\ \ \___   _____   ____     ___ ___
			 \ \  _ `\/\ '__`\/\_ ,`\ /' __` __`\
			  \ \ \ \ \ \ \L\ \/_/  /_/\ \/\ \/\ \
			   \ \_\ \_\ \ ,__/ /\____\ \_\ \_\ \_\
			    \/_/\/_/\ \ \/  \/____/\/_/\/_/\/_/
			             \ \_\
			              \/_/
		*/
		a.Println(" __")
		a.Println("/\\ \\")
		a.Println("\\ \\ \\___   _____   ____     ___ ___")
		a.Println(" \\ \\  _ `\\/\\ '__`\\/\\_ ,`\\ /' __` __`\\")
		a.Println("  \\ \\ \\ \\ \\ \\ \\L\\ \\/_/  /_/\\ \\/\\ \\/\\ \\")
		a.Println("   \\ \\_\\ \\_\\ \\ ,__/ /\\____\\ \\_\\ \\_\\ \\_\\")
		a.Println("    \\/_/\\/_/\\ \\ \\/  \\/____/\\/_/\\/_/\\/_/")
		a.Println("             \\ \\_\\")
		a.Println("              \\/_/")

		a.Println("\n")
		a.Printf("%s %s - %s\n", color.GreenString(a.Config().Name), color.CyanString(appVersion), color.YellowString(a.Config().Description))
		a.Println("\ntype 'help' for a list of commands\n")

	})
}

func Run() {
	grumbleApp.Run()
}
