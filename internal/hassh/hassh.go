package hassh

import (
	"crypto/md5"
	"net"
	"strings"
	"sync"

	"github.com/edutko/hassh-go/x/crypto/ssh"
)

type SSHInfo struct {
	Banner  string
	Kexinit ssh.KexinitInfo
}

func Digest(info ssh.KexinitInfo) []byte {
	kex := strings.Join(info.KexAlgos, ",")
	cphr := strings.Join(info.PeerCiphers, ",")
	macs := strings.Join(info.PeerMACs, ",")
	comp := strings.Join(info.PeerCompression, ",")

	all := strings.Join([]string{kex, cphr, macs, comp}, ";")
	d := md5.Sum([]byte(all))

	return d[:]
}

func Connect(host, port string) (SSHInfo, error) {
	info := SSHInfo{}
	connected := false
	l := sync.Mutex{}

	config := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		BannerCallback: func(b string) error {
			l.Lock()
			defer l.Unlock()
			info.Banner = b
			return nil
		},
		KexinitCallback: func(ki ssh.KexinitInfo) {
			l.Lock()
			defer l.Unlock()
			info.Kexinit = ki
			connected = true
		},
	}

	conn, err := ssh.Dial("tcp", net.JoinHostPort(host, port), config)
	if err == nil {
		_ = conn.Close()
	}

	if connected {
		return info, nil
	}
	return info, err
}
