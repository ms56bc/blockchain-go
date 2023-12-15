# blockchain-go
A blockchain implementation in go for learning
Blockchain
Blockchain is a distributed ledger of transactions that is replicated across multiple nodes in a peer-to-peer network.

Cryptographic Hash Functions
Hash functions are considered as cryptographic hash functions if they have the following properties:

Deterministic: The same input will always result in the same output.
intractable to reverse: It should be impossible to go from the output to the input.
Collision resistant: It should be impossible to find two different inputs that result in the same output.
small change in input should result in a completely different output
computationally efficient: It should be easy to compute the hash value for any given input.
Proof of Work
Proof of work is a piece of data which is difficult (costly, time-consuming) to produce but easy for others to verify and which satisfies certain requirements. - the idea is to make doing an operation computationally expensive, for instance since the hash is random number we - could ask the sender of email to send only the emails with hash value less than say 1000, since hash are random the - sender has to do a lot of trial and error, by adding extra data to find the right hash value, and hence it very - computationally expensive This is what proof of work is: the work is proven by displaying a particular hash,