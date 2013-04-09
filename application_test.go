// Copyright (C) 2013 Jason McVetta, all rights reserved.

package stormpath

import (
	"github.com/darkhelmet/env"
	"testing"
)

func setupStormpath(t *testing.T) *Stormpath {
	apiId := env.String("STORMPATH_API_ID")
	apiSecret := env.String("STORMPATH_API_SECRET")
}
