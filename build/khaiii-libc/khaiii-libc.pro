TEMPLATE = lib
TARGET = khaiiic
CONFIG += console c++14
CONFIG -= app_bundle
CONFIG -= qt

SOURCES += \
        khaiiic_core.cpp \
        khaiiic.cc

HEADERS += \
    khaiii/KhaiiiApi.hpp \
    khaiii/khaiii_api.h \
    khaiii/khaiii_dev.h \
    khaiiic.h \
    khaiiic_core.h

INCLUDEPATH += \
    khaiii

LIBS += \
    -L"$$PWD/lib" -lkhaiii
