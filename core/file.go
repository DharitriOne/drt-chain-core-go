package core

import (
	"bytes"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pelletier/go-toml"
)

const pemPkHeader = "PRIVATE KEY for "

// ArgCreateFileArgument will hold the arguments for a new file creation method call
type ArgCreateFileArgument struct {
	Directory     string
	Prefix        string
	FileExtension string
}

// OpenFile method opens the file from given path - does not close the file
func OpenFile(relativePath string) (*os.File, error) {
	path, err := filepath.Abs(relativePath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	return f, nil
}

// LoadTomlFile method to open and decode toml file
func LoadTomlFile(dest interface{}, relativePath string) error {
	f, err := OpenFile(relativePath)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	return toml.NewDecoder(f).Decode(dest)
}

// SaveTomlFile will open and save data to toml file
func SaveTomlFile(src interface{}, relativePath string) error {
	f, err := os.Create(relativePath)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	return toml.NewEncoder(f).Encode(src)
}

// LoadTomlFileToMap opens and decodes a toml file as a map[string]interface{}
func LoadTomlFileToMap(relativePath string) (map[string]interface{}, error) {
	f, err := OpenFile(relativePath)
	if err != nil {
		return nil, err
	}

	fileinfo, err := f.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = f.Read(buffer)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	loadedTree, err := toml.Load(string(buffer))
	if err != nil {
		return nil, err
	}

	loadedMap := loadedTree.ToMap()

	return loadedMap, nil
}

// LoadJsonFile method to open and decode json file
func LoadJsonFile(dest interface{}, relativePath string) error {
	f, err := OpenFile(relativePath)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	return json.NewDecoder(f).Decode(dest)
}

// CreateFile opens or creates a file relative to the default path
func CreateFile(arg ArgCreateFileArgument) (*os.File, error) {
	absPath, err := filepath.Abs(arg.Directory)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(absPath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	fileName := time.Now().Format("2006-01-02T15-04-05")
	fileName = strings.Replace(fileName, "T", "-", 1)
	if arg.Prefix != "" {
		fileName = arg.Prefix + "-" + fileName
	}

	return os.OpenFile(
		filepath.Join(absPath, fileName+"."+arg.FileExtension),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		FileModeUserReadWrite)
}

// LoadSkPkFromPemFile loads the secret key and existing public key bytes stored in the file
func LoadSkPkFromPemFile(relativePath string, skIndex int) ([]byte, string, error) {
	if skIndex < 0 {
		return nil, "", ErrInvalidIndex
	}

	file, err := OpenFile(relativePath)
	if err != nil {
		return nil, "", err
	}

	defer func() {
		_ = file.Close()
	}()

	buff, err := io.ReadAll(file)
	if err != nil {
		return nil, "", fmt.Errorf("%w while reading %s file", err, relativePath)
	}
	if len(buff) == 0 {
		return nil, "", fmt.Errorf("%w while reading %s file", ErrEmptyFile, relativePath)
	}

	var blkRecovered *pem.Block

	for i := 0; i <= skIndex; i++ {
		if len(buff) == 0 {
			//less private keys present in the file than required
			return nil, "", fmt.Errorf("%w while reading %s file, invalid index %d", ErrInvalidIndex, relativePath, i)
		}

		blkRecovered, buff = pem.Decode(buff)
		if blkRecovered == nil {
			return nil, "", fmt.Errorf("%w while reading %s file, error decoding", ErrPemFileIsInvalid, relativePath)
		}
	}

	if blkRecovered == nil {
		return nil, "", ErrNilPemBLock
	}

	blockType := blkRecovered.Type

	if strings.Index(blockType, pemPkHeader) != 0 {
		return nil, "", fmt.Errorf("%w missing '%s' in block type", ErrPemFileIsInvalid, pemPkHeader)
	}

	blockTypeString := blockType[len(pemPkHeader):]

	return blkRecovered.Bytes, blockTypeString, nil
}

// LoadAllKeysFromPemFile loads all the secret keys and existing public key bytes stored in the file
func LoadAllKeysFromPemFile(relativePath string) ([][]byte, []string, error) {
	file, err := OpenFile(relativePath)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	buff, err := io.ReadAll(file)
	if err != nil {
		return nil, nil, fmt.Errorf("%w while reading %s file", err, relativePath)
	}
	if len(buff) == 0 {
		return nil, nil, fmt.Errorf("%w while reading %s file", ErrEmptyFile, relativePath)
	}

	var blkRecovered *pem.Block
	privateKeys := make([][]byte, 0)
	publicKeys := make([]string, 0)

	for {
		if len(buff) == 0 {
			break
		}

		blkRecovered, buff = pem.Decode(buff)
		if blkRecovered == nil {
			return nil, nil, fmt.Errorf("%w while reading %s file, error decoding", ErrPemFileIsInvalid, relativePath)
		}
		buff = bytes.TrimSpace(buff)

		blockType := blkRecovered.Type
		if strings.Index(blockType, pemPkHeader) != 0 {
			return nil, nil, fmt.Errorf("%w missing '%s' in block type", ErrPemFileIsInvalid, pemPkHeader)
		}

		blockTypeString := blockType[len(pemPkHeader):]

		privateKeys = append(privateKeys, blkRecovered.Bytes)
		publicKeys = append(publicKeys, blockTypeString)
	}

	return privateKeys, publicKeys, nil
}

// SaveSkToPemFile saves secret key bytes in the file
func SaveSkToPemFile(file *os.File, identifier string, skBytes []byte) error {
	if file == nil {
		return ErrNilFile
	}

	blk := pem.Block{
		Type:  "PRIVATE KEY for " + identifier,
		Bytes: skBytes,
	}

	return pem.Encode(file, &blk)
}
