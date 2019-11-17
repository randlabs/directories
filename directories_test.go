package directories

import (
	"fmt"
	"testing"
)

//------------------------------------------------------------------------------

func TestGetHomeDirectory(t *testing.T) {
	dir, err := GetHomeDirectory()
	if err == nil {
		fmt.Printf("Home directory: %v\n", dir)
	} else {
		t.Errorf("%v", err.Error())
	}
}

func TestGetAppSettingsDirectory(t *testing.T) {
	dir, err := GetAppSettingsDirectory()
	if err == nil {
		fmt.Printf("Application settings directory: %v\n", dir)
	} else {
		t.Errorf("%v", err.Error())
	}
}

func TestGetSystemSettingsDirectory(t *testing.T) {
	dir, err := GetSystemSettingsDirectory()
	if err == nil {
		fmt.Printf("System settings directory: %v\n", dir)
	} else {
		t.Errorf("%v", err.Error())
	}
}
