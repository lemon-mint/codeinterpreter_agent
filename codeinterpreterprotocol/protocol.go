package codeinterpreterprotocol

type RunCommandRequest struct {
	Command string
}

type RunCommandResponse struct {
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`

	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exitCode,omitempty"`
}

type WriteFileRequest struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

type WriteFileResponse struct {
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
}
