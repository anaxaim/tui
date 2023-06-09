package container

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/anaxaim/tui/backend/pkg/utils"
)

type TerraformService struct{}

func NewTerraformService() *TerraformService {
	return &TerraformService{}
}

func (t *TerraformService) PrepareWorkingDirectory(content map[string]string) (string, error) {
	tempDir, err := os.MkdirTemp("", "terraform-")
	if err != nil {
		return "", err
	}

	if err := utils.WriteFiles(tempDir, content); err != nil {
		return "", err
	}

	return tempDir, nil
}

func (t *TerraformService) isContainerRunning(ctx context.Context, containerName string) (bool, error) {
	args := []string{"ps", "-a", "--filter", "name=" + containerName, "--format", "{{.Names}}"}
	cmd := exec.CommandContext(ctx, "docker", args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return false, fmt.Errorf("failed to check if container is running: %w, stderr: %s", err, stderr.String())
	}

	return strings.TrimSpace(stdout.String()) == containerName, nil
}

func (t *TerraformService) RunContainerWithVolume(ctx context.Context, version, workingDir, id string) error {
	containerName := fmt.Sprintf("tui-terraform-%s", version)

	isRunning, err := t.isContainerRunning(ctx, containerName)
	if err != nil {
		return err
	}

	var stderr bytes.Buffer

	if isRunning {
		args := []string{"cp", workingDir, fmt.Sprintf("%s:/terraform/", containerName)}
		cmd := exec.CommandContext(ctx, "docker", args...)

		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to copy working directory to running container: %w, stderr: %s", err, stderr.String())
		}
	} else {
		volume := fmt.Sprintf("%s:/terraform/%s", workingDir, id)
		image := fmt.Sprintf("docker-compose-terraform-%s", version)

		cmd := exec.CommandContext(
			ctx,
			"docker",
			"run", "--rm", "-d", "--name", containerName, "-v", volume, image,
		)

		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to start Terraform container: %w, stderr: %s", err, stderr.String())
		}
	}

	return nil
}

func (t *TerraformService) RunContainer(ctx context.Context, version string) error {
	containerName := fmt.Sprintf("tui-terraform-%s", version)
	image := fmt.Sprintf("docker-compose-terraform-%s", version)

	isRunning, err := t.isContainerRunning(ctx, containerName)
	if err != nil {
		return err
	}

	if isRunning {
		return nil
	}

	cmd := exec.CommandContext(
		ctx,
		"docker",
		"run", "--rm", "-d", "--name", containerName, image,
	)

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to start Terraform container: %w, stderr: %s", err, stderr.String())
	}

	return nil
}

func (t *TerraformService) StopContainer(ctx context.Context, version string) error {
	containerName := "tui-terraform-" + version

	cmd := exec.CommandContext(ctx, "docker", "stop", containerName)

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to stop Terraform container: %w, stderr: %s", err, stderr.String())
	}

	cmd = exec.CommandContext(ctx, "docker", "rm", containerName)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to remove Terraform container: %w, stderr: %s", err, stderr.String())
	}

	return nil
}
