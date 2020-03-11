package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"AY1st/infra"
	"AY1st/model"
	"AY1st/registry"
	"AY1st/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

const (
	ipEnv              = "IP"
	portEnv            = "PORT"
	cacheMaxAgeEnv     = "CACHE_MAX_AGE"
	shutdownTimeoutEnv = "SHUTDOWN_TIMEOUT"
)

func loadServiceRegistrySettings() *registry.ServiceRegistrySettings {
	// service factoryの初期化
	serviceRegistrySettings := &registry.ServiceRegistrySettings{}
	return serviceRegistrySettings
}

// Start starts api server
func Start() error {
	logger := util.GetLogger()

	// ログの出力設定
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.DebugLevel
		logger.Warnf("LOG_LEVEL is not set.")
	}

	// db engine 初期化
	engine, err := infra.SetupDBEngine(logLevel)
	if err != nil {
		logger.Fatal(err)
	}
	defer func() {
		log.Println("engine closed")
		engine.Close()
	}()

	maxAgeStr, ok := os.LookupEnv(cacheMaxAgeEnv)
	if !ok {
		return fmt.Errorf("%v is not set", cacheMaxAgeEnv)
	}
	cacheMaxAge, err := strconv.ParseInt(maxAgeStr, 10, 64)
	if err != nil {
		return err
	}

	loggerAccess := logrus.New()
	loggerAccess.Level = logLevel
	loggerAccess.Out = os.Stdout
	loggerAccess.Formatter = &logrus.JSONFormatter{}

	loggerError := logrus.New()
	loggerError.Level = logLevel
	loggerError.Out = os.Stderr
	loggerError.Formatter = &logrus.JSONFormatter{}
	gin.DefaultErrorWriter = loggerError.Writer()

	serviceRegistrySettings := loadServiceRegistrySettings()
	serviceRegistrySettings.Engine = engine
	registry := registry.NewService(serviceRegistrySettings)

	binding.EnableDecoderUseNumber = true
	// override gin validator
	binding.Validator = &model.StructValidator{}

	showVersion := true
	logger.Infof("API Version %s", version)

	// Ginの初期化
	r := gin.Default()

	// middlewareのロード
	r.Use(VersionMiddleware(showVersion))

	r.Use(LogMiddleware(loggerAccess, time.RFC3339, false))
	r.Use(CORSMiddleware())

	//NOTE ズーパー重要
	r.Use(ServiceKeyMiddleware(registry))

	// auth middlewareの準備
	authenticator, err := New(
		&UserPool{
			Region: os.Getenv("COGNITO_REGION"),
			PoolID: os.Getenv("COGNITO_USER_POOL_ID"),
		},
		&Option{
			NoVerification: false,
		})

	defineRoutes(r, authenticator, registry.NewUsers(), cacheMaxAge)

	ip := os.Getenv(ipEnv)

	port := os.Getenv(portEnv)
	if port == "" {
		port = "443"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", ip, port),
		Handler: r,
	}

	// Start server
	logger.Infof("Server listening on port:%v", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("listen: %s\n", err)
	}

	return nil
}
