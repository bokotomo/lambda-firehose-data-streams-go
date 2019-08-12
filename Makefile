create:
	GOOS=linux GOARCH=amd64 go build -o handlerFirehose
	zip handlerFirehose.zip ./handlerFirehose

