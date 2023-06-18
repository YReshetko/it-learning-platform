package config

type GRPCServer struct {
	Host    string `env:"GRPC_HOST"`
	Port    int    `env:"GRPC_PORT"`
	Network string `env:"GRPC_NETWORK"`
}
