package logs

import (
	"log"

	"github.com/getsentry/raven-go"

	"github.com/rifferreinert/bikescrape"
)

func Fatal(msg error) {
	raven.CaptureErrorAndWait(msg, bikescrape.RavenContext)
	log.Fatal(msg)
}
