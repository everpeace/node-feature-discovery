package exec

import (
	"reflect"
	"testing"
)

func TestSource_ExecuteLableCommand_Success(t *testing.T) {
	commands := []string{"./test_cases/test_success1.sh", "./test_cases/test_success2.sh", ""}
	expected := [][]string{
		[]string{"v1", "v2"},
		[]string{"v1", "v2"},
		[]string{},
	}

	for i, _ := range commands {
		s := Source{Command: commands[i]}
		if features, err := s.Discover(); err != nil {
			t.Errorf("Error: %s", err.Error())
		} else if !reflect.DeepEqual(features, expected[i]) {
			t.Errorf("Not Expected: actual=%s, expected=%s", features, expected)
		} else {
			t.Log("success!")
		}
	}
}

func TestSource_ExecuteLableCommand_CommandError(t *testing.T) {
	s := Source{Command: "echooo"}
	if _, err := s.Discover(); err != nil {
		t.Log("success with error: %s", err.Error())
	} else {
		t.Error("error expected.")
	}

}
