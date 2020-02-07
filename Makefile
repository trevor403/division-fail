GBA    := main.elf
SRCS   := *.go

TINYGO_FLAGS :=

.PHONY : build clean $(GBA)

build : $(GBA)

clean :
	@rm -f $(GBA)

$(GBA) : *.go
	@docker run --rm -it -v $(PWD):/wk -w /wk tinygo/tinygo:0.12.0 tinygo build --target=gameboy-advance $(TINYGO_FLAGS) -o $@ .
