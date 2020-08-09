package handlers

import (
	"go-projects/choose-your-own-adventure/configs"
	"net/http"
)

// Ping - A Handler for checking the heartbeat/health
func Ping(w http.ResponseWriter, req *http.Request, ctx configs.AppContext) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	check := configs.Healthcheck{
		AppName: "Choose-Your-Own-Adventure",
	}
	ctx.Render.JSON(w, http.StatusOK, check)
}
