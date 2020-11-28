# lottery

**lottery** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

___

** Do not `npm install` in Vue
This application normally runs on gitpod, not local. We already changed many files in vue/node_modules.

___

-- Directory 
lottery/
├─ .pi/
├─ .vscode/
├─ app/
├─ cmd/
├─ node_module/
├─ pic/
├─ vue/
│  ├─ node_modules/
│  ├─ public/
│  ├─ src/
├─ x/
│  ├─ lottery/
│  │  ├─ client/
│  │  ├─ keeper/
│  │  ├─ types/
│  │  ├─ spec/

___

## Use mnemonic to sigin
<img src="https://github.com/ppalineee/lottery/blob/main/pic/signin.png" width="800">

## To create lottery
- name: string
- detail: string
- reward: 100token
- price: 10token
<img src="https://github.com/ppalineee/lottery/blob/main/pic/createlottery.png" width="800">

### If everything works correctly: the lottery will show in lottery list
<img src="https://github.com/ppalineee/lottery/blob/main/pic/lotterylist.png" width="800">

## Go to /ticket to Buy ticket by using id and number
- lotteyID: string
- number: 0-9
- time.Now() < drawdate
<img src="https://github.com/ppalineee/lottery/blob/main/pic/buyticket.png" width="800">

### If everything works correctly: the tickets will show in ticketlist
<img src="https://github.com/ppalineee/lottery/blob/main/pic/ticketlist.png" width="800">


## Go to /announce and paste id lottery to announce lottery
- announce by lotteryID
<img src="https://github.com/ppalineee/lottery/blob/main/pic/announce.png" width="800">

### If everything works correctly: the announce will show in announcelist
- time.Now() > drawdate
<img src="https://github.com/ppalineee/lottery/blob/main/pic/announcelist.png" width="800">

