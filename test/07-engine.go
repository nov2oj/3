// +build ignore

package main

import . "code.google.com/p/mx3/engine"

func main() {
	Init()
	defer Close()

	const Nx, Ny = 128, 32
	SetMesh(Nx, Ny, 1, 500e-9/Nx, 125e-9/Ny, 3e-9)

	SetM(1, 1, 0)
	Msat = Const(800e3)
	Aex = Const(13e-12)
	Alpha = Const(1)

	//	B_exch.Autosave(0.1e-9)
	//	B_demag.Autosave(0.1e-9)
	//M.Autosave(0.1e-9)
	B_eff.Autosave(0.2e-9)
	Torque.Autosave(0.1e-9)
	Run(5e-9)
}
