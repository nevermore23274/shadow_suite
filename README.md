# Shadow Suite
This application will be written in Go using the Fyne framework, and basically encompass multiple functions primarily used in the security industry.

# Docker
To build the docker container, this is the one command that should be needed:
```
docker build -t go_network_tool .
```

# To Help With Development
- The easiest way to do this is using the docker container to save yourself from all the configurations. After simply building the container, you can use the Dev Containers extension.
- Press F1 for the drop down menu and type in "Dev Containers: Open Folder in Container"
- It should open in the Root directory (you can see all the folders and files), so just press Open
- Choose "From a predefined container configuration definition..."
- Choose "Debian"
- Choose "bullseye"
- Type in "Go" and check the box, the hit "ok"
- Hit "keep defaults"
- I would also advise you to hit the "show log" in the bottom right corner so you're able to see as the container finishes building.
- After it's done wit the building, install the extension "Code Runner" to be able to run the code from the IDE with a play button in the top right
