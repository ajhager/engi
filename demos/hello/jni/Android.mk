LOCAL_PATH := $(call my-dir)
include $(CLEAR_VARS)

LOCAL_MODULE    := basic
LOCAL_SRC_FILES := $(TARGET_ARCH_ABI)/libbasic.so

include $(PREBUILT_SHARED_LIBRARY)
