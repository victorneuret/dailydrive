
# Daily Drive Music

Self hosted Spotify "Daily Drive" playlist creation and synchronization without the podcasts.

#### What is the "Daily Drive" playlist?
According to Spotify news article:
> It’s a new Spotify playlist that combines the very best of news talk shows, including the relevancy and personality of the hosts, with the best of audio streaming (on demand, personalized playing and discovery). It combines music you love with relevant, timely world updates from reputable sources – all put together in a seamless and unified listening experience. - [*Your Daily Drive: Music and News That’ll Brighten Your Commute*](https://newsroom.spotify.com/2019-06-12/your-daily-drive-music-and-news-thatll-brighten-your-commute/)

#### Motivations:
I like a lot the idea of a playlist mixing different type of music I'm listening to. Spotify already creates several daily mixes but each of them is focused on a specific type of music (pop, rock, etc..). "Daily Drive" is the first daily playlist mixing different music types and I found it pretty relative to my music taste. The negative point is that the playlist is meant to look like a radio, music divided by news and talk shows, but I do only want the music.

#### What is this repository doing?
This repository is a Go program meant to create a new playlist called "My Music Daily Drive" and sync it every day with the Spotify's "Daily Drive" playlist without the podcasts.

#### Result:
Spotify's "Daily Drive" playlist:
![image](https://user-images.githubusercontent.com/34629981/116953817-31de7a80-acc9-11eb-8945-1dc1714781d7.png)
Created and synced playlist without the podcasts:
![image](https://user-images.githubusercontent.com/34629981/116953874-520e3980-acc9-11eb-8779-f08890cae535.png)
****


## Install on your server

### Create Spotify application
- First of all, you must have a Spotify account.
- Go to https://developer.spotify.com/dashboard/applications and login to your account.
- Create a new app by clicking on `CREATE AN APP` button. Give it the name and description you want.
- Click on the `EDIT SETTINGS` button. Add a new `redirect URI`: http://{your-server-ip}:2021/callback or http://localhost:2021/callback if you want to use it locally on you computer.
- Click `SAVE`.
- On the dashboard, keep the displayed `Client ID` and `Client Secret` for later use.

### Deploy the service
- ssh to your server.
- Clone this repository: `git clone git@github.com:victorneuret/dailydrive.git`
- Move to the cloned repository `cd dailydrive`
- Create a `.env` file and add you client id, secret and redirect URL in it:
```
SPOTIFY_ID=YOUR_CLIENT_ID
SPOTIFY_SECRET=YOUR_CLIENT_SECRET
REDIRECT_URL=http://{your-server-ip}:2021/callback
```
- Build and launch the docker image `docker-compose up -d`
- Connect your Spotify account in the launched app by clicking the link in the log `docker logs -f daily-drive`
- Accept the connection on the web page
- The first sync should be executed and logged.
- You can hit CTRL + C to quit the logs.

Your new playlist should be created on your account:
![image](https://user-images.githubusercontent.com/34629981/116954883-145ee000-accc-11eb-9eb9-e7fb4bb6e1ef.png)
The update of the playlist will be automatically executed every day at midnight on your server time.



Enjoy! 🎉
