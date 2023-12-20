# silentparty-event

Ganzen Stack starten: ```yarn dev```
_Achtung: der Stack lässt sich nur mit installierter Go v1.21 und Node 18 | 20 starten. Falls du nur Frontend oder Backend schreiben willst, gern einfach nur den Teil starten._

## Backend: ```/PB-Database```
Das Backend ist in Go auf PocketBase geschrieben.

Kann mit ```yarn dev:pocketbase``` gestartet werden.


## Frontend: ```/SilentParty-Event```
Das Frontend ist eine Nuxt-Multipage-Application.

Kann mit ```yarn dev:nuxt``` gestartet werden.



# Features

Falls Nutzer ihre Tickets bearbeiten können sollen, fügen wir wieder diese Zeile in PocketBase ein
```
@request.data.sold:isset = false && @request.data.used:isset = false && 
@request.data.validated:isset = false
```

# Deployment
POST https://pietschcloud.de/api/endpoints/1/docker/images/load
