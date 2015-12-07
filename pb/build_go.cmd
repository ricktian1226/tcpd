set PROTOC=protoc
set SOURCE=.\
set TARGET=..\src\feiyu.com\protocol
%PROTOC% --go_out=%SOURCE% --proto_path=. .\*.proto
copy /y %SOURCE%\*.go %TARGET%