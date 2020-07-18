# Ghost Session Lambda

![Go](https://img.shields.io/badge/Go-1.14-blue.svg?logo=go&longCache=true&logoColor=white&colorB=88C0D0&style=flat-square&colorA=4c566a)
![GitHub Last Commit](https://img.shields.io/github/last-commit/google/skia.svg?style=flat-square&colorA=4c566a&colorB=a3be8c)
[![GitHub Issues](https://img.shields.io/github/issues/toddbirchard/ghost-session-lambda.svg?style=flat-square&colorA=4c566a&colorB=ebcb8b)](https://github.com/toddbirchard/ghost-session-lambda/issues)
[![GitHub Stars](https://img.shields.io/github/stars/toddbirchard/ghost-session-lambda.svg?style=flat-square&colorB=ebcb8b&colorA=4c566a)](https://github.com/toddbirchard/ghost-session-lambda/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/toddbirchard/ghost-session-lambda.svg?style=flat-square&colorA=4c566a&colorB=ebcb8b)](https://github.com/toddbirchard/ghost-session-lambda/network)

Verifies whether the current user is a Ghost member and returns general account information upon success.

### Example Usage

**Request:**
```bash
$ curl https://hackersandslackers.app/.netlify/functions/ghost-session-lambda
```

**Response:**
```json
{
    "uuid": "df0c49c7-8e71-4e8f-9ab5-97082063c4cd",
    "email": "toddbirchard+jfphhudgshjdiwhekjb@gmail.com",
    "name": null,
    "firstname": null,
    "avatar_image": "https://gravatar.com/avatar/5f2c7111462afaecd8bc5a79d21dfe7e?s=250&d=blank",
    "subscribed": true,
    "subscriptions": [],
    "paid": false
}
```