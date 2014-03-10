package engine

import (
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
	"math"
)

// Bogacki-Shampine solver. 3rd order, 3 evaluations per step, adaptive step.
// 	http://en.wikipedia.org/wiki/Bogacki-Shampine_method
//
// 	k1 = f(tn, yn)
// 	k2 = f(tn + 1/2 h, yn + 1/2 h k1)
// 	k3 = f(tn + 3/4 h, yn + 3/4 h k2)
// 	y{n+1}  = yn + 2/9 h k1 + 1/3 h k2 + 4/9 h k3            // 3rd order
// 	k4 = f(tn + h, y{n+1})
// 	z{n+1} = yn + 7/24 h k1 + 1/4 h k2 + 1/3 h k3 + 1/8 h k4 // 2nd order
type RK23 struct {
	k1 *data.Slice // torque at end of step is kept for beginning of next step
}

func (rk *RK23) Step() {
	m := M.Buffer()
	size := m.Size()

	if FixDt != 0 {
		Dt_si = FixDt
	}

	// upon resize: remove wrongly sized k1
	if rk.k1.Size() != m.Size() {
		rk.Free()
	}

	// first step ever: one-time k1 init and eval
	if rk.k1 == nil {
		rk.k1 = cuda.NewSlice(3, size)
		torqueFn(rk.k1)
	}

	t0 := Time
	// backup magnetization
	m0 := cuda.Buffer(3, size)
	defer cuda.Recycle(m0)
	data.Copy(m0, m)

	k1 := rk.k1 // from previous step
	k2, k3, k4 := cuda.Buffer(3, size), cuda.Buffer(3, size), cuda.Buffer(3, size)
	defer cuda.Recycle(k2)
	defer cuda.Recycle(k3)
	defer cuda.Recycle(k4)

	h := float32(Dt_si * *dt_mul) // internal time step = Dt * gammaLL

	// there is no explicit stage 1: k1 from previous step

	// stage 2
	Time = t0 + (1./2.)*Dt_si
	cuda.Madd2(m, m0, k1, 1, (1./2.)*h) // m = m0*1 + k1*h/2
	//postStep()
	torqueFn(k2)

	// stage 3
	Time = t0 + (3./4.)*Dt_si
	cuda.Madd2(m, m0, k2, 1, (3./4.)*h) // m = m0*1 + k2*3/4
	//postStep()
	torqueFn(k3)

	// 3rd order solution
	τ3 := cuda.Buffer(3, size)
	defer cuda.Recycle(τ3)
	cuda.Madd3(τ3, k1, k2, k3, (2. / 9.), (1. / 3.), (4. / 9.))
	cuda.Madd2(m, m0, τ3, 1, h)
	solverPostStep()

	// error estimate
	τ2 := cuda.Buffer(3, size)
	defer cuda.Recycle(τ2)
	Time = t0 + Dt_si
	torqueFn(k4)
	madd4(τ2, k1, k2, k3, k4, (7. / 24.), (1. / 4.), (1. / 3.), (1. / 8.))

	// determine error
	err := cuda.MaxVecDiff(τ2, τ3) * float64(h)

	// adjust next time step
	if err < MaxErr || Dt_si <= MinDt || FixDt != 0 { // mindt check to avoid infinite loop
		// step OK
		setLastErr(err)
		NSteps++
		Time = t0 + Dt_si
		adaptDt(math.Pow(MaxErr/err, 1./3.))
		data.Copy(rk.k1, k4) // FSAL
	} else {
		// undo bad step
		util.Assert(FixDt == 0)
		Time = t0
		data.Copy(m, m0)
		NUndone++
		adaptDt(math.Pow(MaxErr/err, 1./4.))
	}
}

func (rk *RK23) Free() {
	rk.k1.Free()
	rk.k1 = nil
}

// TODO: into cuda
func madd4(dst, src1, src2, src3, src4 *data.Slice, w1, w2, w3, w4 float32) {
	cuda.Madd3(dst, src1, src2, src3, w1, w2, w3)
	cuda.Madd2(dst, dst, src4, 1, w4)
}
