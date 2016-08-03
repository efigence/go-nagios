package nagios

import (
	"bytes"
	rrand "crypto/rand"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"
	"math/rand"
	"time"
)

const NRPE_MAX_PACKETBUFFER_LENGTH = 1024 // this length is hardcoded in nrpe.c
// this is alphanumeric because original nrpe does the same thing, because they think it is "securer" for some wishy-washy reason

// Nrpe packet types
const (
	NrpeTypeQuery = 1
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
}
// example from check_nrpe:
// 0000   00 02 00 01 22 c6 9f f4 38 69 63 68 65 63 6b 5f  ...."...8icheck_
// 0010   64 69 73 6b 00 00 00 00 00 00 00 00 00 00 00 00  disk............
// 0020   00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
// 0030   00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................

//max 1023 BYTES(not characters), it WILL truncate it if you add more
func (r *NrpePacket) SetMessage(msg string) (err error) {
	if len(msg) >= (NRPE_MAX_PACKETBUFFER_LENGTH - 1) {
		return errors.New("Max message size exceed")
	}
	for i := range r.Buffer {
		r.Buffer[i] = nrpeGarbage[randomSource.Intn(nrpeGarbageCont)]
	}
	copy(r.Buffer[:], msg)
	r.Buffer[len(msg)] = 0
	// just in case of some horribly broken implementation recieves that packet null last byte
	r.Buffer[NRPE_MAX_PACKETBUFFER_LENGTH-1] = 0
	return err
}

// mimic nrpe randomize_buffer

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
