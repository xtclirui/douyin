package user

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"user/conf"
	"user/utils/redis"
)

func main() {
	// 初始化 DB *gorm.DB
	conf.Init()
	// 初始化redis-DB0的连接，follow选择的DB0.
	redis.InitRedis()
	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"), // 微服务名字
		micro.Address("127.0.0.1:8081"),
		micro.Registry(etcdReg), // etcd注册件
		micro.Metadata(map[string]string{"protocol": "http"}),
	)

}
