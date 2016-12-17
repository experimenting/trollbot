# Trollbot

Use in Slack like `troll context` or `troll context @username` or `troll context @user1 @user2`  

```
liuggio 4:47 PM
 troll @liuggio shell
trollBOT 4:47 PM
 Hey @liuggio For me Shell is better then NodeJS
```

# Please contribute add troll to the YAML file

Add your troll [here](./feeds/feed.yml)

# Run 

    export TROLL_SLACK_TOKEN="xoxb-TOKEN-YOURTOKEN" && \
     go build . && \
     ./TrollBot feeds/feed.yml feeds/data.yml
