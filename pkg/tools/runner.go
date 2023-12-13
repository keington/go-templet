package tools

import (
	"go.uber.org/automaxprocs/maxprocs"
	"hash/crc32"
	"log"
	"math/rand"
	"os"
	"time"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/10/30 22:52
 * @file: runner.go
 * @description:
 */

var (
	Hostname string
	Cwd      string
)

func Noop(string, ...interface{}) {}

func Init() {

	_, _ = maxprocs.Set(maxprocs.Logger(Noop))

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var err error
	Hostname, err = os.Hostname()
	if err != nil {
		log.Fatalln("[F] cannot get hostname")
	}

	Cwd = SelfDir()

	rand.NewSource(time.Now().UnixNano() + int64(os.Getpid()+os.Getppid()) + int64(crc32.ChecksumIEEE([]byte(Hostname))))
}
