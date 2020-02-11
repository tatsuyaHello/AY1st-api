package registry

import (
	"AY1st/repository"
	"AY1st/service"

	"github.com/go-xorm/xorm"
)

// TODO: registryの役割をはっきりさせる

const (
	// RepositoryKey はリポジトリレジストリ取得キー名
	RepositoryKey = "repository_registry"
	// ServiceKey はサービスレジストリ取得キー名
	ServiceKey = "service_factory"
)

// RepositoryMaker はリポジトリレジストリ
type RepositoryMaker interface {
	// MASTER
	NewHealthCheck() repository.HealthCheckInterface

	// TRANSACTION
}

// RepositorySettings はリポジトリレジストリの設定
// 全ての設定が必須
type RepositorySettings struct {
	Engine xorm.EngineInterface
}

// Repository はリポジトリレジストリの実装
// インフラ層の依存情報を初期化時に注入する
type Repository struct {
	settings *RepositorySettings
}

// NewRepository initializes factory with injected infra.
func NewRepository(settings *RepositorySettings) RepositoryMaker {
	r := &Repository{
		settings: settings,
	}
	return r
}

// NewHealthCheck returns HealthCheck repository.
func (r *Repository) NewHealthCheck() repository.HealthCheckInterface {
	usersRepo := repository.NewHealthCheck(r.settings.Engine)
	return usersRepo
}

// Servicer はサービスレジストリ
type Servicer interface {
	NewHealthCheck() service.HealthCheckInterface
	NewUsers() service.UsersInterface
}

// ServiceRegistrySettings はサービスレジストリの設定
// 全ての設定が必須
type ServiceRegistrySettings struct {
	Engine xorm.EngineInterface
}

// ServiceRegistry はサービスレジストリの実装
// インフラ層の依存情報を初期化時に注入する
type ServiceRegistry struct {
	settings  *ServiceRegistrySettings
	repoMaker RepositoryMaker
}

// NewService initializes factory with injected infra.
func NewService(settings *ServiceRegistrySettings) *ServiceRegistry {
	r := &ServiceRegistry{
		settings: settings,
	}

	r.repoMaker = NewRepository(&RepositorySettings{Engine: settings.Engine})
	return r
}

// NewHealthCheck returns HealthCheck service.
func (r *ServiceRegistry) NewHealthCheck() service.HealthCheckInterface {
	healthCheckRepo := repository.NewHealthCheck(r.settings.Engine)
	return service.NewHealthCheck(healthCheckRepo)
}

// NewUsers returns Users service.
func (r *ServiceRegistry) NewUsers() service.UsersInterface {
	UsersRepo := repository.NewUsers(r.settings.Engine)
	UserIdentitiesRepo := repository.NewUserIdentities(r.settings.Engine)
	return service.NewUsers(UsersRepo, UserIdentitiesRepo)
}
