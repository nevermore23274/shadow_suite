# Specifies base image
FROM python:3.11
# Install nmap
RUN apt-get update && apt-get install -y nmap
# Set working directory inside container to /app (subsequent commands are executed here)
WORKDIR /app
# Copies requirements.txt file from host systme into current directory (./) in the image
COPY requirements.txt ./requirements.txt
# Upgrades pip if needed
RUN /usr/local/bin/python3 -m pip install --upgrade pip
# Installs packages from requirements.txt
RUN pip3 install -r requirements.txt
# Informs docker that container will listen on this port, and makes it accessible outside container
EXPOSE 8501
# This tells Docker to execute 'streamlit run' which starts the Streamlit server
ENTRYPOINT ["streamlit", "run"]
# Specifies default command to be executed when the container starts
CMD ["Home.py"]