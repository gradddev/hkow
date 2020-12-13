FROM openwrtorg/sdk:ramips-mt7621-19.07.4
RUN cp feeds.conf.default feeds.conf
RUN sed '/^src-git luci/d' -i feeds.conf
RUN sed '/^src-git routing/d' -i feeds.conf
RUN sed '/^src-git telephony/d' -i feeds.conf
RUN ./scripts/feeds update -a
RUN ./scripts/feeds install -a
RUN make defconfig
RUN sed 's/CONFIG_SIGNED_PACKAGES=y/CONFIG_SIGNED_PACKAGES=n/' -i .config
COPY Makefile package/hkow/
COPY files package/hkow/files
