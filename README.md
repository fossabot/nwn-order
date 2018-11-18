# nwn-order
Go program written to enhance nwnxee persistent worlds with an external go program.

![Image of Order](https://github.com/Urothis/nwn-order/blob/master/documentation/Screenshot.png)

> required software
> - Docker
> - Docker compose
>
> Module requirements
> - [Neverwinter Nights enhanced edition]:https://www.beamdog.com/products/neverwinter-nights-enhanced-edition
> - [NWNXEE]:https://nwnx.io/

## Setup help
This software is still early in development so the setup readme is a WIP. 

- Copy the /order folder to your current root folder.
- Compare the env files out of /config to your current docker setup
- Compare docker-compose.yml and add the order service to your current docker-compose file

The two best places to get started using this software:
> https://github.com/Urothis/nwnxee-template

Has a full docker structure with compose file to power it all. 
 
If you are still running into issues feel free to stop by the discord

> https://discord.gg/r6wuUdx

## Project goals
This project started as a way for me to setup a full ci deployment.
This has evolved into a fun way to see what we can do.
Community involvement is appreciated.   
    
>### Completed
- ~~Extra heartbeats~~ 
- ~~UUID generation~~ 

>### In Progress
- Full CI
- Cat facts
- Patreon integration
- Discord bot integration
- Whatever else the community suggests


## NWSCRIPT
### UUID
This function should be passed the player objet only.
It will return the assigned UUID to the player.

Internal scripts attach this uuid to the player tag.
> OrderGetUUIDPlayer(oPC)

Should return the players unique ID

Example return:
> 6fc7438a87d42b2dec552b4fb81b75a2

### Heartbeat
Heartbeat functionality can be enabled via config/nworder.env

>NWN_ORDER_HB_VERBOSE=

Setting to true will disply more logs for heartbeat

>NWN_ORDER_HB_ONE_MINUTE=true

Tickers will need to be enabled or disabled depending on your needs.

Default actions for heartbeat tickers are defined in order_heartbeat.nss

### CI/Github
This requires alittle bit of setup to function.

Requirements:
have a webhook setup for the repo you want to recieve alerts from.
https://developer.github.com/webhooks/creating/

When the docker-compose does go up, order will spit out an external facing IP and port. 

You will need to go into gitub and enable the webhook.

Example:
![Image of Github](https://github.com/Urothis/nwn-order/blob/master/documentation/Github_Screenshot.png)

So when you deliver a webhook, order will accept the webhook and trigger the 
"OrderGithub();"
function inside of order_github.nss
