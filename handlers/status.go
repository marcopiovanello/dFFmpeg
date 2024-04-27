package handlers

import (
	"net/http"

	"github.com/marcopeocchi/sanji/scheduler"
)

func StatusHandler(s scheduler.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
