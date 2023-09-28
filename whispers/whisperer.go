package whispers

import (
	"net/http"
	"time"
)

var (
    WHISPERS_TAXES_LIMIT    int             = 10
    WHISPERS_TIMEOUT        time.Duration   = 30 * time.Second
    WHISPERS_SERVER_DELAY   time.Duration   = 2 * time.Second
    WHISPERS_USER_AGENT     string          = "Mozilla/5.0 (Android 4.4; Mobile; rv:41.0) Geko/41.0 Firefox/41.0"
)

type Whisperer interface {
    WhisperJSON(dst string, msg interface{}) (*http.Response, error)
    WhisperGet(dst string) (*http.Response, error)
}

func NewWhisperer(proto string) Whisperer{
    switch proto {
    case "http":
        return NewHTTPWhisperer()
    default:
        return nil
    }
}
