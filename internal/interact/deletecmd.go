package interact

import (
	"errors"
	"fmt"
	"github.com/desertbit/grumble"
	"gmap/internal/filemanager"
	"gmap/internal/log"
)

func init() {
	grumbleApp.AddCommand(&grumble.Command{
		Name: "deletefp",
		Help: "delete a fingerprint from the database",

		Flags: func(f *grumble.Flags) {
			f.String("H", "hpzm_db", "./os-info", "the database to store the fingerprint")
			f.String("v", "vendor", "", "the vendor of the target")
			f.String("m", "model", "", "the model of the target")
		},
		Run: func(c *grumble.Context) error {
			if c.Flags.String("vendor") == "" || c.Flags.String("model") == "" {
				return errors.New("missing required flags, type 'help deletefp' for more information")
			}
			modelStr := fmt.Sprintf("%s %s", c.Flags.String("vendor"), c.Flags.String("model"))
			hasFp := filemanager.HasFpInOsInfo(modelStr, c.Flags.String("hpzm_db"))
			if !hasFp {
				log.Errorf("vendor: %s model: %s not found in the database", c.Flags.String("vendor"), c.Flags.String("model"))
				return nil
			}
			err := filemanager.DeleteFpFromOsInfo(modelStr, c.Flags.String("hpzm_db"))
			if err != nil {
				log.Errorf("failed to delete fingerprint from the database: %v", err)
				return nil
			}
			log.Infof("vendor: %s model: %s deleted from the database", c.Flags.String("vendor"), c.Flags.String("model"))
			return nil
		},
	})
}
