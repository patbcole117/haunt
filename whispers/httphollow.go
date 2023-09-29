package whispers

import (
    "encoding/json"
    "fmt"
    "io"
	"net/http"
    "time"
)

type HTTPHollow struct {
    Addy    string
    Srv     *http.Server
    Taxes   chan Curse
}
func NewHTTPHollow(a string) *HTTPHollow {
    h := HTTPHollow{Addy: a}
    h.Srv = h.GetServer()
    h.Taxes = make(chan Curse, WHISPERS_TAXES_LIMIT)
    return &h
}

func (h *HTTPHollow) GetServer() *http.Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/", root(h))
    mux.HandleFunc("/taxes/", taxes(h))
    return &http.Server{
            Addr: h.Addy,
            Handler: mux,
            ReadTimeout: WHISPERS_TIMEOUT,
            WriteTimeout: WHISPERS_TIMEOUT,     
    }
}

func (h *HTTPHollow) Listen() {
    h.Srv.ListenAndServe()
    time.Sleep(WHISPERS_SERVER_DELAY)
}

func root(h *HTTPHollow) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("[>]", time.Now().Format(time.RFC1123Z), r.RemoteAddr, r.Method, r.Host + r.URL.Path)
        req, err := io.ReadAll(r.Body)
        if err != nil {
            fmt.Println("[!]", err.Error())
            return
        }
        fmt.Println("[>]", string(req))

        c := NewCurse(CURSE_TYPE_HELLO)
        if len(h.Taxes) > 0 {
            c = <- h.Taxes
        }

        resp, err := json.Marshal(c)
        if err != nil {
            fmt.Println("[!]", err.Error())
            return
        }
        fmt.Println("[<]", string(resp))
        w.Write(resp)
    }
}

func taxes(h *HTTPHollow) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("[>]", time.Now().Format(time.RFC1123Z), r.RemoteAddr, r.Method, r.Host + r.URL.Path)
        req, err := io.ReadAll(r.Body)
        if err != nil {
            fmt.Println("[!]", err.Error())
            return
        }
        fmt.Println("[>]", string(req))

        var c Curse
        if err := json.Unmarshal(req, &c); err != nil {
            fmt.Println("[!]", err.Error())
            return
        }
        h.Taxes <- c

        c = NewCurse(CURSE_TYPE_TAXES_ADDED)
        rep, err := json.Marshal(c)
        if err != nil {
            fmt.Println("[!]", err.Error())
            return
        }
        fmt.Println("[<]", string(rep))
        w.Write(rep)
    }
}