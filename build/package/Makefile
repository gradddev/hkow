include $(TOPDIR)/rules.mk

PKG_NAME:=hkow
PKG_VERSION:=1.1.0

PKG_SOURCE_PROTO:=git
PKG_SOURCE_URL:=https://github.com/AlexeySemigradsky/hkow.git
PKG_SOURCE_VERSION:=33385fe79f0e63de78400c55db8140de3644f6c8

PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_PARALLEL:=1
PKG_USE_MIPS16:=0

GO_PKG:=github.com/AlexeySemigradsky/hkow

include $(INCLUDE_DIR)/package.mk
include $(TOPDIR)/feeds/packages/lang/golang/golang-package.mk

define Package/hkow
	SECTION=:net
 	CATEGORY:=Network
	TITLE:=hkow
	URL:=https://github.com/AlexeySemigradsky/hkow
	DEPENDS:=$(GO_ARCH_DEPENDS)
endef

define Package/hkow/description
The HomeKit bridge for the OpenWrt

endef

$(eval $(call GoBinPackage,hkow))
$(eval $(call BuildPackage,hkow))
