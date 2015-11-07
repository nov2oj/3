package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for minimize kernel
var minimize_code cu.Function

// Stores the arguments for minimize kernel invocation
type minimize_args_t struct {
	arg_mx  unsafe.Pointer
	arg_my  unsafe.Pointer
	arg_mz  unsafe.Pointer
	arg_m0x unsafe.Pointer
	arg_m0y unsafe.Pointer
	arg_m0z unsafe.Pointer
	arg_tx  unsafe.Pointer
	arg_ty  unsafe.Pointer
	arg_tz  unsafe.Pointer
	arg_dt  float32
	arg_N   int
	argptr  [11]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for minimize kernel invocation
var minimize_args minimize_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	minimize_args.argptr[0] = unsafe.Pointer(&minimize_args.arg_mx)
	minimize_args.argptr[1] = unsafe.Pointer(&minimize_args.arg_my)
	minimize_args.argptr[2] = unsafe.Pointer(&minimize_args.arg_mz)
	minimize_args.argptr[3] = unsafe.Pointer(&minimize_args.arg_m0x)
	minimize_args.argptr[4] = unsafe.Pointer(&minimize_args.arg_m0y)
	minimize_args.argptr[5] = unsafe.Pointer(&minimize_args.arg_m0z)
	minimize_args.argptr[6] = unsafe.Pointer(&minimize_args.arg_tx)
	minimize_args.argptr[7] = unsafe.Pointer(&minimize_args.arg_ty)
	minimize_args.argptr[8] = unsafe.Pointer(&minimize_args.arg_tz)
	minimize_args.argptr[9] = unsafe.Pointer(&minimize_args.arg_dt)
	minimize_args.argptr[10] = unsafe.Pointer(&minimize_args.arg_N)
}

// Wrapper for minimize CUDA kernel, asynchronous.
func k_minimize_async(mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, m0x unsafe.Pointer, m0y unsafe.Pointer, m0z unsafe.Pointer, tx unsafe.Pointer, ty unsafe.Pointer, tz unsafe.Pointer, dt float32, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("minimize")
	}

	minimize_args.Lock()
	defer minimize_args.Unlock()

	if minimize_code == 0 {
		minimize_code = fatbinLoad(minimize_map, "minimize")
	}

	minimize_args.arg_mx = mx
	minimize_args.arg_my = my
	minimize_args.arg_mz = mz
	minimize_args.arg_m0x = m0x
	minimize_args.arg_m0y = m0y
	minimize_args.arg_m0z = m0z
	minimize_args.arg_tx = tx
	minimize_args.arg_ty = ty
	minimize_args.arg_tz = tz
	minimize_args.arg_dt = dt
	minimize_args.arg_N = N

	args := minimize_args.argptr[:]
	cu.LaunchKernel(minimize_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("minimize")
	}
}

// maps compute capability on PTX code for minimize kernel.
var minimize_map = map[int]string{0: "",
	20: minimize_ptx_20,
	30: minimize_ptx_30,
	35: minimize_ptx_35,
	50: minimize_ptx_50}

// minimize PTX code for various compute capabilities.
const (
	minimize_ptx_20 = `
.version 4.1
.target sm_20
.address_size 64


.visible .entry minimize(
	.param .u64 minimize_param_0,
	.param .u64 minimize_param_1,
	.param .u64 minimize_param_2,
	.param .u64 minimize_param_3,
	.param .u64 minimize_param_4,
	.param .u64 minimize_param_5,
	.param .u64 minimize_param_6,
	.param .u64 minimize_param_7,
	.param .u64 minimize_param_8,
	.param .f32 minimize_param_9,
	.param .u32 minimize_param_10
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<26>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd1, [minimize_param_0];
	ld.param.u64 	%rd2, [minimize_param_1];
	ld.param.u64 	%rd3, [minimize_param_2];
	ld.param.u64 	%rd4, [minimize_param_3];
	ld.param.u64 	%rd5, [minimize_param_4];
	ld.param.u64 	%rd6, [minimize_param_5];
	ld.param.u64 	%rd7, [minimize_param_6];
	ld.param.u64 	%rd8, [minimize_param_7];
	ld.param.u64 	%rd9, [minimize_param_8];
	ld.param.f32 	%f1, [minimize_param_9];
	ld.param.u32 	%r2, [minimize_param_10];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd3;
	cvta.to.global.u64 	%rd11, %rd2;
	cvta.to.global.u64 	%rd12, %rd1;
	cvta.to.global.u64 	%rd13, %rd9;
	cvta.to.global.u64 	%rd14, %rd8;
	cvta.to.global.u64 	%rd15, %rd7;
	cvta.to.global.u64 	%rd16, %rd6;
	cvta.to.global.u64 	%rd17, %rd5;
	cvta.to.global.u64 	%rd18, %rd4;
	mul.wide.s32 	%rd19, %r1, 4;
	add.s64 	%rd20, %rd18, %rd19;
	add.s64 	%rd21, %rd17, %rd19;
	add.s64 	%rd22, %rd16, %rd19;
	add.s64 	%rd23, %rd15, %rd19;
	add.s64 	%rd24, %rd14, %rd19;
	add.s64 	%rd25, %rd13, %rd19;
	ld.global.f32 	%f2, [%rd23];
	ld.global.f32 	%f3, [%rd24];
	mul.f32 	%f4, %f3, %f3;
	fma.rn.f32 	%f5, %f2, %f2, %f4;
	ld.global.f32 	%f6, [%rd25];
	fma.rn.f32 	%f7, %f6, %f6, %f5;
	mul.f32 	%f8, %f1, %f1;
	mul.f32 	%f9, %f8, %f7;
	mov.f32 	%f10, 0f40800000;
	sub.f32 	%f11, %f10, %f9;
	ld.global.f32 	%f12, [%rd20];
	ld.global.f32 	%f13, [%rd21];
	ld.global.f32 	%f14, [%rd22];
	mul.f32 	%f15, %f1, 0f40800000;
	mul.f32 	%f16, %f15, %f2;
	mul.f32 	%f17, %f15, %f3;
	mul.f32 	%f18, %f15, %f6;
	fma.rn.f32 	%f19, %f11, %f12, %f16;
	fma.rn.f32 	%f20, %f11, %f13, %f17;
	fma.rn.f32 	%f21, %f11, %f14, %f18;
	add.f32 	%f22, %f9, 0f40800000;
	div.rn.f32 	%f23, %f19, %f22;
	add.s64 	%rd26, %rd12, %rd19;
	st.global.f32 	[%rd26], %f23;
	div.rn.f32 	%f24, %f20, %f22;
	add.s64 	%rd27, %rd11, %rd19;
	st.global.f32 	[%rd27], %f24;
	div.rn.f32 	%f25, %f21, %f22;
	add.s64 	%rd28, %rd10, %rd19;
	st.global.f32 	[%rd28], %f25;

BB0_2:
	ret;
}


`
	minimize_ptx_30 = `
.version 4.1
.target sm_30
.address_size 64


.visible .entry minimize(
	.param .u64 minimize_param_0,
	.param .u64 minimize_param_1,
	.param .u64 minimize_param_2,
	.param .u64 minimize_param_3,
	.param .u64 minimize_param_4,
	.param .u64 minimize_param_5,
	.param .u64 minimize_param_6,
	.param .u64 minimize_param_7,
	.param .u64 minimize_param_8,
	.param .f32 minimize_param_9,
	.param .u32 minimize_param_10
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<26>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd1, [minimize_param_0];
	ld.param.u64 	%rd2, [minimize_param_1];
	ld.param.u64 	%rd3, [minimize_param_2];
	ld.param.u64 	%rd4, [minimize_param_3];
	ld.param.u64 	%rd5, [minimize_param_4];
	ld.param.u64 	%rd6, [minimize_param_5];
	ld.param.u64 	%rd7, [minimize_param_6];
	ld.param.u64 	%rd8, [minimize_param_7];
	ld.param.u64 	%rd9, [minimize_param_8];
	ld.param.f32 	%f1, [minimize_param_9];
	ld.param.u32 	%r2, [minimize_param_10];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd10, %rd3;
	cvta.to.global.u64 	%rd11, %rd2;
	cvta.to.global.u64 	%rd12, %rd1;
	cvta.to.global.u64 	%rd13, %rd9;
	cvta.to.global.u64 	%rd14, %rd8;
	cvta.to.global.u64 	%rd15, %rd7;
	cvta.to.global.u64 	%rd16, %rd6;
	cvta.to.global.u64 	%rd17, %rd5;
	cvta.to.global.u64 	%rd18, %rd4;
	mul.wide.s32 	%rd19, %r1, 4;
	add.s64 	%rd20, %rd18, %rd19;
	add.s64 	%rd21, %rd17, %rd19;
	add.s64 	%rd22, %rd16, %rd19;
	add.s64 	%rd23, %rd15, %rd19;
	add.s64 	%rd24, %rd14, %rd19;
	add.s64 	%rd25, %rd13, %rd19;
	ld.global.f32 	%f2, [%rd23];
	ld.global.f32 	%f3, [%rd24];
	mul.f32 	%f4, %f3, %f3;
	fma.rn.f32 	%f5, %f2, %f2, %f4;
	ld.global.f32 	%f6, [%rd25];
	fma.rn.f32 	%f7, %f6, %f6, %f5;
	mul.f32 	%f8, %f1, %f1;
	mul.f32 	%f9, %f8, %f7;
	mov.f32 	%f10, 0f40800000;
	sub.f32 	%f11, %f10, %f9;
	ld.global.f32 	%f12, [%rd20];
	ld.global.f32 	%f13, [%rd21];
	ld.global.f32 	%f14, [%rd22];
	mul.f32 	%f15, %f1, 0f40800000;
	mul.f32 	%f16, %f15, %f2;
	mul.f32 	%f17, %f15, %f3;
	mul.f32 	%f18, %f15, %f6;
	fma.rn.f32 	%f19, %f11, %f12, %f16;
	fma.rn.f32 	%f20, %f11, %f13, %f17;
	fma.rn.f32 	%f21, %f11, %f14, %f18;
	add.f32 	%f22, %f9, 0f40800000;
	div.rn.f32 	%f23, %f19, %f22;
	add.s64 	%rd26, %rd12, %rd19;
	st.global.f32 	[%rd26], %f23;
	div.rn.f32 	%f24, %f20, %f22;
	add.s64 	%rd27, %rd11, %rd19;
	st.global.f32 	[%rd27], %f24;
	div.rn.f32 	%f25, %f21, %f22;
	add.s64 	%rd28, %rd10, %rd19;
	st.global.f32 	[%rd28], %f25;

BB0_2:
	ret;
}


`
	minimize_ptx_35 = `
.version 4.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry minimize(
	.param .u64 minimize_param_0,
	.param .u64 minimize_param_1,
	.param .u64 minimize_param_2,
	.param .u64 minimize_param_3,
	.param .u64 minimize_param_4,
	.param .u64 minimize_param_5,
	.param .u64 minimize_param_6,
	.param .u64 minimize_param_7,
	.param .u64 minimize_param_8,
	.param .f32 minimize_param_9,
	.param .u32 minimize_param_10
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<26>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd1, [minimize_param_0];
	ld.param.u64 	%rd2, [minimize_param_1];
	ld.param.u64 	%rd3, [minimize_param_2];
	ld.param.u64 	%rd4, [minimize_param_3];
	ld.param.u64 	%rd5, [minimize_param_4];
	ld.param.u64 	%rd6, [minimize_param_5];
	ld.param.u64 	%rd7, [minimize_param_6];
	ld.param.u64 	%rd8, [minimize_param_7];
	ld.param.u64 	%rd9, [minimize_param_8];
	ld.param.f32 	%f1, [minimize_param_9];
	ld.param.u32 	%r2, [minimize_param_10];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB5_2;

	cvta.to.global.u64 	%rd10, %rd3;
	cvta.to.global.u64 	%rd11, %rd2;
	cvta.to.global.u64 	%rd12, %rd1;
	cvta.to.global.u64 	%rd13, %rd9;
	cvta.to.global.u64 	%rd14, %rd8;
	cvta.to.global.u64 	%rd15, %rd7;
	cvta.to.global.u64 	%rd16, %rd6;
	cvta.to.global.u64 	%rd17, %rd5;
	cvta.to.global.u64 	%rd18, %rd4;
	mul.wide.s32 	%rd19, %r1, 4;
	add.s64 	%rd20, %rd18, %rd19;
	add.s64 	%rd21, %rd17, %rd19;
	add.s64 	%rd22, %rd16, %rd19;
	add.s64 	%rd23, %rd15, %rd19;
	add.s64 	%rd24, %rd14, %rd19;
	add.s64 	%rd25, %rd13, %rd19;
	ld.global.nc.f32 	%f2, [%rd23];
	ld.global.nc.f32 	%f3, [%rd24];
	mul.f32 	%f4, %f3, %f3;
	fma.rn.f32 	%f5, %f2, %f2, %f4;
	ld.global.nc.f32 	%f6, [%rd25];
	fma.rn.f32 	%f7, %f6, %f6, %f5;
	mul.f32 	%f8, %f1, %f1;
	mul.f32 	%f9, %f8, %f7;
	mov.f32 	%f10, 0f40800000;
	sub.f32 	%f11, %f10, %f9;
	ld.global.nc.f32 	%f12, [%rd20];
	ld.global.nc.f32 	%f13, [%rd21];
	ld.global.nc.f32 	%f14, [%rd22];
	mul.f32 	%f15, %f1, 0f40800000;
	mul.f32 	%f16, %f15, %f2;
	mul.f32 	%f17, %f15, %f3;
	mul.f32 	%f18, %f15, %f6;
	fma.rn.f32 	%f19, %f11, %f12, %f16;
	fma.rn.f32 	%f20, %f11, %f13, %f17;
	fma.rn.f32 	%f21, %f11, %f14, %f18;
	add.f32 	%f22, %f9, 0f40800000;
	div.rn.f32 	%f23, %f19, %f22;
	add.s64 	%rd26, %rd12, %rd19;
	st.global.f32 	[%rd26], %f23;
	div.rn.f32 	%f24, %f20, %f22;
	add.s64 	%rd27, %rd11, %rd19;
	st.global.f32 	[%rd27], %f24;
	div.rn.f32 	%f25, %f21, %f22;
	add.s64 	%rd28, %rd10, %rd19;
	st.global.f32 	[%rd28], %f25;

BB5_2:
	ret;
}


`
	minimize_ptx_50 = `
.version 4.1
.target sm_50
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

.visible .entry minimize(
	.param .u64 minimize_param_0,
	.param .u64 minimize_param_1,
	.param .u64 minimize_param_2,
	.param .u64 minimize_param_3,
	.param .u64 minimize_param_4,
	.param .u64 minimize_param_5,
	.param .u64 minimize_param_6,
	.param .u64 minimize_param_7,
	.param .u64 minimize_param_8,
	.param .f32 minimize_param_9,
	.param .u32 minimize_param_10
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<26>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd1, [minimize_param_0];
	ld.param.u64 	%rd2, [minimize_param_1];
	ld.param.u64 	%rd3, [minimize_param_2];
	ld.param.u64 	%rd4, [minimize_param_3];
	ld.param.u64 	%rd5, [minimize_param_4];
	ld.param.u64 	%rd6, [minimize_param_5];
	ld.param.u64 	%rd7, [minimize_param_6];
	ld.param.u64 	%rd8, [minimize_param_7];
	ld.param.u64 	%rd9, [minimize_param_8];
	ld.param.f32 	%f1, [minimize_param_9];
	ld.param.u32 	%r2, [minimize_param_10];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB5_2;

	cvta.to.global.u64 	%rd10, %rd3;
	cvta.to.global.u64 	%rd11, %rd2;
	cvta.to.global.u64 	%rd12, %rd1;
	cvta.to.global.u64 	%rd13, %rd9;
	cvta.to.global.u64 	%rd14, %rd8;
	cvta.to.global.u64 	%rd15, %rd7;
	cvta.to.global.u64 	%rd16, %rd6;
	cvta.to.global.u64 	%rd17, %rd5;
	cvta.to.global.u64 	%rd18, %rd4;
	mul.wide.s32 	%rd19, %r1, 4;
	add.s64 	%rd20, %rd18, %rd19;
	add.s64 	%rd21, %rd17, %rd19;
	add.s64 	%rd22, %rd16, %rd19;
	add.s64 	%rd23, %rd15, %rd19;
	add.s64 	%rd24, %rd14, %rd19;
	add.s64 	%rd25, %rd13, %rd19;
	ld.global.nc.f32 	%f2, [%rd23];
	ld.global.nc.f32 	%f3, [%rd24];
	mul.f32 	%f4, %f3, %f3;
	fma.rn.f32 	%f5, %f2, %f2, %f4;
	ld.global.nc.f32 	%f6, [%rd25];
	fma.rn.f32 	%f7, %f6, %f6, %f5;
	mul.f32 	%f8, %f1, %f1;
	mul.f32 	%f9, %f8, %f7;
	mov.f32 	%f10, 0f40800000;
	sub.f32 	%f11, %f10, %f9;
	ld.global.nc.f32 	%f12, [%rd20];
	ld.global.nc.f32 	%f13, [%rd21];
	ld.global.nc.f32 	%f14, [%rd22];
	mul.f32 	%f15, %f1, 0f40800000;
	mul.f32 	%f16, %f15, %f2;
	mul.f32 	%f17, %f15, %f3;
	mul.f32 	%f18, %f15, %f6;
	fma.rn.f32 	%f19, %f11, %f12, %f16;
	fma.rn.f32 	%f20, %f11, %f13, %f17;
	fma.rn.f32 	%f21, %f11, %f14, %f18;
	add.f32 	%f22, %f9, 0f40800000;
	div.rn.f32 	%f23, %f19, %f22;
	add.s64 	%rd26, %rd12, %rd19;
	st.global.f32 	[%rd26], %f23;
	div.rn.f32 	%f24, %f20, %f22;
	add.s64 	%rd27, %rd11, %rd19;
	st.global.f32 	[%rd27], %f24;
	div.rn.f32 	%f25, %f21, %f22;
	add.s64 	%rd28, %rd10, %rd19;
	st.global.f32 	[%rd28], %f25;

BB5_2:
	ret;
}


`
)
