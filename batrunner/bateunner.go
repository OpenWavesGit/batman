// batrunner.go
package batrunner

import (
    "fmt"
    "os/exec"
)

// RunBatchFile executes a batch file and returns its output or an error
func RunBatchFile(batchFilePath string) (string, error) {
    // Command to execute the batch file
    cmd := exec.Command("cmd", "/c", batchFilePath)

    // Execute the command
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", fmt.Errorf("error executing command: %v", err)
    }

    return string(output), nil
}
