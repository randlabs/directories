package directories

import (
	"fmt"
	"os"
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

func TestAppFile(t *testing.T) {
	SetAppName("dirtest")
	f, err := Open(AppSettings, "sample.txt", os.O_WRONLY | os.O_CREATE, 0644)
	if err == nil {
		fmt.Printf("Created sample.txt file!\n")
		_, _ = f.Write([]byte("test"))
		_ = f.Close()
	} else {
		t.Errorf("%v", err.Error())
	}
}
