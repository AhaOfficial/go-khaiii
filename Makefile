SH = bash
INSTALL_KHAIII = .github/install_base.sh
INSTALL_KHAIII_LIBC = .github/install_libc_for_go.sh
KHAIII_PATH = .github/build/khaiii
KHAIII_LIB_PATH = .github/build/khaiii-libc
LIB_KHAIII = libkhaiii
LIB_KHAIII_C = libkhaiiic
LIB_KHAIII_C_HEADER = khaiiic.h
KHAIII_DIC_DIR = share/khaiii
USR_LIB_PATH = /usr/local/lib
USR_INCLUDE_PATH = /usr/local/include
USR_SHARE_PATH = /usr/local/share

all:
	@$(SH) $(INSTALL_KHAIII)
	@$(SH) $(INSTALL_KHAIII_LIBC)

install:
	@mkdir -p $(USR_LIB_PATH)
	@mkdir -p $(USR_INCLUDE_PATH)
	@mkdir -p $(USR_SHARE_PATH)
	@cp -pf $(LIB_KHAIII).* $(USR_LIB_PATH)
	@cp -pf $(LIB_KHAIII_C).* $(USR_LIB_PATH)
	@cp -pf $(LIB_KHAIII_C_HEADER) $(USR_INCLUDE_PATH)
	@cp -pfr $(KHAIII_DIC_DIR) $(USR_SHARE_PATH)

clean:
	@rm -f $(LIB_KHAIII).*
	@rm -f $(LIB_KHAIII_C).*
	@rm -f $(LIB_KHAIII_C_HEADER)
	@rm -rf share
	@rm -rf $(KHAIII_PATH)
	@rm -f $(KHAIII_LIB_PATH)/$(LIB_KHAIII_C).*
	@rm -f $(KHAIII_LIB_PATH)/*.o
	@rm -rf $(KHAIII_LIB_PATH)/khaiii
	@rm -rf $(KHAIII_LIB_PATH)/lib
	@rm -f $(KHAIII_LIB_PATH)/Makefile
	@rm -f .github/$(LIB_KHAIII_C_HEADER)
	@rm -f .github/$(LIB_KHAIII_C).*

uninstall:
	@rm -f $(USR_LIB_PATH)/$(LIB_KHAIII).*
	@rm -f $(USR_LIB_PATH)/$(LIB_KHAIII_C).*
	@rm -f $(USR_INCLUDE_PATH)/$(LIB_KHAIII_C_HEADER)
	@rm -rf $(USR_SHARE_PATH)/khaiii

