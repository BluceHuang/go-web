package cache

import (
	"log"
	"sync"

	"github.com/go-redis/redis"
)

var redisOnce sync.Once
var redisMutex sync.Mutex
var redisClusterClientMap map[string]*redis.ClusterClient
var redisClientMap map[string]*redis.Client

type Cacher interface {
	Client() *redis.ClusterClient
}

type Cache struct {
	ClientClusterMap map[string]*redis.ClusterClient
	ClientMap        map[string]*redis.Client
}

func BuildCache() Cache {
	return Cache{
		ClientClusterMap: redisClusterClientMap,
		ClientMap:        redisClientMap,
	}
}

func (cache Cache) Client() *redis.ClusterClient {
	return cache.ClientClusterMap["master"]
}

func init() {
	redisConfigs := goweb.ServerConfig().Redis
	redisOnce.Do(func() {
		redisMutex.Lock()
		defer redisMutex.Unlock()

		redisClientMap = make(map[string]*redis.Client, len(redisConfigs))
		redisClusterClientMap = make(map[string]*redis.ClusterClient, len(redisConfigs))
		for _, redisConfig := range redisConfigs {
			var err error
			var redisOptions *redis.Options

			if len(redisConfig.Addrs) == 1 {
				// 解析redis链接地址
				if redisOptions, err = redis.ParseURL(redisConfig.Addrs[0]); err != nil {
					log.Printf("redis parse config failed: %v", err)
					return
				}

				// 连接redis数据库
				client := redis.NewClient(redisOptions)
				if _, err := client.Ping().Result(); err != nil {
					log.Printf("redis ping failed: %v", err)
					return
				}

				// 保存mongodb客户端
				redisClientMap[redisConfig.Name] = client
			} else {
				// 连接redis cluster数据库
				client := redis.NewClusterClient(&redis.ClusterOptions{
					Addrs:    redisConfig.Addrs,    //set redis cluster url
					Password: redisConfig.Password, //set password
				})

				if _, err := client.Ping().Result(); err != nil {
					log.Printf("redis ping failed: %v", err)
					return
				}

				// 保存mongodb客户端
				redisClusterClientMap[redisConfig.Name] = client
			}
		}
	})
}
