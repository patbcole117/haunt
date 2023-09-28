package whispers

import (
    "bytes"
    "encoding/json"
	"net/http"
    "time"
)

type HTTPWhisperer struct {
	c *http.Client
}

func NewHTTPWhisperer() *HTTPWhisperer {
    return &HTTPWhisperer{
        c: &http.Client{Timeout: WHISPERS_TIMEOUT},
    }
}

func (w *HTTPWhisperer) WhisperJSON(dst string, msg interface{}) (*http.Response, error) {
    b, err := json.Marshal(msg)
    if err != nil {
        return nil, err
    } 
    body := bytes.NewReader(b)

    req, err := http.NewRequest(http.MethodPost, dst, body)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", WHISPERS_USER_AGENT)
    req.Header.Set("Date", time.Now().Format(time.RFC1123))
    
    res, err := w.c.Do(req)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (w *HTTPWhisperer) WhisperGet(dst string) (*http.Response, error) {
    req, err := http.NewRequest(http.MethodGet, dst, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("User-Agent", WHISPERS_USER_AGENT)
    req.Header.Set("Date", time.Now().Format(time.RFC1123))
    
    res, err := w.c.Do(req)
    if err != nil {
        return nil, err
    }
    return res, nil
}
