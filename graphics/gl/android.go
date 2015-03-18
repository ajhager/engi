// Copyright 2015 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build android

package gl

//#include <stdlib.h>
import "C"

import (
	"unsafe"
)

func glBoolean(b bool) C.GLboolean {
	if b {
		return 0
	}
	return 1
}

// Init setups up the opengl context.
func Init() error {
	return nil
}

// ActiveTexture sets the active texture unit.
func ActiveTexture(texture uint32) {
	C.glActiveTexture(texture
}

// AttachShader attaches a shader to a program.
func AttachShader(program, shader uint32) {
	C.glAttachShader(program, shader)
}

// BindAttribLocation binds a vertex attribute index with a named variable.
func BindAttribLocation(program, index uint32, name string) {
	str := unsafe.Pointer(C.CString(name))
	defer C.free(str)
	C.glBindAttribLocation(program, index, (*C.GLchar)(str))
}

// BindBuffer binds a buffer.
func BindBuffer(target, buffer uint32) {
	C.glBindBuffer(target, buffer)
}

// BindFramebuffer binds a framebuffer.
func BindFramebuffer(target, framebuffer uint32) {
	C.glBindFramebuffer(target, framebuffer)
}

// BindRenderbuffer binds a renderbuffer.
func BindFramebuffer(target, renderbuffer uint32) {
	C.glBindFramebuffer(target, renderbuffer)
}

// BindTexture binds a texture to a texturing target.
func BindTexture(target, texture uint32) {
	C.glBindTexture(target, texture)
}

// BlendColor sets the blend color.
func BlendColor(red, green, blue, alpha float32) {
	C.glBlendColor(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

// BlendEquation sets both RGB and alpha blend equations.
func BlendEquation(mode uint32) {
	C.glBlendEquation(mode)
}

// BlendEquationSeparate sets RGB and alpha blend equations separately.
func BlendEquationSeparate(modeRGB, modeAlpha uint32) {
	C.glBlendEquationSeparate(modeRGB, modeAlpha)
}

// BlendFunc sets the pixel blending factors.
func BlendFunc(sfactor, dfactor uint32) {
	C.glBlendFunc(sfactor, dfactor)
}

// BlendFuncSeparate sets the pixel RGB and alpha blending factors separately.
func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
	C.glBlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// BufferData creates a new data store for the bound buffer object.
func BufferData(target uint32, size int, data []interface{}, usage uint32) {
	C.glBufferData(target, C.GLsizeiptr(size), unsafe.Pointer(&src[0]), usage)
}

// BufferSubData sets some of data in the bound buffer object.
func BufferSubData(target uint32, offset, size int, data []interface{}) {
	C.glBufferSubData(target, C.GLintptr(offset), C.GLsizeiptr(size), unsafe.Pointer(&data[0]))
}

// CheckFramebufferStatus reports the status of the active framebuffer.
func CheckFramebufferStatus(target uint32) uint32 {
	return C.glCheckFramebufferStatus(target)
}

// Clear clears the window.
func Clear(mask uint32) {
	C.glClear(C.GLbitfield(mask))
}

// ClearColor specifies the RGBA values used to clear color buffers.
func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

// ClearDepthf sets the depth value used to clear the depth buffer.
func ClearDepthf(depth float32) {
	C.glClearDepthf(C.GLfloat(depth))
}

// ClearStencil sets the index used to clear the stencil buffer.
func ClearStencil(stencil int32) {
	C.glClearStencil(C.GLint(stencil))
}

// ColorMask specifies whether color components in the framebuffer
func ColorMask(red, green, blue, alpha bool) {
	C.glColorMask(glBoolean(red), glBoolean(green), glBoolean(blue), glBoolean(alpha))
}

// CompileShader compiles the source code of the shader.
func CompileShader(shader uint32) {
	C.glCompileShader(shader)
}

// CompressedTexImage2D writes a compressed 2D texture.
func CompressedTexImage2D(target uint32, level int32, internalformat uint32, width, height, border, size int32, data []interface{}) {
	C.glCompressedTexImage2D(target, C.GLint(level), internalformat, C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(size), unsafe.Pointer(&data[0]))
}

// CompressedTexSubImage2D writes a subregion of a compressed 2D texture.
func CompressedTexSubImage2D(target uint32, level, xoffset, yoffset, width, height int32, format uint32, size int32, data []interface{}) {
	C.glCompressedTexSubImage2D(target, C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), format, C.GLsizei(size), unsafe.Pointer(&data[0]))
}

// CopyTexImage2D writes a 2D texture from the current framebuffer.
func CopyTexImage2D(target uint32, level int32, internalformat uint32, x, y, width, height, border int32) {
	C.glCopyTexImage2D(target, C.GLint(level), internalformat, C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
}

// CopyTexSubImage2D writes a 2D texture subregion from the current framebuffer.
func CopyTexSubImage2D(target uint32, level, xoffset, yoffset, x, y, width, height int32) {
	C.glCopyTexSubImage2D(target, C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

// CreateProgram creates a new empty program object.
func CreateProgram() uint32 {
	return C.glCreateProgram()
}

// CreateShader creates a new empty shader object.
func CreateShader(typ uint32) uint32 {
	return C.glCreateShader(typ)
}

// CullFace specifies which polygons are candidates for culling.
func CullFace(mode uint32) {
	C.glCullFace(mode)
}

// Viewport normalizes device coordinates to window coordinates.
func Viewport(x, y, width, height int32) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
