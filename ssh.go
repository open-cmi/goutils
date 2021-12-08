package goutils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SSHServer ssh server
type SSHServer struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	ConnType   string `json:"conntype"`
	Password   string `json:"password"`
	SecretFile string `json:"secretfile"`
}

// NewSSHServer new ssh server
func NewSSHServer(host string, port int, conntype string, user string, password string, secretfile string) *SSHServer {
	var server SSHServer = SSHServer{
		User:       user,
		Password:   password,
		Host:       host,
		ConnType:   conntype,
		Port:       port,
		SecretFile: secretfile,
	}
	return &server
}

// SSHConnect ssh connect
func (s *SSHServer) SSHConnect() (*ssh.Client, error) {
	var (
		auth         []ssh.AuthMethod = []ssh.AuthMethod{}
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)
	fmt.Println(s)
	// get auth method
	if s.ConnType == "password" {
		auth = append(auth, ssh.Password(s.Password))
	} else {
		key, err := ioutil.ReadFile(s.SecretFile)
		if err != nil {
			return nil, errors.New("secret file read failed")
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, errors.New("parse private key failed")
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User:            s.User,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", s.Host, s.Port)

	client, err = ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// SSHRun run ssh
func (s *SSHServer) SSHRun(cmd string) error {

	var stdOut, stdErr bytes.Buffer

	client, err := s.SSHConnect()
	if err != nil {
		fmt.Printf("connect server failed: %s\n", err.Error())
		return err
	}
	defer client.Close()

	// create session
	session, err := client.NewSession()
	if err != nil {
		return err
	}

	session.Stdout = &stdOut
	session.Stderr = &stdErr

	err = session.Run(cmd)
	if err != nil {
		fmt.Printf("remote server run command failed: %s\n", err.Error())
		return err
	}
	if stdErr.String() != "" {
		return errors.New(stdErr.String())
	}
	fmt.Printf("%s\n", stdOut.String())
	return nil
}

// SSHCopyToRemote ssh copy from local to remote
func (s *SSHServer) SSHCopyToRemote(local string, remote string) error {

	client, err := s.SSHConnect()
	if err != nil {
		fmt.Printf("connect server failed: %s\n", err.Error())
		return err
	}
	defer client.Close()

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		fmt.Printf("sftp new client failed: %s\n", err.Error())
		return err
	}

	defer sftpClient.Close()

	//获取当前目录
	if strings.HasPrefix(remote, "./") {
		cwd, _ := sftpClient.Getwd()
		remote = sftp.Join(cwd, remote)
	}
	//上传文件,如果远端为文件夹，还需要重新拼接remote文件
	remoteFile, err := sftpClient.Create(remote)
	if err != nil {
		fmt.Printf("remote server create failed: %s\n", err.Error())
		return err
	}
	defer remoteFile.Close()

	//打开本地文件file.dat
	localFile, err := os.Open(local)
	if err != nil {
		fmt.Printf("open local file failed: %s\n", err.Error())
		return err
	}
	defer localFile.Close()

	//本地文件流拷贝到上传文件流
	n, err := io.Copy(remoteFile, localFile)
	if err != nil {
		fmt.Printf("copy from local to remote failed: %s\n", err.Error())
		return err
	}

	//获取本地文件大小
	localFileInfo, err := os.Stat(local)
	if err != nil {
		return err
	}

	if localFileInfo.Size() != n {
		return errors.New("copy from local to remote failed")
	}
	return nil
}

// SSHCopyToLocal copy to local
func (s *SSHServer) SSHCopyToLocal(remote string, local string) error {
	client, err := s.SSHConnect()
	if err != nil {
		fmt.Printf("connect server failed: %s\n", err.Error())
		return err
	}
	defer client.Close()

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		fmt.Printf("sftp new client failed: %s\n", err.Error())
		return err
	}

	defer sftpClient.Close()

	//获取当前目录
	if strings.HasPrefix(remote, "./") {
		cwd, _ := sftpClient.Getwd()
		remote = sftp.Join(cwd, remote)
	}

	//下载文件
	remoteFile, err := sftpClient.Open(remote)
	if err != nil {
		fmt.Printf("open remote file failed: %s\n", err.Error())
		return err
	}
	defer remoteFile.Close()

	localFile, err := os.Create(local)
	if err != nil {
		fmt.Printf("create local file failed: %s\n", err.Error())
		return err
	}
	defer localFile.Close()
	n, err := io.Copy(localFile, remoteFile)
	if err != nil {
		fmt.Printf("copy from local to remote failed: %s\n", err.Error())
		return err
	}

	//获取远程文件大小
	remoteFileInfo, err := sftpClient.Stat(remote)
	if err != nil {
		fmt.Printf("remote file stat failed: %s\n", err.Error())
		return err
	}

	if n != remoteFileInfo.Size() {
		fmt.Printf("copy from remote to local failed: %s\n", err.Error())
		return err
	}
	return nil
}
