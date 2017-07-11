package main

import (
	"net"
	"time"
	"vncproxy/client"
	"vncproxy/common"
	"vncproxy/encodings"
	"vncproxy/logger"
	listeners "vncproxy/tee-listeners"
)

func main() {

	//nc, err := net.Dial("tcp", "192.168.1.101:5903")
	nc, err := net.Dial("tcp", "localhost:5903")

	if err != nil {
		logger.Errorf("error connecting to vnc server: %s", err)
	}
	var noauth client.ClientAuthNone
	authArr := []client.ClientAuth{&client.PasswordAuth{Password: "Ch_#!T@8"}, &noauth}

	vncSrvMessagesChan := make(chan common.ServerMessage)

	rec := listeners.NewRecorder("c:/Users/betzalel/recording.rbs")

	clientConn, err := client.NewClientConn(nc,
		&client.ClientConfig{
			Auth:            authArr,
			ServerMessageCh: vncSrvMessagesChan,
			Exclusive:       true,
		})

	clientConn.Listeners.AddListener(rec)
	clientConn.Connect()

	if err != nil {
		logger.Errorf("error creating client: %s", err)
	}
	// err = clientConn.FramebufferUpdateRequest(false, 0, 0, 1024, 768)
	// if err != nil {
	// 	logger.Errorf("error requesting fb update: %s", err)
	// }

	tight := encodings.TightEncoding{}
	tightPng := encodings.TightPngEncoding{}
	//rre := encodings.RREEncoding{}
	//zlib := encodings.ZLibEncoding{}
	//zrle := encodings.ZRLEEncoding{}
	cpyRect := encodings.CopyRectEncoding{}
	//coRRE := encodings.CoRREEncoding{}
	//hextile := encodings.HextileEncoding{}
	// file, _ := os.OpenFile("stam.bin", os.O_CREATE|os.O_RDWR, 0755)
	// defer file.Close()

	//tight.SetOutput(file)
	clientConn.SetEncodings([]common.Encoding{&cpyRect, &tightPng, &tight})

	go func() {
		for {
			err = clientConn.FramebufferUpdateRequest(true, 0, 0, 1280, 800)
			if err != nil {
				logger.Errorf("error requesting fb update: %s", err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	//go func() {
	for msg := range vncSrvMessagesChan {
		logger.Debugf("message type: %d, content: %v\n", msg.Type(), msg)
	}
	//}()

	//clientConn.Close()
}
