package nagios

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestStatus(t *testing.T) {
	file, err := os.Open("t-data/status.dat.local")
	if err != nil {
		file, err = os.Open("t-data/status.dat")
		if err != nil {
			t.Logf("%s", err)
			t.FailNow()
		}
	}
	Convey("Load status.dat", t, func() {
		s, err := LoadStatus(file)
		_ = err
		Convey("Parsing", func() {
			So(err, ShouldEqual, nil)
			So(len(s.Host), ShouldNotEqual, 0)
			So(len(s.Service), ShouldNotEqual, 0)
		})
	})
	_ = fmt.Sprintf(`dummy`)

	if err != nil {
		t.Logf("%s", err)
		t.FailNow()
	}
}

func BenchmarkStatus(b *testing.B) {
	file, err := os.Open("t-data/status.local.dat")
	if err != nil {
		file, err = os.Open("t-data/status.dat")
		if err != nil {
			b.Logf("%s", err)
			b.FailNow()
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file.Seek(0, 0)
		LoadStatus(file)
	}
}
