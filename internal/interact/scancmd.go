package interact

import (
	"errors"
	"fmt"
	"github.com/desertbit/grumble"
	"gmap/internal/filemanager"
	"gmap/internal/log"
	"gmap/internal/scan"
)

func init() {
	scanCmd := &grumble.Command{
		Name: "addfp",
		Help: "scan a known target and add its fingerprint to the database",
		// Aliases: []string{"run"},

		Flags: func(f *grumble.Flags) {
			f.Int("t", "times", 0, "number of times to run the scan")
			f.String("T", "target", "", "the target to scan")
			f.String("v", "vendor", "", "the vendor of the target")
			f.String("m", "model", "", "the model of the target")
			f.String("F", "database", "./os-info", "the database to store the fingerprint")
		},

		Run: func(c *grumble.Context) error {
			if c.Flags.String("target") == "" || c.Flags.Int("times") == 0 || c.Flags.String("vendor") == "" || c.Flags.String("model") == "" {
				return errors.New("missing required flags, type 'help addfp' for more information")
			}
			fullName := fmt.Sprintf("%s %s", c.Flags.String("vendor"), c.Flags.String("model"))
			hasFp := filemanager.HasFpInOsInfo(fullName, c.Flags.String("database"))
			if hasFp {
				log.Errorf("vendor: %s model: %s already exists in the database", c.Flags.String("vendor"), c.Flags.String("model"))
				return nil
			}
			fpStr := scan.GenFinalFingerprint(c.Flags.String("target"), c.Flags.Int("times"), c.Flags.String("vendor"), c.Flags.String("model"))
			err := filemanager.AddFpToOsInfo(fpStr, c.Flags.String("database"))
			if err != nil {
				log.Errorf("failed to add fingerprint to the database: %v", err)
			}
			log.Infof("target: %s with vendor model: %s %s added to the database", c.Flags.String("target"), c.Flags.String("vendor"), c.Flags.String("model"))
			return nil
		},
	}
	grumbleApp.AddCommand(scanCmd)

	detectCmd := &grumble.Command{
		Name: "detect",
		Help: "detect the alive hosts",
		Flags: func(f *grumble.Flags) {
			f.String("T", "target", "", "the target to scan")
		},
		Run: func(c *grumble.Context) error {
			target := c.Flags.String("target")
			aliveHosts, err := scan.Detect(target)
			if err != nil {
				log.Errorf("failed to detect alive hosts: %v", err)
				return nil
			}
			fmt.Printf("alive hosts:\n")
			for _, host := range aliveHosts {
				fmt.Println(host)
			}
			fmt.Println("total:", len(aliveHosts))
			return nil
		},
	}

	detectOsCmd := &grumble.Command{
		Name: "os",
		Help: "detect the os of the target",
		Flags: func(f *grumble.Flags) {
			f.String("T", "target", "", "the target to scan")
		},
		Run: func(c *grumble.Context) error {
			target := c.Flags.String("target")
			osInfo, err := scan.DetectOs(target)
			if err != nil {
				log.Errorf("failed to detect os of the target: %v", err)
			}
			//fmt.Println("os of the target:")
			//for host, os := range osInfo {
			//	fmt.Println("host:", host, "os:", os)
			//}
			fmt.Println("total:", len(osInfo))
			return nil
		},
	}

	detectCmd.AddCommand(detectOsCmd)
	grumbleApp.AddCommand(detectCmd)
}
