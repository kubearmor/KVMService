// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeArmor

package core

import (
	//"context"

	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	kc "github.com/kubearmor/KVMService/src/common"
	ct "github.com/kubearmor/KVMService/src/constants"
	etcd "github.com/kubearmor/KVMService/src/etcd"
	kg "github.com/kubearmor/KVMService/src/log"
	"github.com/kubearmor/KVMService/src/service/cilium"
	gs "github.com/kubearmor/KVMService/src/service/genscript"
	ks "github.com/kubearmor/KVMService/src/service/server"
	tp "github.com/kubearmor/KVMService/src/types"
)

// ClientConn is the wrapper for a grpc client conn
type ClientConn struct {
	// conn      *grpc.ClientConn
	// unhealthy bool
}

// ====================== //
// == KubeArmor Daemon == //
// ====================== //

// StopChan Channel
var StopChan chan struct{}

// init Function
func init() {
	StopChan = make(chan struct{})
}

// KVMS Structure
type KVMS struct {
	EtcdClient *etcd.EtcdClient
	Server     *ks.Server

	// gRPC
	gRPCPort  string
	LogPath   string
	LogFilter string

	IdentityConnPool []ClientConn

	MapEtcdEWIdentityLabels map[string]string
	EtcdEWLabels            []string

	// Host Security policies
	HostSecurityPolicies     []tp.HostSecurityPolicy
	HostSecurityPoliciesLock *sync.RWMutex

	// Virtual Machine policies and mappers
	VirtualMachineSecurityPolicies     []tp.VirtualMachineSecurityPolicy
	VirtualMachineSecurityPoliciesLock *sync.RWMutex

	MapIdentityToEWName           map[uint16]string
	MapEWNameToIdentity           map[string]uint16
	MapIdentityToLabel            map[uint16]string
	MapLabelToIdentities          map[string][]uint16
	MapVirtualMachineConnIdentity map[uint16]ClientConn

	ClusterPort      uint16
	ClusteripAddress string
	EtcdPort         uint16
	PodIpAddress     string

	// WgDaemon Handler
	WgDaemon sync.WaitGroup
}

// NewKVMSDaemon Function
func NewKVMSDaemon(port, etcdPort int, isnonk8s bool) *KVMS {
	kg.Print("Initializing all the KVMS daemon attributes")
	var err error
	dm := new(KVMS)

	dm.ClusterPort = uint16(port)
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	dm.PodIpAddress = localAddr.IP.String()

	if !isnonk8s {
		// Get kvmservice external ip
		dm.ClusteripAddress, err = kc.GetExternalIP(ct.KvmServiceAccountName)
		if err != nil {
			kg.Err(err.Error())
			return nil
		}
	} else {
		kc.IsNonK8sEnv = true
		// Set IP as localhost for nonk8s environment
		dm.ClusteripAddress = dm.PodIpAddress
	}

	dm.EtcdPort = uint16(etcdPort)
	dm.EtcdClient = etcd.NewEtcdClient()

	dm.MapEtcdEWIdentityLabels = make(map[string]string)
	dm.EtcdEWLabels = make([]string, 0)

	dm.gRPCPort = ""
	dm.LogPath = ""
	dm.LogFilter = ""
	dm.IdentityConnPool = nil

	dm.Server = ks.NewServerInit(dm.PodIpAddress, dm.ClusteripAddress, strconv.FormatUint(uint64(dm.ClusterPort), 10), dm.EtcdClient)

	dm.HostSecurityPolicies = []tp.HostSecurityPolicy{}
	dm.HostSecurityPoliciesLock = new(sync.RWMutex)
	dm.VirtualMachineSecurityPoliciesLock = new(sync.RWMutex)

	dm.MapIdentityToEWName = make(map[uint16]string)
	dm.MapEWNameToIdentity = make(map[string]uint16)

	dm.MapIdentityToLabel = make(map[uint16]string)
	dm.MapLabelToIdentities = make(map[string][]uint16)
	dm.MapVirtualMachineConnIdentity = make(map[uint16]ClientConn)

	dm.WgDaemon = sync.WaitGroup{}
	kg.Print("KVMService attributes got initialized")

	return dm
}

// DestroyKVMS Function
func (dm *KVMS) DestroyKVMS() {

	// wait for a while
	time.Sleep(time.Second * 1)

	// wait for other routines
	kg.Print("Waiting for remaining routine terminations")
	dm.WgDaemon.Wait()
}

// ==================== //
// == Signal Handler == //
// ==================== //

// GetOSSigChannel Function
func GetOSSigChannel() chan os.Signal {
	c := make(chan os.Signal, 1)

	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)

	return c
}

// ========== //
// == Main == //
// ========== //

// KVMSDaemon Function
func KVMSDaemon(portPtr, etcdPort int, nonk8s bool) {

	// create a daemon
	dm := NewKVMSDaemon(portPtr, etcdPort, nonk8s)

	// wait for a while
	time.Sleep(time.Second * 1)

	// == //
	gs.InitGenScript(dm.ClusterPort, dm.EtcdPort, dm.ClusteripAddress)

	if !nonk8s {
		if K8s.InitK8sClient() {
			// watch host security policies
			kg.Print("K8S Client is successfully initialize")

			kg.Print("Watcher triggered for the host policies")
			go dm.WatchHostSecurityPolicies()

		} else {
			kg.Print("K8S client initialization got failed")
		}
	} else {
		// Start http server
		kg.Print("Starting HTTP Server")
		go ks.InitHttpServer(dm.UpdateHostSecurityPolicies, dm.HandleNetworkPolicyUpdates, dm.HandleVm, dm.ListOnboardedVms)

		kg.Print("Starting Cilium Node Registration Observer")
		cilium.NodeRegisterWatcherInit(dm.EtcdClient, dm.MapEWNameToIdentity)
	}

	kg.Print("Triggered the keepalive ETCD client")
	go dm.EtcdClient.KeepAliveEtcdConnection()

	kg.Print("Starting gRPC server")
	go dm.Server.InitServer()

	// wait for a while
	time.Sleep(time.Second * 1)

	// listen for interrupt signals
	sigChan := GetOSSigChannel()
	<-sigChan
	close(StopChan)
	close(ks.PolicyChan)

	// destroy the daemon
	dm.DestroyKVMS()
}
