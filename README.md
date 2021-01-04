# Log

## Issue #1: cannot find runtime/cgo for when building for linux_arm64 + DefaultStatic=True

```shell
% uname -a
Linux arch01 5.10.3-arch1-1 #1 SMP PREEMPT Sun, 27 Dec 2020 10:50:46 +0000 x86_64 GNU/Linux

% go env | grep -E 'GO(HOST)?(ARCH|OS)'
GOARCH="amd64"
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"

% cat .plzconfig | grep -v ';'

[go]
GoTool = //tools/go:toolchain|go
importpath = github.com/ekhabarov/plz-crosscompile-issue
DefaultStatic = True

[buildconfig]
default-docker-repo = hub.docker.com

% plz clean
Cleaning in background; you may continue to do pleasing things in this repo in the meantime.

% plz run //app:image
Sending build context to Docker daemon     153B
Step 1/1 : FROM golang:1.15-alpine3.12
 ---> 1463476d8605
Successfully built 1463476d8605
Successfully tagged hub.docker.com/golang:7f7a30aa28b42aea05146be9827d27bb57268bc34526b3d08372dd7f9275c0e6
Sending build context to Docker daemon  2.263MB
Step 1/5 : FROM hub.docker.com/golang:7f7a30aa28b42aea05146be9827d27bb57268bc34526b3d08372dd7f9275c0e6
 ---> 1463476d8605
Step 2/5 : WORKDIR /app
 ---> Using cache
 ---> 896ec2798d58
Step 3/5 : COPY . .
 ---> Using cache
 ---> 1a47f9b6e0db
Step 4/5 : ENTRYPOINT ["/app/bin"]
 ---> Using cache
 ---> e42104c05a4c
Step 5/5 : EXPOSE 8080
 ---> Using cache
 ---> 4d0970924fcf
Successfully built 4d0970924fcf
Successfully tagged hub.docker.com/app_image:image

% plz run //app:image --arch linux_amd64
14:27:01.639 CRITICAL: Target //app:image not found in build graph

% plz run //app:image --arch linux_arm64
Build stopped after 24.27s. 1 target failed:
    ///linux_arm64//app:bin
Error building target ///linux_arm64//app:bin: exit status 1
loadinternal: cannot find runtime/cgo
/home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/bin/tools/go/toolchain/pkg/tool/linux_amd64/link: running /usr/bin/gcc failed: exit status 1
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: Relocations in generic ELF (EM: 183)
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: Relocations in generic ELF (EM: 183)
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: Relocations in generic ELF (EM: 183)
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: Relocations in generic ELF (EM: 183)
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: Relocations in generic ELF (EM: 183)
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: Relocations in generic ELF (EM: 183)
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: Relocations in generic ELF (EM: 183)
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: Relocations in generic ELF (EM: 183)
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: Relocations in generic ELF (EM: 183)
/usr/bin/ld: /home/ekhabarov/develop/go/src/github.com/ekhabarov/plz-crosscompile-issue/plz-out/tmp/linux_arm64/app/bin._build/go.o: error adding symbols: file in wrong format
collect2: error: ld returned 1 exit status
```

## Issue #2: target not found in build graph when building for linux_arm64 + DefaultStatic=False

```shell
% uname -a
Linux arch01 5.10.3-arch1-1 #1 SMP PREEMPT Sun, 27 Dec 2020 10:50:46 +0000 x86_64 GNU/Linux

% go env | grep -E 'GO(HOST)?(ARCH|OS)'
GOARCH="amd64"
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"

% plz clean
Cleaning in background; you may continue to do pleasing things in this repo in the meantime.
% cat .plzconfig | grep -v ';'

[go]
GoTool = //tools/go:toolchain|go
importpath = github.com/ekhabarov/plz-crosscompile-issue

[buildconfig]
default-docker-repo = hub.docker.com

% plz run //app:image
Sending build context to Docker daemon     153B
Step 1/1 : FROM golang:1.15-alpine3.12
 ---> 1463476d8605
Successfully built 1463476d8605
Successfully tagged hub.docker.com/golang:7f7a30aa28b42aea05146be9827d27bb57268bc34526b3d08372dd7f9275c0e6
Sending build context to Docker daemon  1.873MB
Step 1/5 : FROM hub.docker.com/golang:7f7a30aa28b42aea05146be9827d27bb57268bc34526b3d08372dd7f9275c0e6
 ---> 1463476d8605
Step 2/5 : WORKDIR /app
 ---> Using cache
 ---> 896ec2798d58
Step 3/5 : COPY . .
 ---> Using cache
 ---> bf69a442fc8b
Step 4/5 : ENTRYPOINT ["/app/bin"]
 ---> Using cache
 ---> cbb56fc1b965
Step 5/5 : EXPOSE 8080
 ---> Using cache
 ---> 423fbbe9f194
Successfully built 423fbbe9f194
Successfully tagged hub.docker.com/app_image:image

% plz run //app:image --arch linux_amd64
15:08:35.102 CRITICAL: Target //app:image not found in build graph

% plz run //app:image --arch linux_arm64
15:08:55.655 CRITICAL: Target //app:image not found in build graph
```

## Issue #3: using --arch flag with a host arch value and without specifying tools arch.

```shell
% uname -a
Linux arch01 5.10.3-arch1-1 #1 SMP PREEMPT Sun, 27 Dec 2020 10:50:46 +0000 x86_64 GNU/Linux

% go env | grep -E 'GO(HOST)?(ARCH|OS)'
GOARCH="amd64"
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"

% plz run //app:image
Sending build context to Docker daemon     153B
Step 1/1 : FROM golang:1.15-alpine3.12
 ---> 1463476d8605
Successfully built 1463476d8605
Successfully tagged hub.docker.com/golang:7f7a30aa28b42aea05146be9827d27bb57268bc34526b3d08372dd7f9275c0e6
Sending build context to Docker daemon  2.263MB
Step 1/5 : FROM hub.docker.com/golang:7f7a30aa28b42aea05146be9827d27bb57268bc34526b3d08372dd7f9275c0e6
 ---> 1463476d8605
Step 2/5 : WORKDIR /app
 ---> Using cache
 ---> 896ec2798d58
Step 3/5 : COPY . .
 ---> Using cache
 ---> 1a47f9b6e0db
Step 4/5 : ENTRYPOINT ["/app/bin"]
 ---> Using cache
 ---> e42104c05a4c
Step 5/5 : EXPOSE 8080
 ---> Using cache
 ---> 4d0970924fcf
Successfully built 4d0970924fcf
Successfully tagged hub.docker.com/app_image:image

% plz run //app:image --arch linux_amd64
14:32:19.608 CRITICAL: Target //app:image not found in build graph

% plz run  ///linux_amd64//app:image
Sending build context to Docker daemon     153B
Step 1/1 : FROM golang:1.15-alpine3.12
<skipped>

% plz run ///linux_amd64//app:image --arch linux_amd64
Sending build context to Docker daemon     153B
Step 1/1 : FROM golang:1.15-alpine3.12
<skipped>
```

