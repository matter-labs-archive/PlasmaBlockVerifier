# Description

In this repository resides an initial prototype of the block verifier. The code is dirty and serves for now only for purposes of benchmarking. A verifiying client allows everyone to monitor the Plasma activity and provide fraud proofs in case of malicious operator's behavior. Block verifier should also keep various Merkle proofs for user's UTXO.

At the moment the achieved block verification speed is `~20000 TPS`. That is lower that the maximum achieved block production speed with `~30000 TPS` and should be improved, as well as quasi-sharing procedure implemented that will allow small clients to only monitor a part of their address space.

### Authors

- Alex Vlasov, [@shamatar](https://github.com/shamatar)