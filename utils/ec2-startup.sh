sudo yum update -y
sudo yum install -y golang

echo -e "export GOROOT=/usr/lib/golang\nexport GOPATH=$HOME/projects\nexport PATH=$PATH:$GOROOT/bin\n" >> multiple.txt

source ~/.bash_profile

curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install