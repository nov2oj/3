package engine

import (
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
	"reflect"
)

type magnetization struct {
	buffered
}

func (b *magnetization) Set(src *data.Slice) {
	if src.Mesh().Size() != b.buffer.Mesh().Size() {
		src = data.Resample(src, b.buffer.Mesh().Size())
	}
	data.Copy(b.buffer, src)
	cuda.Normalize(b.buffer, vol)
}

func (b *magnetization) LoadFile(fname string) {
	b.Set(LoadFile(fname))
}

// Sets the magnetization inside the shape
// TODO: a bit slowish
func (m *magnetization) SetInShape(region Shape, conf Config) {
	if region == nil {
		region = universe
	}
	host := m.buffer.HostCopy()
	h := host.Vectors()
	n := m.Mesh().Size()
	c := m.Mesh().CellSize()
	dx := (float64(n[2]/2) - 0.5) * c[2]
	dy := (float64(n[1]/2) - 0.5) * c[1]
	dz := (float64(n[0]/2) - 0.5) * c[0]

	for i := 0; i < n[0]; i++ {
		z := float64(i)*c[0] - dz
		for j := 0; j < n[1]; j++ {
			y := float64(j)*c[1] - dy
			for k := 0; k < n[2]; k++ {
				x := float64(k)*c[2] - dx
				if region(x, y, z) { // inside
					m := conf(x, y, z)
					h[0][i][j][k] = float32(m[2])
					h[1][i][j][k] = float32(m[1])
					h[2][i][j][k] = float32(m[0])
				}
			}
		}
	}
	m.Set(host)
}

// set m to config in region
func (m *magnetization) SetRegion(region int, conf Config) {
	host := m.buffer.HostCopy()
	h := host.Vectors()
	n := m.Mesh().Size()
	c := m.Mesh().CellSize()
	dx := (float64(n[2]/2) - 0.5) * c[2]
	dy := (float64(n[1]/2) - 0.5) * c[1]
	dz := (float64(n[0]/2) - 0.5) * c[0]
	r := byte(region)

	for i := 0; i < n[0]; i++ {
		z := float64(i)*c[0] - dz
		for j := 0; j < n[1]; j++ {
			y := float64(j)*c[1] - dy
			for k := 0; k < n[2]; k++ {
				x := float64(k)*c[2] - dx

				if regions.arr[i][j][k] == r {
					m := conf(x, y, z)
					h[0][i][j][k] = float32(m[2])
					h[1][i][j][k] = float32(m[1])
					h[2][i][j][k] = float32(m[0])
				}
			}
		}
	}
	m.Set(host)
}

func (m *magnetization) SetValue(v interface{})  { m.SetInShape(nil, v.(Config)) }
func (m *magnetization) InputType() reflect.Type { return reflect.TypeOf(Config(nil)) }
func (m *magnetization) Type() reflect.Type      { return reflect.TypeOf(new(magnetization)) }
func (m *magnetization) Eval() interface{}       { return m }
