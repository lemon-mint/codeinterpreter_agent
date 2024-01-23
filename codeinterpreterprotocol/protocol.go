package codeinterpreterprotocol

type RunCommandRequest struct {
	Command string `json:"command" description:"The command to run in bash (e.g. 'python3 main.py')" required:"true"`
}

type RunCommandResponse struct {
	Success bool   `json:"success,omitempty" description:"Whether the command was successful"`
	Error   string `json:"error,omitempty" description:"If the command was not successful, the error message"`

	Stdout   string `json:"stdout" description:"The standard output of the command"`
	Stderr   string `json:"stderr" description:"The standard error of the command"`
	ExitCode int    `json:"exitCode,omitempty" description:"The exit code of the command"`
}

type WriteFileRequest struct {
	Filename string `json:"filename" description:"The filename to write to (e.g. 'main.py')" required:"true"`
	Content  string `json:"content" description:"The content to write to the file (e.g. '#main.py\nimport pandas as pd\nimport numpy as np\n')" required:"true"`
}

type WriteFileResponse struct {
	Success bool   `json:"success,omitempty" description:"Whether the file was written successfully"`
	Error   string `json:"error,omitempty" description:"If the file was not written successfully, the error message"`
}
