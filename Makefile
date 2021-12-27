# mclist
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: mclist

mclist:
	$(GO) build $(GOFLAGS)

clean:
	$(RM) mclist

install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f mclist $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/mclist

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/mclist

.DEFAULT_GOAL := all

.PHONY: all clean install uninstall
