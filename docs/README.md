# Documentation
These are just the bigger picture of available endpoints, some of these are requiring a request body and query params. So if you want to test the endpoints, refer to import JSON collections on this folder to Postman or the online ver:
[https://www.postman.com/kresnawan1/bcc-canteen](https://www.postman.com/kresnawan1/bcc-canteen)

```
GET    /api/v1/user/profile             --> 
PATCH  /api/v1/user/profile             --> 
PATCH  /api/v1/user/:uid                --> 
GET    /api/v1/user                     --> 
DELETE /api/v1/user/:uid                -->

POST   /api/v1/auth/login               --> 
POST   /api/v1/auth/register            --> 
DELETE /api/v1/auth/logout              --> 
GET    /api/v1/auth/token               -->

GET    /api/v1/order/:oid               --> 
GET    /api/v1/order                    --> 
POST   /api/v1/order                    --> 
POST   /api/v1/order/:oid/pay           --> 
POST   /api/v1/order/:oid/feedback      --> 
PATCH  /api/v1/order/:oid               -->

GET    /api/v1/canteen                  --> 
GET    /api/v1/canteen/:cid/menu        --> 
GET    /api/v1/canteen/:cid/feedback    -->
GET    /api/v1/canteen/my               -->
POST   /api/v1/canteen/:cid/menu        --> 
GET    /api/v1/canteen/:cid/order       -->
DELETE /api/v1/canteen/feedback/:fid    --> 
POST   /api/v1/canteen                  --> 
DELETE /api/v1/canteen/:cid             --> 
GET    /api/v1/canteen/owner            --> 
POST   /api/v1/canteen/:cid/owner       --> 
DELETE /api/v1/canteen/:cid/owner       --> 
GET    /api/v1/canteen/:cid/owner       -->

GET    /api/v1/cart                     --> 
POST   /api/v1/cart                     --> 
DELETE /api/v1/cart                     --> 

DELETE /api/v1/feedback/:fid            --> 

GET    /api/v1/menu                     --> 
PATCH  /api/v1/menu/:mid                --> 
PUT    /api/v1/menu/:mid                --> 
DELETE /api/v1/menu/:mid                --> 

GET    /api/v1/owner                    --> 
POST   /api/v1/owner                    --> 
PATCH  /api/v1/owner                    --> 
PATCH  /api/v1/owner/:owid              --> 
```