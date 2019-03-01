package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/obiknows/social_media_example_site/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
