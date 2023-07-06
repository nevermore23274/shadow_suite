# Shadow Suite
This application is written in Python using the Streamlit framework. This tool is mostly a pet project, but the intent is for it to encompass multiple functions primarily used in the security industry.

# Docker
To build the docker container, this is the one command that should be needed:
```
docker build -t net_tool:latest -f Docker/Dockerfile .
```
And run it with:
```
docker run -p 8501:8501 -v <path_to_folder>:/app net_tool:latest
```
After making changes, you'll need to kill the container and rerun it in order for the updates to take. After you're done, I would suggest pruning in order to clear up space made by this process:
```
docker system prune -a
```

Then in your browser you can navigate to 127.0.0.1:8501 in order to have the app open.