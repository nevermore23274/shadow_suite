# Shadow Suite
This application is written in Go using the Fyne framework, and encompasses multiple functions primarily used in the security industry.

# Docker
To build the docker container, this is the one command that should be needed:
```
docker build -t go_network_tool .
```

# To Help With Development
- The easiest way to do this is using the docker container to save yourself from all the configurations. After simply building the container, you can use the Dev Containers and Remote Development extensions.
- Press F1 for the drop down menu and type in "Dev Containers: Open Folder in Container"
- It should open in the Root directory (you can see all the folders and files), so just press Open
- Choose "From a predefined container configuration definition..."
- Choose "Debian"
- Choose "bullseye"
- Type in "Go" and check the box, then type in "Code Runner" and check the box, then hit "ok"
- Hit "keep defaults"
- I would also advise you to hit the "show log" in the bottom right corner so you're able to see as the container finishes building.
- After it's done wit the building, install the extension "Code Runner" to be able to run the code from the IDE with a play button in the top right

# Updating Go
If you're in need of updating Go, here are some instructions assuming you're on Debian:
### 1. Uninstall the exisiting version
As mentioned [here](https://golang.org/doc/install#install), to update a go version you will first need to uninstall the original version.
```
$ sudo rm -rf /usr/local/go
```

### 2. Install the new version
Go to the [downloads](https://golang.org/dl/) page and download the binary release suitable for your system.

### 3. Extract the archive file
To extract the archive file:
```
$ sudo tar -C /usr/local -xzf /home/nikhita/Downloads/go1.8.1.linux-amd64.tar.gz
```

### 4. Make sure that your PATH contains `/usr/local/go/bin`
```
$ echo $PATH | grep "/usr/local/go/bin"
```
