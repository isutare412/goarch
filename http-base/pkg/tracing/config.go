package tracing

type Config struct {
	Enabled                 bool
	ServiceName             string
	Environment             string
	JaegerCollectorEndpoint string
}
