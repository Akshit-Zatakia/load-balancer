package lb

import (
	"context"
	"time"

	serverpool "github.com/Akshit-Zatakia/load-balancer/server-pool"
	"github.com/Akshit-Zatakia/load-balancer/utils"
)

func LauchHealthCheck(ctx context.Context, sp serverpool.ServerPool) {
	t := time.NewTicker(time.Second * 20)
	utils.Logger.Info("Starting health check...")
	for {
		select {
		case <-t.C:
			go serverpool.HealthCheck(ctx, sp)
		case <-ctx.Done():
			utils.Logger.Info("Closing Health Check")
			return
		}
	}
}
