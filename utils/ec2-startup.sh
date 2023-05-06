sudo yum update -y
sudo yum install -y golang

echo -e "export GOROOT=/usr/lib/golang\nexport GOPATH=$HOME/projects\nexport PATH=$PATH:$GOROOT/bin\n" >> multiple.txt

source ~/.bash_profile

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

aws configure
gets the secrets from env variables in mac/goland

install google-cloud-sdk

exit 0


brew install protobuf
protoc --version

brew install grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"

### to generate both client and server grpc code
 PATH="${PATH}:${HOME}/go/bin" protoc services/user.proto \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --proto_path=.

###### Deploy app to ec2 instance ######

##copy zip from local machine to ec2 instance
scp -r -i SwipeMeter/ec2keypair.pem ./SwipeMeter.zip ec2-user@ec2-18-191-253-123.us-east-2.compute.amazonaws.com:/home/ec2-user/

##To run app on ec2 as a service using systemd (so it wont stop on exiting)
1. create a .service file in /etc/systemd/system/
2. After creating/modifying any file run ```sudo systemctl daemon-reload```
3. sudo systemctl start SwipeMeter
4. sudo systemctl status SwipeMeter
5. sudo systemctl enable SwipeMeter [Enables your app to start on machine boot]


