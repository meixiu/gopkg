package fs

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// stat()
func Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}

// pathinfo()
// -1: all; 1: dirname; 2: basename; 4: extension; 8: filename
// Usage:
// Pathinfo("/home/go/path/src/php2go/php2go.go", 1|2|4|8)
func Pathinfo(path string, options int) map[string]string {
	if options == -1 {
		options = 1 | 2 | 4 | 8
	}
	info := make(map[string]string)
	if (options & 1) == 1 {
		info["dirname"] = filepath.Dir(path)
	}
	if (options & 2) == 2 {
		info["basename"] = filepath.Base(path)
	}
	if ((options & 4) == 4) || ((options & 8) == 8) {
		basename := ""
		if (options & 2) == 2 {
			basename = info["basename"]
		} else {
			basename = filepath.Base(path)
		}
		p := strings.LastIndex(basename, ".")
		filename, extension := "", ""
		if p > 0 {
			filename, extension = basename[:p], basename[p+1:]
		} else if p == -1 {
			filename = basename
		} else if p == 0 {
			extension = basename[p+1:]
		}
		if (options & 4) == 4 {
			info["extension"] = extension
		}
		if (options & 8) == 8 {
			info["filename"] = filename
		}
	}
	return info
}

// file_exists()
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// is_file()
func IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// is_dir()
func IsDir(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	fm := fd.Mode()
	return fm.IsDir(), nil
}

// filesize()
func FileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

// file_put_contents()
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}

// file_get_contents()
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}

// unlink()
func Unlink(filename string) error {
	return os.Remove(filename)
}

// delete()
func Delete(filename string) error {
	return os.Remove(filename)
}

// copy()
func Copy(source, dest string) (bool, error) {
	fd1, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer fd1.Close()
	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	defer fd2.Close()
	_, e := io.Copy(fd2, fd1)
	if e != nil {
		return false, e
	}
	return true, nil
}

// is_readable()
func IsReadable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_RDONLY, 0)
	return err == nil
}

// is_writeable()
func IsWriteable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_WRONLY, 0)
	return err == nil
}

// rename()
func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// touch()
func Touch(filename string) (bool, error) {
	fd, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return false, err
	}
	fd.Close()
	return true, nil
}

// mkdir()
func Mkdir(filename string, mode os.FileMode) error {
	return os.Mkdir(filename, mode)
}

// getcwd()
func Getcwd() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}

// realpath()
func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

// basename()
func Basename(path string) string {
	return filepath.Base(path)
}

// chmod()
func Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// chown()
func Chown(filename string, uid, gid int) bool {
	return os.Chown(filename, uid, gid) == nil
}

// fclose()
func Fclose(handle *os.File) error {
	return handle.Close()
}

// filemtime()
func Filemtime(filename string) (int64, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fd.Close()
	fileinfo, err := fd.Stat()
	if err != nil {
		return 0, err
	}
	return fileinfo.ModTime().Unix(), nil
}

// fgetcsv()
func Fgetcsv(handle *os.File, length int, delimiter rune) ([][]string, error) {
	reader := csv.NewReader(handle)
	reader.Comma = delimiter
	// TODO length limit
	return reader.ReadAll()
}

// disk_free_space()
// Only supports Unix/Linux operating systems
func DiskFreeSpace(directory string) (uint64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(directory, &fs)
	if err != nil {
		return 0, err
	}
	return fs.Bfree * uint64(fs.Bsize), nil
}

// disk_total_space()
// Only supports Unix/Linux operating systems
func DiskTotalSpace(directory string) (uint64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(directory, &fs)
	if err != nil {
		return 0, err
	}
	return fs.Blocks * uint64(fs.Bsize), nil
}

// glob()
func Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

// umask()
func Umask(mask int) int {
	return syscall.Umask(mask)
}
