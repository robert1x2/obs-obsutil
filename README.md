OBS Util
========

Documentation
-------------

How to install on linux (CentOS 8 used):

1) **Download and Install**

Find binaries [here](https://github.com/opentelekomcloud/obs-obsutil/raw/master/release/5.1.15.zip).

In target ECS (e.g. Cent OS 8) downloaded it using command:
```
linux@ssh ]$ mkdir obsutil
linux@ssh ]$ wget https://github.com/opentelekomcloud/obs-obsutil/raw/master/release/5.1.15.zip
linux@ssh ]$ unzip 5.1.15.zip
linux@ssh ]$ cd 5.1.15
linux@ssh ]$ tar xf obsutil_linux_amd64_5.1.15.tar.gz
linux@ssh ]$ cd obsutil_linux_amd64_5.1.15
```
2) **Configure**

In current directory run obsutil without any args the first time. It will place ~/.obsutilconfig in your home directory.

```
linux@ssh ]$ ./obsutil
```

Open the config file with your prefered text editor (I like nano) and add ENDPOINT, AK and SK.

```
linux@ssh ]$ nano ~/obsutilconfig
```

OBS Endpoint you find [here](https://github.com/opentelekomcloud/obs-obsutil/blob/master/User%20Guide.pdf) which is obs.eu-de.otc.t-systems.com
How to create and optain AK/SK you find here:

Creating an AK and SK
[direct link](https://docs.otc.t-systems.com/sdk-java-devg/obs/en-us_topic_0142800555.html) or [OBS user guide](https://docs.otc.t-systems.com/obs/doc/download/pdf/obs-usermanual.pdf) search for "creating an AK and SK".

First lines of .obsutilconfig look like:
```
endpoint=http://obs.eu-de.otc.t-systems.com
ak=*** Provide your Access Key ***
sk=*** Provide your Secret Key ***
```

3) **Run obsutil**

List all buckets:
```
./obsutil ls
```

Copy file setup.sh from current directory to OBS bucket t-systems
```
./obsutil cp ./setup.sh obs://t-systems
```

Further Installation Notes
--------------------------
Version 5.1.15

Dependent SDK Version: 
Go SDK 2.2.12

Install command:

ARM:
set GOPATH=<your go path>
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go install -ldflags "-X main.AesKey=<your aes key of which the length must be 16> -X main.AesIv=<your aes iv of which the length must be 16> -X main.CloudType=dt" obsutil

Linux:
set GOPATH=<your go path>
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go install -ldflags "-X main.AesKey=<your aes key of which the length must be 16> -X main.AesIv=<your aes iv of which the length must be 16> -X main.CloudType=dt" obsutil


Windows:
set GOPATH=<your go path>
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go install -ldflags "-X main.AesKey=<your aes key of which the length must be 16> -X main.AesIv=<your aes iv of which the length must be 16> -X main.CloudType=dt" obsutil

MacOs:
set GOPATH=<your go path>
set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=amd64
go install -ldflags "-X main.AesKey=<your aes key of which the length must be 16> -X main.AesIv=<your aes iv of which the length must be 16> -X main.CloudType=dt" obsutil





