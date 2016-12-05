package nagios

import (
	"bytes"
	rrand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"math/rand"
	"time"
)

const NRPE_MAX_PACKETBUFFER_LENGTH = 1024 // this length is hardcoded in nrpe.c
const NrpePacketSize = 1036               // struct size + 2 bytes. No idea why, but nrpe C client requires it
// this is alphanumeric because original nrpe does the same thing, because they think it is "securer" for some wishy-washy reason

// Nrpe packet types
const (
	NrpeTypeQuery    = 1
	NrpeTypeResponse = 2
)

var nrpeGarbage = []byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()'"[]{},./<>?`)
var nrpeGarbageCont = len(nrpeGarbage)

var randSeed int64
var randomSource *rand.Rand

func init() {
	if randSeed == 0 {
		b := make([]byte, 4)
		_, err := rrand.Read(b)
		if err != nil {
			randSeed = time.Now().UTC().UnixNano()
		} else {
			randSeed, _ = binary.Varint(b)
		}
		randomSource = rand.New(rand.NewSource(randSeed))
	}
}

type NrpePacket struct {
	Version    int16
	Type       int16
	Crc        uint32
	ResultCore int16
	Buffer     [NRPE_MAX_PACKETBUFFER_LENGTH]byte
	Tail       [2]byte // doesn't work without it, even if common.h packet struct says it should be 1024, it only works when total buffer is 1026 bytes
}

//max 1023 BYTES(not characters), it WILL truncate it if you add more
func (r *NrpePacket) SetMessage(msg string) (err error) {
	if len(msg) >= (NRPE_MAX_PACKETBUFFER_LENGTH - 1) {
		return errors.New("Max message size exceed")
	}
	// mimic nrpe randomize_buffer
	for i := range r.Buffer {
		_ = i
		r.Buffer[i] = nrpeGarbage[randomSource.Intn(nrpeGarbageCont)]
	}
	copy(r.Buffer[:], msg)
	r.Buffer[len(msg)] = 0
	// just in case of some horribly broken implementation recieves that packet null last byte
	r.Buffer[NRPE_MAX_PACKETBUFFER_LENGTH-1] = 0
	return err
}

func ReadNrpeBytes(b []byte) (p *NrpePacket, err error) {
	// we accept bigger packets just in case of some buggy implementation. rest of packet is ignored
	if len(b) < NrpePacketSize {
		return &NrpePacket{}, fmt.Errorf("Wrong packet size %d, should be %d", len(b), NrpePacketSize)
	}
	r := bytes.NewReader(b)
	return ReadNrpe(r)
}
func ReadNrpe(r io.Reader) (p *NrpePacket, err error) {
	var nrpe NrpePacket
	err = binary.Read(r, binary.BigEndian, &nrpe)
	if err != nil {
		return &nrpe, err
	}
	// TODO checksum test
	return &nrpe, err
}

// calculate crc, set version and type
// Should be called before generating packet
func (r *NrpePacket) PrepareRequest() (err error) {
	r.Version = 2
	r.Type = 1
	r.Crc = 0
	packet, err := r.GenerateBytes()
	if err != nil {
		return err
	}
	r.Crc = crc32.ChecksumIEEE(packet)
	return err
}
func (r *NrpePacket) PrepareResponse() (err error) {
	r.Version = 2
	r.Type = 2
	r.Crc = 0
	packet, err := r.GenerateBytes()
	if err != nil {
		return err
	}
	r.Crc = crc32.ChecksumIEEE(packet)
	return err
}

func (r *NrpePacket) GetMessage() (str string, err error) {
	b := bytes.NewBuffer(r.Buffer[:])
	msg, err := b.ReadBytes('\000')
	msg = bytes.Trim(msg, "\000")
	return string(msg), err
}

func (r *NrpePacket) Generate(w io.Writer) (err error) {
	err = binary.Write(w, binary.BigEndian, r)
	return err
}

func (r *NrpePacket) GenerateBytes() (b []byte, err error) {
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, r)

	return buf.Bytes(), err
}

func (r *NrpePacket) CheckCRC() {}
