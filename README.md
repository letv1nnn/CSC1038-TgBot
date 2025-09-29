# CSC1038 LabBot

LabBot is your personal Computer Science study assistant on Telegram! It helps students quickly find useful resources and provide solutions to CSC1038 lab problems when you need them.

*this project was created as a practice task after learning golang for one week.*

Only docker image available. In this week gonna deploy it on a local Kubernetes cluster.


### Build and Run Locally
```sh
git clone https://github.com/letv1nnn/CSC1038-TgBot.git && cd CSC1038-TgBot
docker build -t csc1038bot:latest .
docker run --rm -it -e TOKEN="" csc1038bot:latest       # put telegram token in the placeholder
```