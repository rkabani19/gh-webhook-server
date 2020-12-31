# GitHub Webhook Server
Skeleton server to launch workflows triggered by GitHub webhook events.

## Local Development
GitHub requires a publically reachable address so it can send out messages. Use [ngrok](https://ngrok.com/) to create a forwarding address for your local host. 

```
ngrok http 8080
```

Configure webhooks on your repository as explained [here](https://docs.github.com/en/free-pro-team@latest/developers/webhooks-and-events/configuring-your-server-to-receive-payloads).

Clone this repository and `cd` into it.
```
git clone git@github.com:rkabani19/gh-webhook-server.git && cd gh-webhook-server
```

Start the server
```
go run main.go
```

When a webhook event is triggered the server will execute your defined action.
