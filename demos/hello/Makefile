desktop:
	go build

browser:
	gopherjs build

android:
	mkdir -p jni/armeabi
	CGO_ENABLED=1 GOOS=android GOARCH=arm GOARM=7 go build -ldflags="-shared" -o jni/armeabi/libbasic.so .
	ndk-build NDK_DEBUG=1
	ant debug
