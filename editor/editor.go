package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"

	"github.com/vorstenbosch/templator/yamldata"
)

//
// Heavily inspired from https://samrapdev.com/capturing-sensitive-input-with-editor-in-golang-from-the-cli/
//

// openFileInEditor opens filename in a text editor.
func openFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		if runtime.GOOS == "windows" {
			editor = "notepad.exe"
		} else {
			editor = "vim"
		}
	}

	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// CaptureInputFromEditor opens a temporary file in a text editor and returns
// the written bytes on success or an error on failure. It handles deletion
// of the temporary file behind the scenes.
func CaptureInputFromEditor(d map[string]string) ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()

	// Putting the prepared YAML in the file
	b, err := yamldata.PrepareYamlFromData(d)

	if err != nil {
		return []byte{}, err
	}

	err = ioutil.WriteFile(filename, b, 0777)
	if err != nil {
		return []byte{}, err
	}

	// Defer removal of the temporary file in case any of the next steps fail.
	defer os.Remove(filename)

	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	if err = openFileInEditor(filename); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
