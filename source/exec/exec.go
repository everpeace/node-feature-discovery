package exec

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Source struct {
	Command string
}

func (s Source) Name() string { return "exec" }

func (s Source) Discover() ([]string, error) {
	out, err := exec.Command("sh", "-c", s.Command).Output()

	if err != nil {
		return nil, fmt.Errorf("error occured in executing %s: %s", []string{"sh", "-c", s.Command}, err.Error())
	}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	features := make([]string, 0)
	for scanner.Scan() {
		if feature := strings.TrimSpace(scanner.Text()); feature != "" {
			features = append(features, feature)
		}
	}

	return features, nil
}
