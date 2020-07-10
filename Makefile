arm:
	GOARCH=arm go build -o redirector-arm -v -a

clean:
	rm -f redirector-*
