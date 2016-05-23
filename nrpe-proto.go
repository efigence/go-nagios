package nagios

import (
	"hash/crc32"
	"encoding/binary"
	"bytes"
	"io"
	"math/rand"
	rrand "crypto/rand"
	"time"
)

const NRPE_MAX_PACKETBUFFER_LENGTH = 1024 // this length is hardcoded in nrpe.c
var nrpeGarbage = []byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()'"[]{},./<>?`)
var nrpeGarbageCont = len(nrpeGarbage)


var randSeed int64
var randomSource *rand.Rand
func init() {
	if randSeed == 0 {
		b := make([]byte,4)
		_ , err := rrand.Read(b)
		if err != nil {
			randSeed=time.Now().UTC().UnixNano()
		} else {
			randSeed, _ =  binary.Varint(b)
		}
		randomSource = rand.New(rand.NewSource(randSeed))
	}
}



type NrpePacket struct {
	Version int16
	Type int16
	Crc uint32
	ResultCore int16
	Buffer [NRPE_MAX_PACKETBUFFER_LENGTH]byte
}

//max 1023 BYTES(not characters), it WILL truncate it if you add more
func (r *NrpePacket)SetMessage(msg string) (err error) {
	for i := range r.Buffer {
		r.Buffer[i] = nrpeGarbage[randomSource.Intn(nrpeGarbageCont)]
    }
	copy(r.Buffer[:], msg)
	r.Buffer[NRPE_MAX_PACKETBUFFER_LENGTH-1] = 0
	return err
}

// mimic nrpe randomize_buffer



// calculate crc, set version and type
func (r *NrpePacket)PrepareRequest() (err error) {
	r.Version=1
	r.Type=1
	r.Crc=0
	packet, err := r.GenerateBytes()
	if err != nil { return err }
	r.Crc=crc32.ChecksumIEEE(packet)
	return err
}

func (r *NrpePacket)Generate(w io.Writer) (err error) {
	err = binary.Write(w, binary.BigEndian, r)
	return err
}

func (r *NrpePacket)GenerateBytes() (b []byte, err error) {
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, r)

	return buf.Bytes(), err
}


func (r *NrpePacket)CheckCRC() {}