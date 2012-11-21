package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["copypad"] = COPYPAD }

const COPYPAD = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00002f05_00000000-9_copypad.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/copypad.cu"

.visible .entry copypad(
	.param .u64 copypad_param_0,
	.param .u32 copypad_param_1,
	.param .u32 copypad_param_2,
	.param .u32 copypad_param_3,
	.param .u64 copypad_param_4,
	.param .u32 copypad_param_5,
	.param .u32 copypad_param_6,
	.param .u32 copypad_param_7,
	.param .u32 copypad_param_8,
	.param .u32 copypad_param_9,
	.param .u32 copypad_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<40>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [copypad_param_0];
	ld.param.u32 	%r17, [copypad_param_2];
	ld.param.u32 	%r18, [copypad_param_3];
	ld.param.u64 	%rd4, [copypad_param_4];
	ld.param.u32 	%r19, [copypad_param_5];
	ld.param.u32 	%r20, [copypad_param_6];
	ld.param.u32 	%r21, [copypad_param_7];
	ld.param.u32 	%r22, [copypad_param_8];
	ld.param.u32 	%r23, [copypad_param_9];
	ld.param.u32 	%r24, [copypad_param_10];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 13 1
	mov.u32 	%r1, %ntid.y;
	mov.u32 	%r2, %ctaid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r25, %r1, %r2, %r3;
	.loc 2 14 1
	mov.u32 	%r4, %ntid.x;
	mov.u32 	%r5, %ctaid.x;
	mov.u32 	%r6, %tid.x;
	mad.lo.s32 	%r26, %r4, %r5, %r6;
	.loc 2 16 1
	setp.lt.s32 	%p1, %r25, %r20;
	setp.lt.s32 	%p2, %r26, %r21;
	and.pred  	%p3, %p2, %p1;
	setp.lt.s32 	%p4, %r25, %r17;
	and.pred  	%p5, %p3, %p4;
	setp.lt.s32 	%p6, %r26, %r18;
	and.pred  	%p7, %p5, %p6;
	.loc 2 23 1
	setp.gt.s32 	%p8, %r19, 0;
	.loc 2 16 1
	and.pred  	%p9, %p8, %p7;
	@!%p9 bra 	BB0_3;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 23 1
	add.s32 	%r28, %r6, %r24;
	mad.lo.s32 	%r29, %r4, %r5, %r28;
	add.s32 	%r30, %r3, %r23;
	mad.lo.s32 	%r31, %r1, %r2, %r30;
	mad.lo.s32 	%r32, %r22, %r17, %r31;
	mad.lo.s32 	%r38, %r18, %r32, %r29;
	mul.lo.s32 	%r8, %r18, %r17;
	mad.lo.s32 	%r37, %r21, %r25, %r26;
	mul.lo.s32 	%r10, %r21, %r20;
	mov.u32 	%r39, 0;

BB0_2:
	.loc 2 25 1
	mul.wide.s32 	%rd5, %r37, 4;
	add.s64 	%rd6, %rd2, %rd5;
	mul.wide.s32 	%rd7, %r38, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.f32 	%f1, [%rd6];
	st.global.f32 	[%rd8], %f1;
	.loc 2 23 1
	add.s32 	%r38, %r38, %r8;
	add.s32 	%r37, %r37, %r10;
	.loc 2 23 18
	add.s32 	%r39, %r39, 1;
	.loc 2 23 1
	setp.lt.s32 	%p10, %r39, %r19;
	@%p10 bra 	BB0_2;

BB0_3:
	.loc 2 27 2
	ret;
}


`
