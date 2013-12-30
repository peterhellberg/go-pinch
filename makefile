#
# Old school makefile, just type make on the command line
#

all: pinch test1 test2 done

GOSRCS = 					\
	main.go 				\
	pinch/pinch.go

TEST_URL  = http://peterhellberg.github.com/pinch/test.zip
TEST_FILE = data.json

pinch: $(GOSRCS)
	@echo =[COMPILE]======================================================
	go build -o=pinch-test

test1:
	@echo =[TEST 1]=======================================================
	./pinch-test $(TEST_URL)

test2:
	@echo =[TEST 2]=======================================================
	./pinch-test $(TEST_URL) $(TEST_FILE)

done:
	@echo =[DONE]=========================================================
	rm pinch-test
