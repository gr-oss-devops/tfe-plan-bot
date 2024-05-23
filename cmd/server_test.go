package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestServerCmd(t *testing.T) {
	cmd := &cobra.Command{
		Use: "tfe-plan-bot",
	}
	serverCmdConfig.Path = "nonexistent.yml"

	err := serverCmd(cmd, []string{})
	if err == nil {
		t.Errorf("Expected error due to nonexistent config file, got nil")
	}

	expectedError := "failed to read server config: failed fetching server config file: nonexistent.yml: stat nonexistent.yml: no such file or directory"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
	}
}

// TODO: fix/add tests
// func TestServerCmd_EmptyConfig(t *testing.T) {
// 	cmd := &cobra.Command{
// 		Use: "tfe-plan-bot",
// 	}
// 	serverCmdConfig.Path = "empty.yml"

// 	err := serverCmd(cmd, []string{})
// 	if err == nil {
// 		t.Errorf("Expected error due to empty config file, got nil")
// 	}

// 	expectedError := "failed to read server config: empty config file: empty.yml"
// 	if err.Error() != expectedError {
// 		t.Errorf("Expected error '%s', got '%s'", expectedError, err.Error())
// 	}
// }

// func TestServerCmd_ValidConfig(t *testing.T) {
// 	cmd := &cobra.Command{
// 		Use: "tfe-plan-bot",
// 	}
// 	serverCmdConfig.Path = "tfe-plan.yml"

// 	err := serverCmd(cmd, []string{})
// 	if err != nil {
// 		t.Errorf("Expected no error, got '%s'", err.Error())
// 	}
// }
