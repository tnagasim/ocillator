package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootCmd_VersionFlag(t *testing.T) {
	cmd := newRootCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetArgs([]string{"--version"})

	err := cmd.Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "v0.0.1") {
		t.Errorf("expected version output to contain 'v0.0.1', got: %q", buf.String())
	}
}

func TestRootCmd_HelpListsSubcommands(t *testing.T) {
	cmd := newRootCmd()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetArgs([]string{"--help"})

	err := cmd.Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	help := buf.String()
	for _, sub := range []string{"sync", "release", "deploy"} {
		if !strings.Contains(help, sub) {
			t.Errorf("expected help to list subcommand %q, got:\n%s", sub, help)
		}
	}
}

func TestSyncCmd_RunsWithoutError(t *testing.T) {
	cmd := newRootCmd()
	cmd.SetArgs([]string{"sync"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("sync subcommand returned unexpected error: %v", err)
	}
}

func TestReleaseCmd_RunsWithoutError(t *testing.T) {
	cmd := newRootCmd()
	cmd.SetArgs([]string{"release"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("release subcommand returned unexpected error: %v", err)
	}
}

func TestDeployCmd_RunsWithoutError(t *testing.T) {
	cmd := newRootCmd()
	cmd.SetArgs([]string{"deploy"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("deploy subcommand returned unexpected error: %v", err)
	}
}

func TestRootCmd_UnknownSubcommandReturnsError(t *testing.T) {
	cmd := newRootCmd()
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.SetArgs([]string{"nonexistent"})

	if err := cmd.Execute(); err == nil {
		t.Fatal("expected error for unknown subcommand, got nil")
	}
}
