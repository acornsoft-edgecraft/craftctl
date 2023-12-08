package common

type NodeType string
type State string

const (
	PASS State = "PASS"
	FAIL State = "FAIL"
	WARN State = "WARN"
	INFO State = "INFO"

	MASTER       NodeType = "master"
	NODE         NodeType = "node"
	ETCD         NodeType = "etcd"
	CONTROLPLANE NodeType = "controlplane"
	POLICIES     NodeType = "policies"
)
