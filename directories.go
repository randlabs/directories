package directories

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

//------------------------------------------------------------------------------

type Location int

const (
	Home           Location = 1
	AppSettings             = 2
	SystemSettings          = 3
)

var appName string
var pathSep = string(filepath.Separator)

//------------------------------------------------------------------------------

// GetHomeDirectory gets the user's home directory
func GetHomeDirectory() (string, error) {
	d, err := getHomeDirectory()
	if err == nil {
		if !strings.HasSuffix(d, pathSep) {
			d += pathSep
		}
	}
	return d, err
}

// GetAppSettingsDirectory gets the directory to save local app settings.
// Should be used for non system-wide applications.
func GetAppSettingsDirectory() (string, error) {
	d, err := getAppSettingsDirectory()
	if err == nil {
		if !strings.HasSuffix(d, pathSep) {
			d += pathSep
		}
		if len(appName) > 0 {
			d += appName + pathSep
		}
	}
	return d, err
}

// GetSystemSettingsDirectory gets the directory to save app settings.
// Should be used for system-wide applications.
func GetSystemSettingsDirectory() (string, error) {
	d, err := getSystemSettingsDirectory()
	if err == nil {
		if !strings.HasSuffix(d, pathSep) {
			d += pathSep
		}
		if len(appName) > 0 {
			d += appName + pathSep
		}
	}
	return d, err
}

// SetAppName sets the name of the current application.
// GetAppSettingsDirectory and GetSystemSettingsDirectory appends the application
// name automatically to the returned string.
func SetAppName(name string) {
	appName = name
	return
}

// BuildFilename returns the complete file name using the desired location as a base.
func BuildFilename(location Location, name string) (string, error) {
	var dir string
	var err error

	//check name
	if pathSep == "\\" {
		name = strings.ReplaceAll(name, "/", pathSep)
	} else {
		name = strings.ReplaceAll(name, "\\", pathSep)
	}

	for strings.HasPrefix(name, pathSep) {
		name = name[1:]
	}
	if len(name) == 0 {
		return "", errors.New("invalid file name")
	}
	if strings.HasSuffix(name, pathSep) {
		return "", errors.New("cannot open a directory")
	}

	//get directory
	switch location {
	case Home:
		dir, err = GetHomeDirectory()
		if err != nil {
			return "", err
		}

	case AppSettings:
		dir, err = GetAppSettingsDirectory()
		if err != nil {
			return "", err
		}

	case SystemSettings:
		dir, err = GetSystemSettingsDirectory()
		if err != nil {
			return "", err
		}

	default:
		return "", errors.New("invalid directory location")
	}

	// done
	return dir + name, nil
}

// Open is a shortcut to open and/or create a file under the desired location.
func Open(location Location, name string, flag int, perm os.FileMode) (*os.File, error) {
	var err error

	//build complete name
	name, err = BuildFilename(location, name)
	if err != nil {
		return nil, err
	}

	//create parent directories if they don't exist
	if (flag & os.O_CREATE) != 0 {
		err = os.MkdirAll(filepath.Dir(name), perm)
		if err != nil {
			return nil, err
		}
	}

	//open/create file
	return os.OpenFile(name, flag, perm)
}
