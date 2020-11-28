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

- use mnemonic to sigin
<img src="https://github.com/ppalineee/lottery/blob/main/pic/signin.png" width="800">

- go to /announce and paste id lottery to announce lottery
<img src="https://github.com/ppalineee/lottery/blob/main/pic/announce.png" width="800">

<img src="https://github.com/ppalineee/lottery/blob/main/pic/announcelist.png" width="800">

<img src="https://github.com/ppalineee/lottery/blob/main/pic/buyticket.png" width="800">

<img src="https://github.com/ppalineee/lottery/blob/main/pic/createlottery.png" width="800">

<img src="https://github.com/ppalineee/lottery/blob/main/pic/lotterylist.png" width="800">

<img src="https://github.com/ppalineee/lottery/blob/main/pic/ticketlist.png" width="800">

