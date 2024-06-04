package interact

import (
	"github.com/desertbit/grumble"
	"gmap/internal/filemanager"
	"gmap/internal/log"
)

func init() {
	grumbleApp.AddCommand(&grumble.Command{
		Name: "merge",
		Help: "Merge custom fingerprints to the database",
		Flags: func(f *grumble.Flags) {
			f.String("n", "main_db", "", "file to merge")
			f.String("g", "hpzm_db", "os-info", "output file")
		},
		Run: func(c *grumble.Context) error {
			nmapDb := c.Flags.String("main_db")
			gmapDb := c.Flags.String("hpzm_db")
			if nmapDb == "" {
				var err error
				nmapDb, err = filemanager.GetNmapOsDbPath()
				if err != nil {
					return err
				}
			} else {
				nmapDb = nmapDb
			}
			// log.Infof("merging %s to %s", gmapDb, nmapLoc)
			err := filemanager.MergeFiles(nmapDb, gmapDb)
			if err != nil {
				return err
			}
			log.Infof("merge successful")
			return nil
		},
	})
}
