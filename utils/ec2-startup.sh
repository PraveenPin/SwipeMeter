sudo yum update -y
sudo yum install -y golang

echo -e "export GOROOT=/usr/lib/golang\nexport GOPATH=$HOME/projects\nexport PATH=$PATH:$GOROOT/bin\n" >> multiple.txt

source ~/.bash_profile

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

brew install protobuf
protoc --version

brew install grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"

exit 0
### to generate both client and server grpc code
 PATH="${PATH}:${HOME}/go/bin" protoc services/user.proto \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --proto_path=.
