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

Accept: application/json, text/plain, */*
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJqYW4iLCJyb2xlIjoxLCJzY29wZSI6ImRlZmF1bHQiLCJleHAiOjE2OTg3NTIxNjV9.EQVwb6sqow1WX_7rbayVAgsUOwzRkgMqc3X2IENvs8k
Content-Type: application/x-tar
Dnt: 1
Referer: https://pietschcloud.de/
Sec-Ch-Ua: "Chromium";v="118", "Google Chrome";v="118", "Not=A?Brand";v="99"
Sec-Ch-Ua-Mobile: ?0
Sec-Ch-Ua-Platform: "Linux"
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36