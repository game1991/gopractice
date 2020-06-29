
GCC=gcc

build:
	$(GCC) -o libdencrypt.so -fPIC -shared dencrypt.c md5.c
	$(GCC) -c dencrypt.c md5.c
	ar r libdencrypt.a dencrypt.o md5.o

test:
	$(GCC) -o dencrypt_test *.c
	time ./dencrypt_test
