# About TuruBot
TuruBot is telegram bot program that made with Go.

This bot has a lot of feature, we will update it ASAP.

## Installation Tutorial
1. **Install Go compiler**
   - go to [https://go.dev/](https://go.dev) and download the latest golang version

2. **Clone this repository**
    ```bash
   git clone https://github.com/n0paleon/turubot
   cd turubot
   
3. **Compile source code**
    ```bash
   cd scripts
   make

4. **Set your bot token**
    - rename file `config_examples.yaml` to `config.yaml`
    - go to [@BotFather](https://t.me/BotFather) at Telegram and make your own bot token
    - open `config.yaml` then put your own bot token

5. **Start the bot**
    - check the `bin` directory after doing step-1, make sure there is a file named `turubot`
    - run this file using command `./bin/turubot` from the root project
