package server

import (
	"net"

	ct "github.com/kubearmor/KVMService/src/constants"
	etcd "github.com/kubearmor/KVMService/src/etcd"
	kg "github.com/kubearmor/KVMService/src/log"
	pb "github.com/kubearmor/KVMService/src/service/protobuf"
	tp "github.com/kubearmor/KVMService/src/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	PolicyChan  chan tp.K8sKubeArmorHostPolicyEventWithIdentity
	ClusterIp   string
	Clusterport string
	podIp       string
	EtcdClient  *etcd.EtcdClient
)

// Variables / Struct
type Server struct {
	podIp      string
	port       string
	EtcdClient *etcd.EtcdClient
}

func NewServerInit(ipAddress, ClusterIpAddress, portVal string, Etcd *etcd.EtcdClient) *Server {
	kg.Printf("Initiliazing the KVMServer => podip:%v clusterIP:%v clusterPort:%v", ipAddress, ClusterIpAddress, portVal)
	podIp = ipAddress
	Clusterport = portVal
	EtcdClient = Etcd
	ClusterIp = ClusterIpAddress
	return &Server{podIp: ipAddress, port: portVal, EtcdClient: EtcdClient}
}

func (s *Server) InitServer() {
	// TCP connection - Listen on port specified in input
	PolicyChan = make(chan tp.K8sKubeArmorHostPolicyEventWithIdentity)
	tcpConn, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		kg.Printf("Error listening on port %s", s.port)
		return
	} else {
		kg.Printf("Successfully KVMServer Listening on port %s", s.port)
	}

	creds, err := credentials.NewServerTLSFromFile(ct.ServerCertPath, ct.ServerKeyPath)
	if err != nil {
		kg.Errf("Failed to setup TLS: %v", err)
		return
	}

	// Start gRPC server and register for protobuf
	gRPCServer := grpc.NewServer(grpc.Creds(creds))
	if gRPCServer == nil {
		kg.Err("Failed to serve gRPCServer is null")
		return
	}

	// Register KVM Server
	pb.RegisterKVMServer(gRPCServer, &KVMServer{})

	// Register CLIHandler Server
	pb.RegisterHandleCliServer(gRPCServer, &CLIServer{})

	err = gRPCServer.Serve(tcpConn)
	if err != nil {
		kg.Err("Failed to serve")
		return
	}
}
