include $(TOPDIR)/rules.mk

PKG_NAME:=hkow
PKG_VERSION:=1.5.2

PKG_SOURCE_PROTO:=git
PKG_SOURCE_URL:=https://github.com/gradddev/hkow.git
PKG_SOURCE_VERSION:=v$(PKG_VERSION)

PKG_BUILD_DEPENDS:=golang/host
PKG_BUILD_PARALLEL:=1
PKG_USE_MIPS16:=0

GO_PKG:=github.com/gradddev/hkow

include $(INCLUDE_DIR)/package.mk
include $(TOPDIR)/feeds/packages/lang/golang/golang-package.mk

define Package/hkow
	SECTION=:net
 	CATEGORY:=Network
	TITLE:=hkow
	URL:=https://github.com/gradddev/hkow
	DEPENDS:=$(GO_ARCH_DEPENDS)
endef

define Package/hkow/description
The HomeKit bridge for the OpenWrt

endef

define Package/hkow/install
	$(call GoPackage/Package/Install/Bin,$(1))

	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/hkow $(1)/etc/init.d/hkow
endef

$(eval $(call GoBinPackage,hkow))
$(eval $(call BuildPackage,hkow))
