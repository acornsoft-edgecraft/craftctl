package model

import "github.com/edgecraft/edge-benchmarks/pkg/common"

// kube-bench output
type KbOutput struct {
	Controls []Control `json:"controls"`
	Totals   Total     `json:"totals"`
}

type Control struct {
	ID              string          `json:"id"`
	Version         string          `json:"version"`
	DetectedVersion string          `json:"detected_version"`
	Text            string          `json:"text"`
	NodeType        common.NodeType `json:"node_type"`
	Tests           []Test          `json:"tests"`
}

type Total struct {
	Pass int `json:"total_pass"`
	Fail int `json:"total_fail"`
	Warn int `json:"total_warn"`
	Info int `json:"total_info"`
}

type Test struct {
	ID      string   `json:"section"`
	Type    string   `json:"type"`
	Pass    int      `json:"pass"`
	Fail    int      `json:"fail"`
	Warn    int      `json:"warn"`
	Info    int      `json:"info"`
	Desc    string   `json:"desc"`
	Results []Result `json:"results"`
}

type Result struct {
	ID             string       `json:"test_number"`
	Desc           string       `json:"test_desc"`
	Audit          string       `json:"audit"`
	Type           string       `json:"type"`
	Remediation    string       `json:"remediation"`
	TestInfo       []string     `json:"test_info"`
	Status         common.State `json:"status"`
	ActualValue    string       `json:"actual_value"`
	Scored         bool         `json:"scored"`
	IsMultiple     bool         `json:"IsMultiple"`
	ExpectedResult string       `json:"expected_result"`
	Reason         string       `json:"reason"`
}

// Results
type Results struct {
	Version         string `json:"version"`
	DetectedVersion string `json:"detected_version"`
	Nodes           []Node `json:"nodes"`
	Totals          Total  `json:"totals"`
}

type Node struct {
	NodeName string `json:"node_name"`
	KbOutput
}
