#
# Old school makefile, just type make on the command line
#

all: pinch test1 test2 done

GOSRCS = 					\
	pinch.go 				\
	src/go-pinch/pinch.go

TEST_URL  = http://peterhellberg.github.com/pinch/test.zip
TEST_FILE = data.json

pinch: $(GOSRCS)
	@echo =[COMPILE]======================================================
	GOPATH=`pwd` go build -o=pinch-test pinch.go

test1:
	@echo =[TEST 1]=======================================================
	./pinch-test $(TEST_URL)

test2:
	@echo =[TEST 2]=======================================================
	./pinch-test $(TEST_URL) $(TEST_FILE)

done:
	@echo =[DONE]=========================================================
	rm pinch-test

clean:
	rm -vf *~ */*~ */*/*~ pinch
