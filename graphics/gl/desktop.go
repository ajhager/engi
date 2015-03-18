// Copyright 2015 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js,!android

package gl

import (
	"github.com/go-gl/gl/v2.1/gl"
)

// Init setups up the opengl context.
func Init() error {
	return gl.Init()
}

// ActiveTexture sets the active texture unit.
func ActiveTexture(texture uint32) {
	gl.ActiveTexture(texture)
}

// AttachShader attaches a shader to a program.
func AttachShader(program, shader uint32) {
	gl.AttachShader(program, shader)
}

// BindAttribLocation binds a vertex attribute index with a named variable.
func BindAttribLocation(program, index uint32, name string) {
	gl.BindAttribLocation(program, index, gl.Str(name+"\x00"))
}

// BindBuffer binds a buffer.
func BindBuffer(target, buffer uint32) {
	gl.BindBuffer(target, buffer)
}

// BindFramebuffer binds a framebuffer.
func BindFramebuffer(target, framebuffer uint32) {
	gl.BindFramebuffer(target, framebuffer)
}

// BindRenderbuffer binds a renderbuffer.
func BindRenderbuffer(target, renderbuffer uint32) {
	gl.BindRenderbuffer(target, renderbuffer)
}

// BindTexture binds a texture to a texturing target.
func BindTexture(target, texture uint32) {
	gl.BindTexture(target, texture)
}

// BlendColor sets the blend color.
func BlendColor(red, green, blue, alpha float32) {
	gl.BlendColor(red, green, blue, alpha)
}

// BlendEquation sets both RGB and alpha blend equations.
func BlendEquation(mode uint32) {
	gl.BlendEquation(mode)
}

// BlendEquationSeparate sets RGB and alpha blend equations separately.
func BlendEquationSeparate(modeRGB, modeAlpha uint32) {
	gl.BlendEquationSeparate(modeRGB, modeAlpha)
}

// BlendFunc sets the pixel blending factors.
func BlendFunc(sfactor, dfactor uint32) {
	gl.BlendFunc(sfactor, dfactor)
}

// BlendFuncSeparate sets the pixel RGB and alpha blending factors separately.
func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
	gl.BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// BufferData creates a new data store for the bound buffer object.
func BufferData(target uint32, size int, data []interface{}, usage uint32) {
	gl.BufferData(target, size, gl.Ptr(data), usage)
}

// BufferSubData sets some of data in the bound buffer object.
func BufferSubData(target uint32, offset, size int, data []interface{}) {
	gl.BufferSubData(target, offset, size, gl.Ptr(data))
}

// CheckFramebufferStatus reports the status of the active framebuffer.
func CheckFramebufferStatus(target uint32) uint32 {
	return gl.CheckFramebufferStatus(target)
}

// Clear clears the window.
func Clear(mask uint32) {
	gl.Clear(mask)
}

// ClearColor specifies the RGBA values used to clear color buffers.
func ClearColor(red, green, blue, alpha float32) {
	gl.ClearColor(red, green, blue, alpha)
}

// ClearDepthf sets the depth value used to clear the depth buffer.
func ClearDepthf(depth float32) {
	gl.ClearDepthf(depth)
}

// ClearStencil sets the index used to clear the stencil buffer.
func ClearStencil(stencil int32) {
	gl.ClearStencil(stencil)
}

// ColorMask specifies whether color components in the framebuffer
func ColorMask(red, green, blue, alpha bool) {
	gl.ColorMask(red, green, blue, alpha)
}

// CompileShader compiles the source code of the shader.
func CompileShader(shader uint32) {
	gl.CompileShader(shader)
}

// CompressedTexImage2D writes a compressed 2D texture.
func CompressedTexImage2D(target uint32, level int32, internalformat uint32, width, height, border, size int32, data []interface{}) {
	gl.CompressedTexImage2D(target, level, internalformat, width, height, border, size, gl.Ptr(data))
}

// CompressedTexSubImage2D writes a subregion of a compressed 2D texture.
func CompressedTexSubImage2D(target uint32, level, xoffset, yoffset, width, height int32, format uint32, size int32, data []interface{}) {
	gl.CompressedTexSubImage2D(target, level, xoffset, yoffset, width, height, format, size, gl.Ptr(data))
}

// CopyTexImage2D writes a 2D texture from the current framebuffer.
func CopyTexImage2D(target uint32, level int32, internalformat uint32, x, y, width, height, border int32) {
	gl.CopyTexImage2D(target, level, internalformat, x, y, width, height, border)
}

// CopyTexSubImage2D writes a 2D texture subregion from the current framebuffer.
func CopyTexSubImage2D(target uint32, level, xoffset, yoffset, x, y, width, height int32) {
	gl.CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height)
}

// CreateProgram creates a new empty program object.
func CreateProgram() uint32 {
	return gl.CreateProgram()
}

// CreateShader creates a new empty shader object.
func CreateShader(typ uint32) uint32 {
	return gl.CreateShader(typ)
}

// CullFace specifies which polygons are candidates for culling.
func CullFace(mode uint32) {
	gl.CullFace(mode)
}

// Viewport normalizes device coordinates to window coordinates.
func Viewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}
