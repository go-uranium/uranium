# [WIP] [DRAFT] USER API DEFINITION

Var:

|Name|Value|Explanation|
|---|---|---|
|id|uid|id == 0 stands for current user(self)|
|username|username|username only|
|user|username or email|for login|


## GET /user/
### GET /user/webauthn/new

## POST /user/
### POST /user/login
### POST /user/signup/start
### POST /user/signup/finish
### POST /user/new
### POST /user/forget

## /user/:id/
### GET /user/:id/info
Permission:

|Role|Access|
|---|---|
|User|true|
|Self|true|
|Self(sudo)|true|
|Mod|true|
|Admin|true|

Success:
```json
{
	"uid": 6,
	"username": "Richard",
	"username_lowercase": "richard",
	"electrons": 8963,
	"mod": true,
	"admin": true,
	"created_at": "2021-08-19T03:46:14.188311914Z"
}
```

Failed:
```json
{
	"err": true,
	"msg": "User not found."
}
```
```json
{
	"err": true,
	"msg": "Invalid UID."
}
```
...


### GET /user/:id/basic
|Role|Access|
|---|---|
|User|true|
|Self|true|
|Self(sudo)|true|
|Mod|true|
|Admin|true|

### GET /user/:id/profile
|Role|Access|
|---|---|
|User|true|
|Self|true|
|Self(sudo)|true|
|Mod|true|
|Admin|true|

### GET /user/:user/auth_methods
Permission:

|Role|Access|
|---|---|
|User|true|
|Self|true|
|Self(sudo)|true|
|Mod|true|
|Admin|true|


Params:

|Key|Value|Explanation|
|---|---|---|
|type|user, mod, admin|auth type|


Success:
```json
{
	"password": true,
	"webauthn": true
}
```
```json
{
	"password": true,
	"webauthn": false
}
```
...

Failed:
```json
{
	"err": true,
	"msg": "User not found."
}
```
...

### GET /user/:id/sudo/auth
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### GET /user/:id/sudo/totp
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### GET /user/:id/sudo/webauthn
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

## GET /user/username/:username/
### GET /user/username/:username/info
### GET /user/username/:username/basic
### GET /user/username/:username/profile

## /user/:id/
### POST /user/:id/username
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### POST /user/:id/avatar
|Role|Access|
|---|---|
|User|false|
|Self|true|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### POST /user/:id/profile
|Role|Access|
|---|---|
|User|false|
|Self|true|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### POST /user/:id/sudo/email
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### POST /user/:id/sudo/password
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### POST /user/:id/sudo/security_email
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### POST /user/:id/sudo/totp

### POST /user/:id/sudo/totp

### POST /user/:id/mod/locked
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|false|
|Mod|true|
|Admin|true|

### POST /user/:id/admin/disabled
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|false|
|Mod|true|
|Admin|true|

### POST /user/:id/admin/electrons
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|false|
|Mod|false|
|Admin|true|

### POST /user/:id/admin/mod
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|false|
|Mod|false|
|Admin|true|

## DELETE /user/:id/
### DELETE /user/:id/sudo/user
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

