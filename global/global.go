package global

import (
	"sync"

	"github.com/go-redis/redis/v8"

	"go-generate/utils/timer"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"go-generate/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	YAN_DB     *gorm.DB
	YAN_DBList map[string]*gorm.DB
	YAN_REDIS  *redis.Client
	YAN_CONFIG config.Server
	YAN_VP     *viper.Viper
	// YAN_LOG    *oplogging.Logger
	YAN_LOG                 *zap.Logger
	YAN_Timer               timer.Timer = timer.NewTimerTask()
	YAN_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return YAN_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := YAN_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
