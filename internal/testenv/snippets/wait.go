package snippets

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/pomerium/pomerium/internal/testenv"
	"github.com/pomerium/pomerium/pkg/grpcutil"
	"github.com/pomerium/pomerium/pkg/telemetry/trace"
)

func WaitStartupComplete(env testenv.Environment, timeout ...time.Duration) time.Duration {
	if env.GetState() == testenv.NotRunning {
		panic("test bug: WaitStartupComplete called before starting the test environment")
	}
	_, span := trace.Continue(env.Context(), "snippets.WaitStartupComplete")
	defer span.End()
	start := time.Now()
	recorder := env.NewLogRecorder()
	if len(timeout) == 0 {
		timeout = append(timeout, 1*time.Minute)
	}
	ctx, ca := context.WithTimeout(env.Context(), timeout[0])
	defer ca()
	recorder.WaitForMatch(map[string]any{
		"syncer-id":   "databroker",
		"syncer-type": "type.googleapis.com/pomerium.config.Config",
		"message":     "listening for updates",
	}, timeout...)
	cc, err := grpc.Dial(env.DatabrokerURL().Value(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(grpcutil.WithUnarySignedJWT(env.SharedSecret)),
		grpc.WithChainStreamInterceptor(grpcutil.WithStreamSignedJWT(env.SharedSecret)),
	)
	env.Require().NoError(err)
	env.Require().True(cc.WaitForStateChange(ctx, connectivity.Ready))
	cc.Close()
	return time.Since(start)
}
