# lottery

**lottery** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

** Do not `npm install` in Vue
This application normally runs on gitpod, not local. We already changed many files in vue/node_modules.

-- Directory 
  - vue 
    contain Frontend UI/logic
  - x/lottery 
    - Bussiness logic link to Frontend
  - x/lottery/handleMsgCreate*.go
    - hadler Controller for handleMsgCreate*.go
    - handleRestMsg From Vue
    
![alt text](https://github.com/ppalineee/lottery/tree/main/pic/announce.png)

![alt text](https://github.com/ppalineee/lottery/tree/main/pic/announcelist.png)

