.PHONY: dummy build clean

ifndef DEST
override DEST = destination.txt
endif

# When unintentionally running `make` alone without argument, it will trigger 
# harmless first target and prevent other unnecessary target (e.g., build) 
dummy:
	echo ""

build:
	go mod tidy && \
	go build -o resumable-copy .

clean:
	rm -f $(DEST)
