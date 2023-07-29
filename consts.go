package logger

const (
	KeyServiceName = "service.name"

	KeyRequestID = "request_id"

	KeyGRPCService          = "grpc.service"
	KeyGRPCMethod           = "grpc.method"
	KeyGRPCRequestBody      = "grpc.request.body"
	KeyGRPCRequestStartTime = "grpc.request.start_time"
	KeyGRPCRequestDeadline  = "grpc.request.deadline"
	KeyGRPCRequestResponse  = "grpc.request.response"

	KeyWatcherName   = "watcher.name"
	KeyWatcherParams = "watcher.params"

	KeyRMQConnectionName   = "rmq.connection_name"
	KeyRMQServerName       = "rmq.server_name"
	KeyRMQExchange         = "rmq.exchange"
	KeyRMQRoutingKey       = "rmq.routing_key"
	KeyRMQMsgBody          = "rmq.msg.body"
	KeyRMQHandlerStartTime = "rmq.handler.start_time"
	KeyRMQHandlerPanicMsg  = "rmq.handler.panic_msg"
)
