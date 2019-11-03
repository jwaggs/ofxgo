package main

import (
	"flag"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
)

type command struct {
	Name        string
	Description string
	Flags       *flag.FlagSet
	CheckFlags  func() bool // Check the flag values after they're parsed, printing errors and returning false if they're incorrect
	Do          func()      // Run the command (only called if CheckFlags returns true)
}

func (c *command) usage() {
	fmt.Printf("Usage of %s:\n", c.Name)
	c.Flags.PrintDefaults()
}

// flags common to all server transactions
var serverURL, username, password, org, fid, appID, appVer, ofxVersion, clientUID string
var noIndentRequests bool

func defineServerFlags(f *flag.FlagSet) {
	f.StringVar(&serverURL, "url", "https://ofx.chase.com", "Financial institution's OFX Server URL (see ofxhome.com if you don't know it)")
	f.StringVar(&clientUID, "clientuid", os.Getenv("CHASE_CLIENTUID"), "Client UID (only required by a few FIs, like Chase)")
	f.StringVar(&username, "username", os.Getenv("CHASE_USERNAME"), "Your username at financial institution")
	f.StringVar(&password, "password", "", "Your password at financial institution")
	f.StringVar(&org, "org", "B1", "'ORG' for your financial institution")
	f.StringVar(&fid, "fid", "10898", "'FID' for your financial institution")
	f.StringVar(&appID, "appid", "QWIN", "'APPID' to pretend to be")
	f.StringVar(&appVer, "appver", "2700", "'APPVER' to pretend to be")
	f.StringVar(&ofxVersion, "ofxversion", "220", "OFX version to use")
	f.BoolVar(&noIndentRequests, "noindent", false, "Don't indent OFX requests")
}

func checkServerFlags() bool {
	var ret bool = true
	if len(serverURL) == 0 {
		fmt.Println("Error: Server URL empty")
		ret = false
	}
	if len(username) == 0 {
		fmt.Println("Error: Username empty")
		ret = false
	}

	if ret && len(password) == 0 {
		fmt.Printf("Password for %s: ", username)
		pass, err := gopass.GetPasswd()
		if err != nil {
			fmt.Printf("Error reading password: %s\n", err)
			ret = false
		} else {
			password = string(pass)
		}
	}
	return ret
}
