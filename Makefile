SUBDIRS := src/app test

all:
	@for d in $(SUBDIRS); do \
	$(MAKE) -C $$d all; \
	done

	mv src/app/dstp-agent ./
	mv test/client-ctl ./

agent:
	@for d in src/app; do \
	$(MAKE) -C $$d agent; \
	done

client:
	@for d in test; do \
	$(MAKE) -C $$d client; \
	done

clean:
	@for d in $(SUBDIRS); do \
	$(MAKE) -C $$d clean; \
	done
