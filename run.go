package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"os/exec"

	"github.com/lemon-mint/codeinterpreter_agent/codeinterpreterprotocol"
)

func WriteFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func RunCommand(command string, res *codeinterpreterprotocol.RunCommandResponse) {
	cmd := exec.Command("bash", "-c", command)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	res.ExitCode = cmd.ProcessState.ExitCode()
	if err != nil {
		res.Error = err.Error()
	}
	res.Stdout = stdout.String()
	res.Stderr = stderr.String()
}

func RunCommandHandler(w http.ResponseWriter, r *http.Request) {
	var req codeinterpreterprotocol.RunCommandRequest
	var res codeinterpreterprotocol.RunCommandResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	RunCommand(req.Command, &res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteFileHandler(w http.ResponseWriter, r *http.Request) {
	var req codeinterpreterprotocol.WriteFileRequest
	var res codeinterpreterprotocol.WriteFileResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := WriteFile(req.Filename, req.Content); err != nil {
		res.Success = false
		res.Error = err.Error()
	}

	res.Success = true
	res.Error = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
