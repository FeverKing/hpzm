package interact

import "github.com/desertbit/grumble"

func init() {
	versionCmd := &grumble.Command{
		Name: "version",
		Help: "print the version of the application",
		Run: func(c *grumble.Context) error {
			c.App.Printf("%s-%s\n", appName, appVersion)
			return nil
		},
	}
	grumbleApp.AddCommand(versionCmd)
}
