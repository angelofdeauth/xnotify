// @File:     inf.go
// @Created:  2020-03-19 20:41:05
// @Modified: 2020-03-29 00:52:18
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package inf sets up various informational fields.
package inf

import (
	"time"

	"github.com/angelofdeauth/xnotify/pkg/rtc"
	"github.com/urfave/cli/v2"
)

// Set sets the information fields of the cli.App.
func Set(a *cli.App, rtc *rtc.RunTimeCfg) {
	a.Name = rtc.App
	a.Version = rtc.Version
	a.Compiled, _ = time.Parse(time.UnixDate, rtc.Time)
	a.Copyright = "Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>"
	a.Description = "xnotify is a filesystem event based workflow automation daemon.\n"
	a.Usage = "A filesystem event based workflow automation daemon."
	a.Metadata = map[string]interface{}{
		"Commit": rtc.Commit,
	}
}
