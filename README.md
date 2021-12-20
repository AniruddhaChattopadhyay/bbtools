# BlueBeak Installation Guide

## Task 1: Building the CLI

The directory structure looks like

![Untitled](BlueBeak%20Installation%20Guide%20ff97e6ea437442f8b90b2eb638a38b89/Untitled.png)

### Fastest process:

1)cd into bbtools

2)cd into bbtools/bin

3)I have executables for linux arm64, windows amd64 and mac (darwin amd64),If you have any other system then skip this section and go to the next section.

Assuming you are on linux execute the command

```bash
./bbtools
```

![Untitled](BlueBeak%20Installation%20Guide%20ff97e6ea437442f8b90b2eb638a38b89/Untitled%201.png)

**Run install command:**

```bash
./bbtools install
```

**Run greet command:**

```bash
./bbtools greet hello ron
```

![Untitled](BlueBeak%20Installation%20Guide%20ff97e6ea437442f8b90b2eb638a38b89/Untitled%202.png)

### Building CLI binary using Go:

**Installing GO:**

This requires us to install go. There are a couple of ways to install go.

1) Downloading go installer from their official site:

[Download and install](https://go.dev/doc/install)

2) Using make file to install go:

I have also created a Makefile that installs and add the go to path for linux.

 

```bash
make
```

This installs go and then runs go install to create the binary which is automatically added to path.

Now we can be in any directory and use bbtools. (like run bbtools greet hello ron)

3) In ubuntu we can use snap to install go.

```bash
sudo apt update
sudo apt install snapd
sudo snap install go --classic
```

Check go installation using:

```bash
go version
```

**Building the binary:**

Now that go is installed we can build the executable. (No need to do this if installed using make as make already has built the binary executable)

```bash
cd bbtools
go install
bbtools install
bbtools greet hello ron
```

We can run bbtools from any directory once go install has been run. We don’t necessarily need to be inside bbtools directory.

## Task 2: Building nanovm using ops to run a R script

**Install ops:**

Refer to 

[Getting Started](https://docs.ops.city/ops/getting_started)

```bash
curl https://ops.city/get.sh -sSfL | sh
sudo apt-get install qemu-kvm qemu-utils
```

**Running hello.R script:**

```bash
cd bbtools
ops pkg load R_3.4.4 -c config.json
```

![Untitled](BlueBeak%20Installation%20Guide%20ff97e6ea437442f8b90b2eb638a38b89/Untitled%203.png)

Thus we get the desired output

### Side note:

I was ideally supposed to pull the docker image from local registry and then run the hello.R using that but despite trying several different base images that didn’t happen. I also spoke to the main collaborator and creater and CEO of nanovms and he said that there is an issue with Rscript specifically that deters the run of a R script using the Rscript command. I am attaching everything I did in the page below.

[ops from docker](https://www.notion.so/ops-from-docker-f8059a27626149f9a25da179c8481295)