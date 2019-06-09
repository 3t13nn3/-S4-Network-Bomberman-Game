#go makefile

GOCMD = go
GOCLEAN=$(GOCMD) clean
BUILDCMD = $(GOCMD) build
OUT = bomber
CC   = go build

all:
	$(BUILDCMD) -o $(OUT) -v

clean:  
	$(GOCLEAN)
	rm -f $(OUT)