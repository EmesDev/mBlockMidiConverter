package services

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
	"go.bug.st/serial"
)

type connection struct {
	BaudRate   int
	SerialPort string
	connection serial.Port
}

type Connection interface {
	InitConnection() error
	UpdateSerialPort(serialPort string)
	Read(buff []byte)
}

type ConnectionPorts []string

var SerialPorts ConnectionPorts

func (c *connection) UpdateSerialPort(serialPort string) {
	c.SerialPort = serialPort
}

func (s *connection) InitConnection() error {
	port, err := serial.Open(s.SerialPort, &serial.Mode{
		BaudRate: s.BaudRate,
	})
	if err != nil {
		return err
	}

	s.connection = port

	return nil
}

func GetPorts() {
	Ports, err := serial.GetPortsList()
	if err != nil {
		panic(err)
	}
	SerialPorts = Ports
}

func NewConnection(baudRate int, serialPort string) *connection {
	return &connection{
		BaudRate:   baudRate,
		SerialPort: serialPort,
	}
}
func (c *connection) Read(buff []byte) {
	var dataBuffer string

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	for {
		n, err := c.connection.Read(buff)
		if err != nil {
			log.Println(err)
		}

		if n == 0 {
			fmt.Println("\nEOF")
		}

		dataBuffer += string(buff[:n])

		if strings.Contains(dataBuffer, "\r\n") {
			dataBuffer = strings.Replace(dataBuffer, "\r\n", "", -1)

			switch dataBuffer {
			case "10":
				kb.SetKeys(keybd_event.VK_Y)
				kb.Press()
				time.Sleep(100 * time.Millisecond)
				kb.Release()

				fmt.Println("Tocando prato")
			case "20":
				kb.SetKeys(keybd_event.VK_X)
				kb.Press()
				time.Sleep(100 * time.Millisecond)
				kb.Release()

				fmt.Println("Tocando Bumbo")
			case "30":
				kb.SetKeys(keybd_event.VK_G)
				kb.Press()
				time.Sleep(100 * time.Millisecond)
				kb.Release()

				fmt.Println("Tocando agudo")
			case "40":
				kb.SetKeys(keybd_event.VK_H)
				kb.Press()
				time.Sleep(100 * time.Millisecond)
				kb.Release()

				fmt.Println("Tocando grave")
			case "50":
				kb.SetKeys(keybd_event.VK_S)
				kb.Press()
				time.Sleep(100 * time.Millisecond)
				kb.Release()

				fmt.Println("Tocando tarola")
			case "60":
				kb.SetKeys(keybd_event.VK_E)
				kb.Press()
				time.Sleep(100 * time.Millisecond)
				kb.Release()

				fmt.Println("Tocando prato")
			}

			dataBuffer = ""
		}

	}
}
