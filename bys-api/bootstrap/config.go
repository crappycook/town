package bootstrap

type Config struct {
	ServiceName string        `yaml:"service_name"`
	Log         LogConfig     `yaml:"log"`
	RestApi     RestApiConfig `yaml:"rest_api"`
	// grpc config
	HostStatusRpc RpcConfig              `yaml:"host_status_rpc"`
	Mysql         map[string]MysqlConfig `yaml:"mysql"` // keyed by db name
}

type LogConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	Dir    string `yaml:"dir"`
}

type RestApiConfig struct {
	ListenPort string `yaml:"listen_port"`
}

type RpcConfig struct {
	Addr string `yaml:"addr"`
}

type MysqlConfig struct {
	Addr     string `yaml:"addr"`
	DBName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
