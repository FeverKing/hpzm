package interact

import (
	"fmt"
	"github.com/desertbit/grumble"
	"gmap/internal/filemanager"
	"gmap/internal/log"
)

func init() {

	listCmd := &grumble.Command{
		Name: "listfp",
		Help: "list all fingerprints in the database",
		// Aliases: []string{"run"},

		Flags: func(f *grumble.Flags) {
			f.String("H", "hpzm_db", "./os-info", "the database to store hpzm fingerprint")
		},
		Run: func(c *grumble.Context) error {
			fps, err := filemanager.GetFpFromOsInfo(c.Flags.String("hpzm_db"))
			if err != nil {
				log.Errorf("failed to get fingerprints from the database: %v", err)
				return nil
			}
			fmt.Printf("fingerprints of OS/device in the database:\n")
			for _, fp := range fps {
				fmt.Println(fp)
			}
			fmt.Println("total:", len(fps))
			return nil
		},
	}

	// grumbleApp.AddCommand(listCmd)

	listNCmd := &grumble.Command{
		Name: "n",
		Help: "list all fingerprints in nmap database",

		Flags: func(f *grumble.Flags) {
			f.String("N", "nmap_db", "", "the database to store the all fingerprint")
		},
		Run: func(c *grumble.Context) error {
			nmapDb := c.Flags.String("nmap_db")
			if nmapDb == "" {
				var err error
				nmapDb, err = filemanager.GetNmapOsDbPath()
				if err != nil {
					return err
				}
			} else {
				nmapDb = nmapDb
			}
			fps, err := filemanager.GetFpFromOsInfo(nmapDb)
			if err != nil {
				log.Errorf("failed to get fingerprints from the database: %v", err)
			}
			fmt.Printf("fingerprints of OS/device in nmap database:\n")
			for _, fp := range fps {
				fmt.Println(fp)
			}
			fmt.Println("total:", len(fps))
			return nil
		},
	}
	listCmd.AddCommand(listNCmd)
	grumbleApp.AddCommand(listCmd)
}
