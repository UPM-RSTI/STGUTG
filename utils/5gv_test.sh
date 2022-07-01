#!/usr/bin/env bash

# Check OS

if [ -f /etc/os-release ]; then
    # freedesktop.org and systemd
    . /etc/os-release
    OS=$NAME
    VER=$VERSION_ID
else
    # Fall back to uname, e.g. "Linux <version>", also works for BSD, etc.
    OS=$(uname -s)
    VER=$(uname -r)
    echo "This Linux version is too old: $OS:$VER, we don't support!"
    exit 1
fi

sudo -v
if [ $? == 1 ]
then
    echo "Without root permission, you cannot run the test due to our test is using namespace"
    exit 1
fi


while getopts 'o' OPT;
do
    case $OPT in
        o) DUMP_NS=True;;
    esac
done
shift $(($OPTIND - 1))

TEST_POOL="TestInsertMongoDB|TestRegistration|TestGUTIRegistration|TestServiceRequest|TestXnHandover|TestN2Handover|TestDeregistration|TestPDUSessionReleaseRequest|TestPaging|TestNon3GPP|TestReSynchronisation"
if [[ ! "$1" =~ $TEST_POOL ]]
then
    echo "Usage: $0 [ ${TEST_POOL//|/ | } ]"
    exit 1
fi

cp config/smfcfg.conf config/test/smfcfg.test.conf

GOPATH=$HOME/go

if [ $OS == "Ubuntu" ]; then
    GOROOT=/usr/local/go
elif [ $OS == "Fedora" ]; then
    GOROOT=/usr/lib/golang
fi

PATH=$PATH:$GOPATH/bin:$GOROOT/bin

export GIN_MODE=release

cd src/upf/build
./bin/free5gc-upfd &

#cd ../../test
cd src/test
$GOROOT/bin/go test -v -vet=off -run $1


sleep 3
#sudo killall -15 free5gc-upfd
#sleep 1

cd ../..
mkdir -p testkeylog
for KEYLOG in $(ls *sslkey.log); do
    mv $KEYLOG testkeylog
done


sleep 2
sudo killall -15 free5gc-upfd
sleep 1
