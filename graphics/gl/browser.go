// Copyright 2015 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js

package gl

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

var (
	context           *js.Object
	textures          = make(map[uint32]*js.Object)
	textureId         = uint32(0)
	buffers           = make(map[uint32]*js.Object)
	bufferId          = uint32(0)
	framebuffers      = make(map[uint32]*js.Object)
	framebufferId     = uint32(0)
	renderbuffers     = make(map[uint32]*js.Object)
	renderbufferId    = uint32(0)
	programs          = make(map[uint32]*js.Object)
	programId         = uint32(0)
	shaders           = make(map[uint32]*js.Object)
	shaderId          = uint32(0)
	uniformLocations  = make(map[uint32]map[int32]*js.Object)
	uniformLocationId = int32(0)
	currentProgram    = uint32(0)
)

// Init setups up the opengl context.
func Init(webglContext *js.Object) error {
	context = webglContext

	webglProto := js.Global.Get("WebGLRenderingContext").Get("prototype")
	same := webglProto == js.Global.Get("Object").Call("getPrototypeOf", context)
	if !same {
		return errors.New("Supplied object is not a vaild rendering context")
	}

	if context.Call("getExtension", "WEBGL_compressed_texture_s3tc") == nil {
		return errors.New("CompressedTextImage2D and CompressedTexSubImage2D unavailable")
	}

	return nil
}

// ActiveTexture sets the active texture unit.
func ActiveTexture(texture uint32) {
	context.Call("activeTexture", texture)
}

// AttachShader attaches a shader to a program.
func AttachShader(program, shader uint32) {
	context.Call("attachShader", programs[program], shaders[shader])
}

// BindAttribLocation binds a vertex attribute index with a named variable.
func BindAttribLocation(program, index uint32, name string) {
	context.Call("bindAttribLocation", programs[program], index, name)
}

// BindBuffer binds a buffer.
func BindBuffer(target, buffer uint32) {
	context.Call("bindBuffer", target, buffers[buffer])
}

// BindFramebuffer binds a framebuffer.
func BindFramebuffer(target, framebuffer uint32) {
	context.Call("bindFramebuffer", target, framebuffers[framebuffer])
}

// BindRenderbuffer binds a renderbuffer.
func BindRenderbuffer(target, renderbuffer uint32) {
	context.Call("bindRenderbuffer", target, renderbuffers[renderbuffer])
}

// BindTexture binds a texture to a texturing target.
func BindTexture(target, texture uint32) {
	context.Call("bindTexture", target, textures[texture])
}

// BlendColor sets the blend color.
func BlendColor(red, green, blue, alpha float32) {
	context.Call("blendColor", red, green, blue, alpha)
}

// BlendEquation sets both RGB and alpha blend equations.
func BlendEquation(mode uint32) {
	context.Call("blendEquation", mode)
}

// BlendEquationSeparate sets RGB and alpha blend equations separately.
func BlendEquationSeparate(modeRGB, modeAlpha uint32) {
	context.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

// BlendFunc sets the pixel blending factors.
func BlendFunc(sfactor, dfactor uint32) {
	context.Call("blendFunc", sfactor, dfactor)
}

// BlendFuncSeparate sets the pixel RGB and alpha blending factors separately.
func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha uint32) {
	context.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// BufferData creates a new data store for the bound buffer object.
func BufferData(target uint32, size int, data []interface{}, usage uint32) {
	context.Call("bufferData", target, data, usage)
}

// BufferSubData sets some of data in the bound buffer object.
func BufferSubData(target uint32, offset, size int, data []interface{}) {
	context.Call("bufferSubData", target, offset, data)
}

// CheckFramebufferStatus reports the status of the active framebuffer.
func CheckFramebufferStatus(target uint32) uint32 {
	return uint32(context.Call("checkFramebufferStatus", target).Uint64())
}

// Clear clears the window.
func Clear(mask uint32) {
	context.Call("clear", mask)
}

// ClearColor specifies the RGBA values used to clear color buffers.
func ClearColor(red, green, blue, alpha float32) {
	context.Call("clearColor", red, green, blue, alpha)
}

// ClearDepthf sets the depth value used to clear the depth buffer.
func ClearDepthf(depth float32) {
	context.Call("clearDepth", depth)
}

// ClearStencil sets the index used to clear the stencil buffer.
func ClearStencil(stencil int32) {
	context.Call("clearStencil", stencil)
}

// ColorMask specifies whether color components in the framebuffer
func ColorMask(red, green, blue, alpha bool) {
	context.Call("colorMask", red, green, blue, alpha)
}

// CompileShader compiles the source code of the shader.
func CompileShader(shader uint32) {
	context.Call("compileShader", shaders[shader])
}

// CompressedTexImage2D writes a compressed 2D texture.
func CompressedTexImage2D(target uint32, level int32, internalformat uint32, width, height, border, size int32, data []interface{}) {
	context.Call("compressedTexImage2D", target, level, internalformat, width, height, border, data)
}

// CompressedTexSubImage2D writes a subregion of a compressed 2D texture.
func CompressedTexSubImage2D(target uint32, level, xoffset, yoffset, width, height int32, format uint32, size int32, data []interface{}) {
	context.Call("compressedTexSubImage2D", target, level, xoffset, yoffset, width, height, format, data)
}

// CopyTexImage2D writes a 2D texture from the current framebuffer.
func CopyTexImage2D(target uint32, level int32, internalformat uint32, x, y, width, height, border int32) {
	context.Call("copyTexImage2D", target, level, internalformat, x, y, width, height, border)
}

// CopyTexSubImage2D writes a 2D texture subregion from the current framebuffer.
func CopyTexSubImage2D(target uint32, level, xoffset, yoffset, x, y, width, height int32) {
	context.Call("copyTexSubImage2D", target, level, xoffset, yoffset, x, y, width, height)
}

// CreateProgram creates a new empty program object.
func CreateProgram() uint32 {
	programId += 1
	programs[programId] = context.Call("createProgram")
	return programId
}

// CreateShader creates a new empty shader object.
func CreateShader(typ uint32) uint32 {
	shaderId += 1
	shaders[shaderId] = context.Call("createShader", typ)
	return shaderId
}

// CullFace specifies which polygons are candidates for culling.
func CullFace(mode uint32) {
	context.Call("cullFace", mode)
}

// Viewport normalizes device coordinates to window coordinates.
func Viewport(x, y, width, height int32) {
	context.Call("viewport", x, y, width, height)
}
