# [WIP] [DRAFT] USER API DEFINITION

|ID|Method|URL|Permission|Explain|
|---|---|---|---|---|
|10101|GET|/get/user/:id/info|11111|Get user.User|
|10102|GET|/get/user/:id/basic|11111|Get user.Basic|
|10103|GET|/get/user/:id/profile|11111|Get user.Profile|
|10104|GET|/get/user/:id/sudo/auth|00101|Get user.Auth|
|20101|POST|/update/user/:id/username|00101|Update user.User.Username|
|20102|POST|/update/user/:id/avatar|01101|Update user.User.Avatar|
|20103|POST|/update/user/:id/profile|01101|Update user.Profile|
|20104|POST|/update/user/:id/sudo/email|00101|Update user.Auth.Email|
|20105|POST|/update/user/:id/sudo/password|00101|Update user.Auth.Password|
|20106|POST|/update/user/:id/sudo/security_email|00101|Update user.Auth.SecurityEmail|
|20107|POST|/update/user/:id/mod/locked|00011|Update user.User.Locked and <br/> Update user.User.LockedTill|
|20108|POST|/update/user/:id/admin/disabled|00011|Update user.User.Disabled|
|20109|POST|/update/user/:id/admin/electrons|00001|Update user.User.Electrons|
|20110|POST|/update/user/:id/admin/mod|00001|Update user.User.Mod|
|30101|DELETE|/delete/user/:id/sudo/user|00101|Delete user(multi actions)|
## /get/user/:id/
### [10101] GET /get/user/:id/info
|Role|Access|
|---|---|
|User|true|
|Self|true|
|Self(sudo)|true|
|Mod|true|
|Admin|true|

### [10102] GET /get/user/:id/basic
|Role|Access|
|---|---|
|User|true|
|Self|true|
|Self(sudo)|true|
|Mod|true|
|Admin|true|

### [10103] GET /get/user/:id/profile
|Role|Access|
|---|---|
|User|true|
|Self|true|
|Self(sudo)|true|
|Mod|true|
|Admin|true|
### [10104] GET /get/user/:id/sudo/auth
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|


## /update/user/:id/
### [20101] POST /update/user/:id/username
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### [20102] POST /update/user/:id/avatar
|Role|Access|
|---|---|
|User|false|
|Self|true|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### [20103] POST /update/user/:id/profile
|Role|Access|
|---|---|
|User|false|
|Self|true|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### [20104] POST /update/user/:id/sudo/email
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### [20105] POST /update/user/:id/sudo/password
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### [20106] POST /update/user/:id/sudo/security_email
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

### [20107] POST /update/user/:id/mod/locked
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|false|
|Mod|true|
|Admin|true|

### [20108] POST /update/user/:id/admin/disabled
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|false|
|Mod|true|
|Admin|true|

### [20109] POST /update/user/:id/admin/electrons
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|false|
|Mod|false|
|Admin|true|

### [20110] POST /update/user/:id/admin/mod
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|false|
|Mod|false|
|Admin|true|

## POST /create/user/
### POST /create/user/signup
### POST /create/user/user
### POST /create/user/forget

## DELETE /delete/user/:id/
### [40101] DELETE /delete/user/:id/sudo/user
|Role|Access|
|---|---|
|User|false|
|Self|false|
|Self(sudo)|true|
|Mod|false|
|Admin|true|

