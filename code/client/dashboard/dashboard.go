package dashboard

import (
	"fmt"
	"net/http"

	"github.com/jkstack/natpass/code/client/global"
	"github.com/jkstack/natpass/code/client/pool"
	"github.com/jkstack/natpass/code/client/rule"
)

// Dashboard dashboard object
type Dashboard struct {
	cfg     *global.Configure
	pl      *pool.Pool
	mgr     *rule.Mgr
	Version string
}

// New create dashboard object
func New(cfg *global.Configure, pl *pool.Pool, mgr *rule.Mgr, version string) *Dashboard {
	return &Dashboard{
		cfg:     cfg,
		pl:      pl,
		mgr:     mgr,
		Version: version,
	}
}

// ListenAndServe listen and serve http handler
func (db *Dashboard) ListenAndServe(addr string, port uint16) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/info", db.Info)
	mux.HandleFunc("/api/rules", db.Rules)
	mux.HandleFunc("/", db.Render)
	svr := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", addr, port),
		Handler: mux,
	}
	return svr.ListenAndServe()
}
