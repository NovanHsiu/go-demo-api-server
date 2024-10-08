package common

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	uCipher "github.com/NovanHsiu/go-demo-api-server/internal/domain/cipher"
)

const randomText = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var Cipher = uCipher.NewCipher(24, 11, "This is a private key!")
var ImageExtName = []string{".jpg", ".jpeg", ".png", ".bmp", ".gif"}

// ContainsString Contains find element slice contains the element or not
func ArrayContainsString(arr []string, v string) bool {
	return matchArrayString(arr, v, false)
}

// HasString find the exactly same string in array
func ArrayHasString(arr []string, v string) bool {
	return matchArrayString(arr, v, true)
}

func matchArrayString(arr []string, v string, isExact bool) bool {
	for _, arrv := range arr {
		if isExact {
			if v == arrv {
				return true
			}
		} else {
			if strings.Contains(arrv, v) {
				return true
			}
		}
	}
	return false
}

func GetExecutionDir() string {
	exdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return exdir
}

// GetRandomString 取得指定長度的隨機字串，長度必需大於0小於等於100，否則會回傳空字串 ""
func GetRandomString(length int) (rndtxt string) {
	rand.Seed(time.Now().UnixNano())
	if length < 1 || length > 100 {
		return rndtxt
	}
	for i := 0; i < length; i++ {
		rndtxt += string(randomText[rand.Intn(len(randomText))])
	}
	return rndtxt
}

// GetRandomTempDirName 取得指定長度且不重複的暫存資料夾路徑
func GetRandomTempDirName(basePath string, length int) (rndtxt string) {
	rndtxt = GetRandomString(length)
	for j := 0; j < 10; j++ {
		if _, err := os.Stat(basePath + "/" + rndtxt); os.IsNotExist(err) {
			break
		}
		time.Sleep(10 * time.Millisecond)
		rndtxt = GetRandomString(length)
	}
	return basePath + "/" + rndtxt
}

// CopyFile copy a file from source to destination path
func CopyFile(src, dst string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()
	to, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer to.Close()
	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}
	return nil
}

// MatchDatePattern check date string is match pattern "YYYY-MM-DD"，ex:"2015-11-26"
func MatchDatePattern(date string) bool {
	var validDate = regexp.MustCompile(`^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$`)
	return validDate.MatchString(date)
}

// ToRocDate change date to roc year format, ex: 1951-12-11 -> 401211
func ToRocDate(date string) string {
	dateSplit := strings.Split(date, "-")
	adyear, _ := strconv.Atoi(dateSplit[0])
	dateSplit[0] = strconv.Itoa(adyear - 1911)
	return strings.Join(dateSplit, "")
}

// RocToADYear ROC YYYMMDD(or YYMMDD) format to AD YYYY-MM-DD format
func RocToADYear(rocDate string) string {
	var yeindex int
	if len(rocDate) == 6 {
		yeindex = 2
	} else if len(rocDate) == 7 {
		yeindex = 3
	} else {
		return "0000-00-00"
	}
	rocyint, _ := strconv.Atoi(rocDate[:yeindex])
	year := strconv.Itoa(rocyint + 1911)
	month := rocDate[yeindex : yeindex+2]
	day := rocDate[yeindex+2 : yeindex+4]
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}

func Execute(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	output := string(out[:])
	return output, err
}

func ExecuteBackground(cmdstr string, args ...string) (int, error) {
	cmd := exec.Command(cmdstr, args...)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		return 0, err
	}
	go func() {
		err = cmd.Wait()
		if err != nil {
			log.Println("ExecuteBackground error: " + cmdstr + fmt.Sprintf(" %v"+err.Error()))
		}
	}()
	return cmd.Process.Pid, nil
}

// GetMacAddress get mac address of this device
func GetMacAddress() string {
	// for windows
	if runtime.GOOS == "windows" {
		stdout, _ := Execute("getmac")
		var macAddress string
		lineSplit := strings.Split(stdout, "\n")
		for _, line := range lineSplit {
			if strings.Contains(line, "Device") {
				macAddress = strings.Split(line, " ")[0]
				break
			}
		}
		return macAddress
	}
	// for liniux
	stdout, _ := Execute("cat", "/sys/class/net/eth0/address")
	return strings.Replace(stdout, "\n", "", 1)
}

// IsImage check filename is image
func IsImage(filename string) bool {
	lfn := strings.ToLower(filename)
	for i := range ImageExtName {
		if strings.HasSuffix(lfn, ImageExtName[i]) {
			return true
		}
	}
	return false
}
