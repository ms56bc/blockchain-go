# Peer Discovery
We are going to use static peering for bootstrap and gossip protocol of state sharing. 
For gossip protocols, lightweight protocols are often preferred, such as UDP. UDP is a connectionless and lightweight transport protocol that can be suitable for gossip protocols. 
It allows for quick communication between peers without the overhead of establishing and maintaining connections.

## NetworkDiscoveryNode

The `BlockchainNetworkDiscovery` interface defines methods for discovering peers, retrieving network peers, and processing incoming messages.

```
type BlockchainNetworkDiscovery interface {
    DiscoverPeers(ctx context.Context)
    GetNetworkPeers() []model.Peer
    ProcessMessage(message *NodeUpdate)
}
```
- `ProcessMessage`: The ProcessMessage method of BlockchainNode processes incoming node update messages. It checks the source and performs add or remove operations accordingly.
- `DiscoverPeers`:The DiscoverPeers method initiates the gossip mechanism to periodically broadcast the node's information to a random subset of other nodes.

### PeerDiscoveryClient
The PeerDiscoveryClient interface defines method for sending discovered peers to other nodes.
```
type PeerDiscoveryClient interface {
        SendDiscoveredPeers(sourceId string, destination model.Peer, discoveredPeers []model.Peer) error
        Close() error
    }
```

### PeerDiscoveryServer
The `PeerDiscoveryServer` interface defines a method for handling incoming messages from peers.
```
    type PeerDiscoveryServer interface {
        HandleIncomingMessage(conn io.Reader) error
    }
```