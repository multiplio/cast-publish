# publish

[![Build Status](https://travis-ci.org/multiplio/cast-publish.svg?branch=master)](https://travis-ci.org/multiplio/cast-publish)
[![Go Report Card](https://goreportcard.com/badge/github.com/multiplio/cast-publish)](https://goreportcard.com/badge/github.com/multiplio/cast-publish)

## routes
| method | route | success | failure | comment |
|:---:|:---|:---|:---|---:|
| GET | /ready | 200 'ok' | - | kubernetes ready probe |
| GET | /twitter/\<user\>/\<post\> | 200 | 400 \<error\> | submit a post for publish to twitter |

## environment
```
PROGRAM_ALIAS=publish
ADDRESS=:3000

POST_URL=

TWITTER_CONSUMER_KEY=
TWITTER_CONSUMER_SECRET=

DATABASE_NAME=
DATABASE_PROTOCOL=
DATABASE_ADDRESS=
DATABASE_OPTIONS=
DATABASE_USER=
DATABASE_PASSWORD=
```

