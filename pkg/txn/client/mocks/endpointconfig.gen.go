// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"crypto/tls"
	"sync"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type EndpointConfig struct {
	TimeoutStub        func(fab.TimeoutType) time.Duration
	timeoutMutex       sync.RWMutex
	timeoutArgsForCall []struct {
		arg1 fab.TimeoutType
	}
	timeoutReturns struct {
		result1 time.Duration
	}
	timeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	OrderersConfigStub        func() []fab.OrdererConfig
	orderersConfigMutex       sync.RWMutex
	orderersConfigArgsForCall []struct{}
	orderersConfigReturns     struct {
		result1 []fab.OrdererConfig
	}
	orderersConfigReturnsOnCall map[int]struct {
		result1 []fab.OrdererConfig
	}
	OrdererConfigStub        func(nameOrURL string) (*fab.OrdererConfig, bool)
	ordererConfigMutex       sync.RWMutex
	ordererConfigArgsForCall []struct {
		nameOrURL string
	}
	ordererConfigReturns struct {
		result1 *fab.OrdererConfig
		result2 bool
	}
	ordererConfigReturnsOnCall map[int]struct {
		result1 *fab.OrdererConfig
		result2 bool
	}
	PeersConfigStub        func(org string) ([]fab.PeerConfig, bool)
	peersConfigMutex       sync.RWMutex
	peersConfigArgsForCall []struct {
		org string
	}
	peersConfigReturns struct {
		result1 []fab.PeerConfig
		result2 bool
	}
	peersConfigReturnsOnCall map[int]struct {
		result1 []fab.PeerConfig
		result2 bool
	}
	PeerConfigStub        func(nameOrURL string) (*fab.PeerConfig, bool)
	peerConfigMutex       sync.RWMutex
	peerConfigArgsForCall []struct {
		nameOrURL string
	}
	peerConfigReturns struct {
		result1 *fab.PeerConfig
		result2 bool
	}
	peerConfigReturnsOnCall map[int]struct {
		result1 *fab.PeerConfig
		result2 bool
	}
	NetworkConfigStub        func() *fab.NetworkConfig
	networkConfigMutex       sync.RWMutex
	networkConfigArgsForCall []struct{}
	networkConfigReturns     struct {
		result1 *fab.NetworkConfig
	}
	networkConfigReturnsOnCall map[int]struct {
		result1 *fab.NetworkConfig
	}
	NetworkPeersStub        func() []fab.NetworkPeer
	networkPeersMutex       sync.RWMutex
	networkPeersArgsForCall []struct{}
	networkPeersReturns     struct {
		result1 []fab.NetworkPeer
	}
	networkPeersReturnsOnCall map[int]struct {
		result1 []fab.NetworkPeer
	}
	ChannelConfigStub        func(name string) *fab.ChannelEndpointConfig
	channelConfigMutex       sync.RWMutex
	channelConfigArgsForCall []struct {
		name string
	}
	channelConfigReturns struct {
		result1 *fab.ChannelEndpointConfig
	}
	channelConfigReturnsOnCall map[int]struct {
		result1 *fab.ChannelEndpointConfig
	}
	ChannelPeersStub        func(name string) []fab.ChannelPeer
	channelPeersMutex       sync.RWMutex
	channelPeersArgsForCall []struct {
		name string
	}
	channelPeersReturns struct {
		result1 []fab.ChannelPeer
	}
	channelPeersReturnsOnCall map[int]struct {
		result1 []fab.ChannelPeer
	}
	ChannelOrderersStub        func(name string) []fab.OrdererConfig
	channelOrderersMutex       sync.RWMutex
	channelOrderersArgsForCall []struct {
		name string
	}
	channelOrderersReturns struct {
		result1 []fab.OrdererConfig
	}
	channelOrderersReturnsOnCall map[int]struct {
		result1 []fab.OrdererConfig
	}
	TLSCACertPoolStub        func() fab.CertPool
	tLSCACertPoolMutex       sync.RWMutex
	tLSCACertPoolArgsForCall []struct{}
	tLSCACertPoolReturns     struct {
		result1 fab.CertPool
	}
	tLSCACertPoolReturnsOnCall map[int]struct {
		result1 fab.CertPool
	}
	TLSClientCertsStub        func() []tls.Certificate
	tLSClientCertsMutex       sync.RWMutex
	tLSClientCertsArgsForCall []struct{}
	tLSClientCertsReturns     struct {
		result1 []tls.Certificate
	}
	tLSClientCertsReturnsOnCall map[int]struct {
		result1 []tls.Certificate
	}
	CryptoConfigPathStub        func() string
	cryptoConfigPathMutex       sync.RWMutex
	cryptoConfigPathArgsForCall []struct{}
	cryptoConfigPathReturns     struct {
		result1 string
	}
	cryptoConfigPathReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EndpointConfig) Timeout(arg1 fab.TimeoutType) time.Duration {
	fake.timeoutMutex.Lock()
	ret, specificReturn := fake.timeoutReturnsOnCall[len(fake.timeoutArgsForCall)]
	fake.timeoutArgsForCall = append(fake.timeoutArgsForCall, struct {
		arg1 fab.TimeoutType
	}{arg1})
	fake.recordInvocation("Timeout", []interface{}{arg1})
	fake.timeoutMutex.Unlock()
	if fake.TimeoutStub != nil {
		return fake.TimeoutStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.timeoutReturns.result1
}

func (fake *EndpointConfig) TimeoutCallCount() int {
	fake.timeoutMutex.RLock()
	defer fake.timeoutMutex.RUnlock()
	return len(fake.timeoutArgsForCall)
}

func (fake *EndpointConfig) TimeoutArgsForCall(i int) fab.TimeoutType {
	fake.timeoutMutex.RLock()
	defer fake.timeoutMutex.RUnlock()
	return fake.timeoutArgsForCall[i].arg1
}

func (fake *EndpointConfig) TimeoutReturns(result1 time.Duration) {
	fake.TimeoutStub = nil
	fake.timeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *EndpointConfig) TimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.TimeoutStub = nil
	if fake.timeoutReturnsOnCall == nil {
		fake.timeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.timeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *EndpointConfig) OrderersConfig() []fab.OrdererConfig {
	fake.orderersConfigMutex.Lock()
	ret, specificReturn := fake.orderersConfigReturnsOnCall[len(fake.orderersConfigArgsForCall)]
	fake.orderersConfigArgsForCall = append(fake.orderersConfigArgsForCall, struct{}{})
	fake.recordInvocation("OrderersConfig", []interface{}{})
	fake.orderersConfigMutex.Unlock()
	if fake.OrderersConfigStub != nil {
		return fake.OrderersConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.orderersConfigReturns.result1
}

func (fake *EndpointConfig) OrderersConfigCallCount() int {
	fake.orderersConfigMutex.RLock()
	defer fake.orderersConfigMutex.RUnlock()
	return len(fake.orderersConfigArgsForCall)
}

func (fake *EndpointConfig) OrderersConfigReturns(result1 []fab.OrdererConfig) {
	fake.OrderersConfigStub = nil
	fake.orderersConfigReturns = struct {
		result1 []fab.OrdererConfig
	}{result1}
}

func (fake *EndpointConfig) OrderersConfigReturnsOnCall(i int, result1 []fab.OrdererConfig) {
	fake.OrderersConfigStub = nil
	if fake.orderersConfigReturnsOnCall == nil {
		fake.orderersConfigReturnsOnCall = make(map[int]struct {
			result1 []fab.OrdererConfig
		})
	}
	fake.orderersConfigReturnsOnCall[i] = struct {
		result1 []fab.OrdererConfig
	}{result1}
}

func (fake *EndpointConfig) OrdererConfig(nameOrURL string) (*fab.OrdererConfig, bool) {
	fake.ordererConfigMutex.Lock()
	ret, specificReturn := fake.ordererConfigReturnsOnCall[len(fake.ordererConfigArgsForCall)]
	fake.ordererConfigArgsForCall = append(fake.ordererConfigArgsForCall, struct {
		nameOrURL string
	}{nameOrURL})
	fake.recordInvocation("OrdererConfig", []interface{}{nameOrURL})
	fake.ordererConfigMutex.Unlock()
	if fake.OrdererConfigStub != nil {
		return fake.OrdererConfigStub(nameOrURL)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.ordererConfigReturns.result1, fake.ordererConfigReturns.result2
}

func (fake *EndpointConfig) OrdererConfigCallCount() int {
	fake.ordererConfigMutex.RLock()
	defer fake.ordererConfigMutex.RUnlock()
	return len(fake.ordererConfigArgsForCall)
}

func (fake *EndpointConfig) OrdererConfigArgsForCall(i int) string {
	fake.ordererConfigMutex.RLock()
	defer fake.ordererConfigMutex.RUnlock()
	return fake.ordererConfigArgsForCall[i].nameOrURL
}

func (fake *EndpointConfig) OrdererConfigReturns(result1 *fab.OrdererConfig, result2 bool) {
	fake.OrdererConfigStub = nil
	fake.ordererConfigReturns = struct {
		result1 *fab.OrdererConfig
		result2 bool
	}{result1, result2}
}

func (fake *EndpointConfig) OrdererConfigReturnsOnCall(i int, result1 *fab.OrdererConfig, result2 bool) {
	fake.OrdererConfigStub = nil
	if fake.ordererConfigReturnsOnCall == nil {
		fake.ordererConfigReturnsOnCall = make(map[int]struct {
			result1 *fab.OrdererConfig
			result2 bool
		})
	}
	fake.ordererConfigReturnsOnCall[i] = struct {
		result1 *fab.OrdererConfig
		result2 bool
	}{result1, result2}
}

func (fake *EndpointConfig) PeersConfig(org string) ([]fab.PeerConfig, bool) {
	fake.peersConfigMutex.Lock()
	ret, specificReturn := fake.peersConfigReturnsOnCall[len(fake.peersConfigArgsForCall)]
	fake.peersConfigArgsForCall = append(fake.peersConfigArgsForCall, struct {
		org string
	}{org})
	fake.recordInvocation("PeersConfig", []interface{}{org})
	fake.peersConfigMutex.Unlock()
	if fake.PeersConfigStub != nil {
		return fake.PeersConfigStub(org)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.peersConfigReturns.result1, fake.peersConfigReturns.result2
}

func (fake *EndpointConfig) PeersConfigCallCount() int {
	fake.peersConfigMutex.RLock()
	defer fake.peersConfigMutex.RUnlock()
	return len(fake.peersConfigArgsForCall)
}

func (fake *EndpointConfig) PeersConfigArgsForCall(i int) string {
	fake.peersConfigMutex.RLock()
	defer fake.peersConfigMutex.RUnlock()
	return fake.peersConfigArgsForCall[i].org
}

func (fake *EndpointConfig) PeersConfigReturns(result1 []fab.PeerConfig, result2 bool) {
	fake.PeersConfigStub = nil
	fake.peersConfigReturns = struct {
		result1 []fab.PeerConfig
		result2 bool
	}{result1, result2}
}

func (fake *EndpointConfig) PeersConfigReturnsOnCall(i int, result1 []fab.PeerConfig, result2 bool) {
	fake.PeersConfigStub = nil
	if fake.peersConfigReturnsOnCall == nil {
		fake.peersConfigReturnsOnCall = make(map[int]struct {
			result1 []fab.PeerConfig
			result2 bool
		})
	}
	fake.peersConfigReturnsOnCall[i] = struct {
		result1 []fab.PeerConfig
		result2 bool
	}{result1, result2}
}

func (fake *EndpointConfig) PeerConfig(nameOrURL string) (*fab.PeerConfig, bool) {
	fake.peerConfigMutex.Lock()
	ret, specificReturn := fake.peerConfigReturnsOnCall[len(fake.peerConfigArgsForCall)]
	fake.peerConfigArgsForCall = append(fake.peerConfigArgsForCall, struct {
		nameOrURL string
	}{nameOrURL})
	fake.recordInvocation("PeerConfig", []interface{}{nameOrURL})
	fake.peerConfigMutex.Unlock()
	if fake.PeerConfigStub != nil {
		return fake.PeerConfigStub(nameOrURL)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.peerConfigReturns.result1, fake.peerConfigReturns.result2
}

func (fake *EndpointConfig) PeerConfigCallCount() int {
	fake.peerConfigMutex.RLock()
	defer fake.peerConfigMutex.RUnlock()
	return len(fake.peerConfigArgsForCall)
}

func (fake *EndpointConfig) PeerConfigArgsForCall(i int) string {
	fake.peerConfigMutex.RLock()
	defer fake.peerConfigMutex.RUnlock()
	return fake.peerConfigArgsForCall[i].nameOrURL
}

func (fake *EndpointConfig) PeerConfigReturns(result1 *fab.PeerConfig, result2 bool) {
	fake.PeerConfigStub = nil
	fake.peerConfigReturns = struct {
		result1 *fab.PeerConfig
		result2 bool
	}{result1, result2}
}

func (fake *EndpointConfig) PeerConfigReturnsOnCall(i int, result1 *fab.PeerConfig, result2 bool) {
	fake.PeerConfigStub = nil
	if fake.peerConfigReturnsOnCall == nil {
		fake.peerConfigReturnsOnCall = make(map[int]struct {
			result1 *fab.PeerConfig
			result2 bool
		})
	}
	fake.peerConfigReturnsOnCall[i] = struct {
		result1 *fab.PeerConfig
		result2 bool
	}{result1, result2}
}

func (fake *EndpointConfig) NetworkConfig() *fab.NetworkConfig {
	fake.networkConfigMutex.Lock()
	ret, specificReturn := fake.networkConfigReturnsOnCall[len(fake.networkConfigArgsForCall)]
	fake.networkConfigArgsForCall = append(fake.networkConfigArgsForCall, struct{}{})
	fake.recordInvocation("NetworkConfig", []interface{}{})
	fake.networkConfigMutex.Unlock()
	if fake.NetworkConfigStub != nil {
		return fake.NetworkConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.networkConfigReturns.result1
}

func (fake *EndpointConfig) NetworkConfigCallCount() int {
	fake.networkConfigMutex.RLock()
	defer fake.networkConfigMutex.RUnlock()
	return len(fake.networkConfigArgsForCall)
}

func (fake *EndpointConfig) NetworkConfigReturns(result1 *fab.NetworkConfig) {
	fake.NetworkConfigStub = nil
	fake.networkConfigReturns = struct {
		result1 *fab.NetworkConfig
	}{result1}
}

func (fake *EndpointConfig) NetworkConfigReturnsOnCall(i int, result1 *fab.NetworkConfig) {
	fake.NetworkConfigStub = nil
	if fake.networkConfigReturnsOnCall == nil {
		fake.networkConfigReturnsOnCall = make(map[int]struct {
			result1 *fab.NetworkConfig
		})
	}
	fake.networkConfigReturnsOnCall[i] = struct {
		result1 *fab.NetworkConfig
	}{result1}
}

func (fake *EndpointConfig) NetworkPeers() []fab.NetworkPeer {
	fake.networkPeersMutex.Lock()
	ret, specificReturn := fake.networkPeersReturnsOnCall[len(fake.networkPeersArgsForCall)]
	fake.networkPeersArgsForCall = append(fake.networkPeersArgsForCall, struct{}{})
	fake.recordInvocation("NetworkPeers", []interface{}{})
	fake.networkPeersMutex.Unlock()
	if fake.NetworkPeersStub != nil {
		return fake.NetworkPeersStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.networkPeersReturns.result1
}

func (fake *EndpointConfig) NetworkPeersCallCount() int {
	fake.networkPeersMutex.RLock()
	defer fake.networkPeersMutex.RUnlock()
	return len(fake.networkPeersArgsForCall)
}

func (fake *EndpointConfig) NetworkPeersReturns(result1 []fab.NetworkPeer) {
	fake.NetworkPeersStub = nil
	fake.networkPeersReturns = struct {
		result1 []fab.NetworkPeer
	}{result1}
}

func (fake *EndpointConfig) NetworkPeersReturnsOnCall(i int, result1 []fab.NetworkPeer) {
	fake.NetworkPeersStub = nil
	if fake.networkPeersReturnsOnCall == nil {
		fake.networkPeersReturnsOnCall = make(map[int]struct {
			result1 []fab.NetworkPeer
		})
	}
	fake.networkPeersReturnsOnCall[i] = struct {
		result1 []fab.NetworkPeer
	}{result1}
}

func (fake *EndpointConfig) ChannelConfig(name string) *fab.ChannelEndpointConfig {
	fake.channelConfigMutex.Lock()
	ret, specificReturn := fake.channelConfigReturnsOnCall[len(fake.channelConfigArgsForCall)]
	fake.channelConfigArgsForCall = append(fake.channelConfigArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("ChannelConfig", []interface{}{name})
	fake.channelConfigMutex.Unlock()
	if fake.ChannelConfigStub != nil {
		return fake.ChannelConfigStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.channelConfigReturns.result1
}

func (fake *EndpointConfig) ChannelConfigCallCount() int {
	fake.channelConfigMutex.RLock()
	defer fake.channelConfigMutex.RUnlock()
	return len(fake.channelConfigArgsForCall)
}

func (fake *EndpointConfig) ChannelConfigArgsForCall(i int) string {
	fake.channelConfigMutex.RLock()
	defer fake.channelConfigMutex.RUnlock()
	return fake.channelConfigArgsForCall[i].name
}

func (fake *EndpointConfig) ChannelConfigReturns(result1 *fab.ChannelEndpointConfig) {
	fake.ChannelConfigStub = nil
	fake.channelConfigReturns = struct {
		result1 *fab.ChannelEndpointConfig
	}{result1}
}

func (fake *EndpointConfig) ChannelConfigReturnsOnCall(i int, result1 *fab.ChannelEndpointConfig) {
	fake.ChannelConfigStub = nil
	if fake.channelConfigReturnsOnCall == nil {
		fake.channelConfigReturnsOnCall = make(map[int]struct {
			result1 *fab.ChannelEndpointConfig
		})
	}
	fake.channelConfigReturnsOnCall[i] = struct {
		result1 *fab.ChannelEndpointConfig
	}{result1}
}

func (fake *EndpointConfig) ChannelPeers(name string) []fab.ChannelPeer {
	fake.channelPeersMutex.Lock()
	ret, specificReturn := fake.channelPeersReturnsOnCall[len(fake.channelPeersArgsForCall)]
	fake.channelPeersArgsForCall = append(fake.channelPeersArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("ChannelPeers", []interface{}{name})
	fake.channelPeersMutex.Unlock()
	if fake.ChannelPeersStub != nil {
		return fake.ChannelPeersStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.channelPeersReturns.result1
}

func (fake *EndpointConfig) ChannelPeersCallCount() int {
	fake.channelPeersMutex.RLock()
	defer fake.channelPeersMutex.RUnlock()
	return len(fake.channelPeersArgsForCall)
}

func (fake *EndpointConfig) ChannelPeersArgsForCall(i int) string {
	fake.channelPeersMutex.RLock()
	defer fake.channelPeersMutex.RUnlock()
	return fake.channelPeersArgsForCall[i].name
}

func (fake *EndpointConfig) ChannelPeersReturns(result1 []fab.ChannelPeer) {
	fake.ChannelPeersStub = nil
	fake.channelPeersReturns = struct {
		result1 []fab.ChannelPeer
	}{result1}
}

func (fake *EndpointConfig) ChannelPeersReturnsOnCall(i int, result1 []fab.ChannelPeer) {
	fake.ChannelPeersStub = nil
	if fake.channelPeersReturnsOnCall == nil {
		fake.channelPeersReturnsOnCall = make(map[int]struct {
			result1 []fab.ChannelPeer
		})
	}
	fake.channelPeersReturnsOnCall[i] = struct {
		result1 []fab.ChannelPeer
	}{result1}
}

func (fake *EndpointConfig) ChannelOrderers(name string) []fab.OrdererConfig {
	fake.channelOrderersMutex.Lock()
	ret, specificReturn := fake.channelOrderersReturnsOnCall[len(fake.channelOrderersArgsForCall)]
	fake.channelOrderersArgsForCall = append(fake.channelOrderersArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("ChannelOrderers", []interface{}{name})
	fake.channelOrderersMutex.Unlock()
	if fake.ChannelOrderersStub != nil {
		return fake.ChannelOrderersStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.channelOrderersReturns.result1
}

func (fake *EndpointConfig) ChannelOrderersCallCount() int {
	fake.channelOrderersMutex.RLock()
	defer fake.channelOrderersMutex.RUnlock()
	return len(fake.channelOrderersArgsForCall)
}

func (fake *EndpointConfig) ChannelOrderersArgsForCall(i int) string {
	fake.channelOrderersMutex.RLock()
	defer fake.channelOrderersMutex.RUnlock()
	return fake.channelOrderersArgsForCall[i].name
}

func (fake *EndpointConfig) ChannelOrderersReturns(result1 []fab.OrdererConfig) {
	fake.ChannelOrderersStub = nil
	fake.channelOrderersReturns = struct {
		result1 []fab.OrdererConfig
	}{result1}
}

func (fake *EndpointConfig) ChannelOrderersReturnsOnCall(i int, result1 []fab.OrdererConfig) {
	fake.ChannelOrderersStub = nil
	if fake.channelOrderersReturnsOnCall == nil {
		fake.channelOrderersReturnsOnCall = make(map[int]struct {
			result1 []fab.OrdererConfig
		})
	}
	fake.channelOrderersReturnsOnCall[i] = struct {
		result1 []fab.OrdererConfig
	}{result1}
}

func (fake *EndpointConfig) TLSCACertPool() fab.CertPool {
	fake.tLSCACertPoolMutex.Lock()
	ret, specificReturn := fake.tLSCACertPoolReturnsOnCall[len(fake.tLSCACertPoolArgsForCall)]
	fake.tLSCACertPoolArgsForCall = append(fake.tLSCACertPoolArgsForCall, struct{}{})
	fake.recordInvocation("TLSCACertPool", []interface{}{})
	fake.tLSCACertPoolMutex.Unlock()
	if fake.TLSCACertPoolStub != nil {
		return fake.TLSCACertPoolStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.tLSCACertPoolReturns.result1
}

func (fake *EndpointConfig) TLSCACertPoolCallCount() int {
	fake.tLSCACertPoolMutex.RLock()
	defer fake.tLSCACertPoolMutex.RUnlock()
	return len(fake.tLSCACertPoolArgsForCall)
}

func (fake *EndpointConfig) TLSCACertPoolReturns(result1 fab.CertPool) {
	fake.TLSCACertPoolStub = nil
	fake.tLSCACertPoolReturns = struct {
		result1 fab.CertPool
	}{result1}
}

func (fake *EndpointConfig) TLSCACertPoolReturnsOnCall(i int, result1 fab.CertPool) {
	fake.TLSCACertPoolStub = nil
	if fake.tLSCACertPoolReturnsOnCall == nil {
		fake.tLSCACertPoolReturnsOnCall = make(map[int]struct {
			result1 fab.CertPool
		})
	}
	fake.tLSCACertPoolReturnsOnCall[i] = struct {
		result1 fab.CertPool
	}{result1}
}

func (fake *EndpointConfig) TLSClientCerts() []tls.Certificate {
	fake.tLSClientCertsMutex.Lock()
	ret, specificReturn := fake.tLSClientCertsReturnsOnCall[len(fake.tLSClientCertsArgsForCall)]
	fake.tLSClientCertsArgsForCall = append(fake.tLSClientCertsArgsForCall, struct{}{})
	fake.recordInvocation("TLSClientCerts", []interface{}{})
	fake.tLSClientCertsMutex.Unlock()
	if fake.TLSClientCertsStub != nil {
		return fake.TLSClientCertsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.tLSClientCertsReturns.result1
}

func (fake *EndpointConfig) TLSClientCertsCallCount() int {
	fake.tLSClientCertsMutex.RLock()
	defer fake.tLSClientCertsMutex.RUnlock()
	return len(fake.tLSClientCertsArgsForCall)
}

func (fake *EndpointConfig) TLSClientCertsReturns(result1 []tls.Certificate) {
	fake.TLSClientCertsStub = nil
	fake.tLSClientCertsReturns = struct {
		result1 []tls.Certificate
	}{result1}
}

func (fake *EndpointConfig) TLSClientCertsReturnsOnCall(i int, result1 []tls.Certificate) {
	fake.TLSClientCertsStub = nil
	if fake.tLSClientCertsReturnsOnCall == nil {
		fake.tLSClientCertsReturnsOnCall = make(map[int]struct {
			result1 []tls.Certificate
		})
	}
	fake.tLSClientCertsReturnsOnCall[i] = struct {
		result1 []tls.Certificate
	}{result1}
}

func (fake *EndpointConfig) CryptoConfigPath() string {
	fake.cryptoConfigPathMutex.Lock()
	ret, specificReturn := fake.cryptoConfigPathReturnsOnCall[len(fake.cryptoConfigPathArgsForCall)]
	fake.cryptoConfigPathArgsForCall = append(fake.cryptoConfigPathArgsForCall, struct{}{})
	fake.recordInvocation("CryptoConfigPath", []interface{}{})
	fake.cryptoConfigPathMutex.Unlock()
	if fake.CryptoConfigPathStub != nil {
		return fake.CryptoConfigPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cryptoConfigPathReturns.result1
}

func (fake *EndpointConfig) CryptoConfigPathCallCount() int {
	fake.cryptoConfigPathMutex.RLock()
	defer fake.cryptoConfigPathMutex.RUnlock()
	return len(fake.cryptoConfigPathArgsForCall)
}

func (fake *EndpointConfig) CryptoConfigPathReturns(result1 string) {
	fake.CryptoConfigPathStub = nil
	fake.cryptoConfigPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *EndpointConfig) CryptoConfigPathReturnsOnCall(i int, result1 string) {
	fake.CryptoConfigPathStub = nil
	if fake.cryptoConfigPathReturnsOnCall == nil {
		fake.cryptoConfigPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cryptoConfigPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *EndpointConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.timeoutMutex.RLock()
	defer fake.timeoutMutex.RUnlock()
	fake.orderersConfigMutex.RLock()
	defer fake.orderersConfigMutex.RUnlock()
	fake.ordererConfigMutex.RLock()
	defer fake.ordererConfigMutex.RUnlock()
	fake.peersConfigMutex.RLock()
	defer fake.peersConfigMutex.RUnlock()
	fake.peerConfigMutex.RLock()
	defer fake.peerConfigMutex.RUnlock()
	fake.networkConfigMutex.RLock()
	defer fake.networkConfigMutex.RUnlock()
	fake.networkPeersMutex.RLock()
	defer fake.networkPeersMutex.RUnlock()
	fake.channelConfigMutex.RLock()
	defer fake.channelConfigMutex.RUnlock()
	fake.channelPeersMutex.RLock()
	defer fake.channelPeersMutex.RUnlock()
	fake.channelOrderersMutex.RLock()
	defer fake.channelOrderersMutex.RUnlock()
	fake.tLSCACertPoolMutex.RLock()
	defer fake.tLSCACertPoolMutex.RUnlock()
	fake.tLSClientCertsMutex.RLock()
	defer fake.tLSClientCertsMutex.RUnlock()
	fake.cryptoConfigPathMutex.RLock()
	defer fake.cryptoConfigPathMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EndpointConfig) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ fab.EndpointConfig = new(EndpointConfig)
