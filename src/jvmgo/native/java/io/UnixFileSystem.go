package io

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"os"
	"path/filepath"
)

func init() {
	_ufs(ufs_initIDs, "initIDs", "()V")
	_ufs(canonicalize0, "canonicalize0", "(Ljava/lang/String;)Ljava/lang/String;")
	_ufs(getBooleanAttributes0, "getBooleanAttributes0", "(Ljava/io/File;)I")
	_ufs(getLastModifiedTime, "getLastModifiedTime", "(Ljava/io/File;)J")
}

func _ufs(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/io/UnixFileSystem", name, desc, method)
}

// private static native void initIDs();
// ()V
func ufs_initIDs(frame *rtda.Frame) {
	// todo
}

// private native String canonicalize0(String path) throws IOException;
// (Ljava/lang/String;)Ljava/lang/String;
func canonicalize0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	pathStr := vars.GetRef(1)

	// todo
	path := rtda.GoString(pathStr)
	path2 := filepath.Clean(path)
	if path2 != path {
		pathStr = rtda.NewJString(path2)
	}

	stack := frame.OperandStack()
	stack.PushRef(pathStr)
}

// public native int getBooleanAttributes0(File f);
// (Ljava/io/File;)I
func getBooleanAttributes0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fileObj := vars.GetRef(1)
	path := _getPath(fileObj)

	// todo
	attributes0 := 0
	if _exists(path) {
		attributes0 |= 0x01
	}
	if _isDir(path) {
		attributes0 |= 0x04
	}

	stack := frame.OperandStack()
	stack.PushInt(int32(attributes0))
}

func _getPath(fileObj *rtc.Obj) string {
	pathStr := fileObj.GetFieldValue("path", "Ljava/lang/String;").(*rtc.Obj)
	return rtda.GoString(pathStr)
}

// http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang
// exists returns whether the given file or directory exists or not
func _exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// http://stackoverflow.com/questions/8824571/golang-determining-whether-file-points-to-file-or-directory
func _isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.IsDir()
	}
	return false
}

// public native long getLastModifiedTime(File f);
//
func getLastModifiedTime(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fileObj := vars.GetRef(1)

	path := _getPath(fileObj)
	t := _getLastModifiedTime(path)

	stack := frame.OperandStack()
	stack.PushLong(t)
}

func _getLastModifiedTime(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.ModTime().UnixNano() / 1000
	}
	return 0
}
