package liletestone

import (
	"fmt"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	//	"github.com/grpc-ecosystem/go-grpc-middleware/tags/logrus"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/karux/go-utils/security"
	"github.com/lileio/lile"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	oauth "google.golang.org/grpc/credentials/oauth"

	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	cm     = &sync.Mutex{}
	Client LiletestoneClient
)

func customClientCodeToLevel(c codes.Code) logrus.Level {
	if c == codes.Unauthenticated {
		// Make this a special case for tests, and an error.
		return logrus.ErrorLevel
	}
	level := grpc_logrus.DefaultClientCodeToLevel(c)
	return level
}

func GetLiletestoneClient() LiletestoneClient {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client
	}

	serviceURL := lile.URLForService("liletestone")

	//TODO: CONFIGURE CIRCUIT BREAKER

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{DisableTimestamp: false}
	logOpts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(customClientCodeToLevel),
	}

	// client connection timeout
	//TODO: retrieve from config, push down into circuit breaker?
	duration, _ := time.ParseDuration("3s")

	// CONFIGURE AUTHORIZATION TOKEN INTERCEPTOR
	//TODO: do we need NewOauthAccess??
	cred := oauth.NewOauthAccess(security.FetchClientToken(security.UserCredentials{}))
	fmt.Println("creds", cred)
	tlsCreds := security.GetClientTLSCreds()

	dOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(duration),
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(cred),
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				lile.ContextClientInterceptor(),
				otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer()),
				grpc_logrus.UnaryClientInterceptor(logrus.NewEntry(logger), logOpts...),
			),
		),
	}
	//dOpts = append(dOpts, grpc.WithPerRPCCredentials(cred))

	/*
	   grpc_ctxtags.Extract(ctx).Set("custom_tags.string", "something").Set("custom_tags.int", 1337)
	   // Extract a single request-scoped logrus.Logger and log messages.
	   l := ctx_logrus.Extract(ctx)

	*/

	// We don't need to error here, as this creates a pool and connections
	// will happen later
	conn, err := grpc.Dial(
		serviceURL,
		dOpts...,
	)

	if err != nil {
		log.Println("GetLiletestoneClient error ", err, "to serviceURL", serviceURL)
	} else {
		//log.Println(conn)
		//log.Println(serviceURL)
	}

	cli := NewLiletestoneClient(conn)
	Client = cli
	return cli
}
