package redis

//// 声明一个全局的rab变量
//var rdb *redis.Client
//
//// 初始化连接
//func InitClient() (err error) {
//	rdb = redis.NewClient(&redis.Options{
//		Addr: fmt.Sprintf("%s:%d",
//			viper.GetString("redis.host"),
//			viper.GetInt("redis.port"),
//		),
//		Password: viper.GetString("viper.password"),
//		DB:       viper.GetInt("redis.db"),
//		PoolSize: viper.GetInt("redis.pool_size"),
//	})
//	_, err = rdb.Ping().Result()
//	return
//}
//
//func Close() {
//	_ = rdb.Close()
//}
