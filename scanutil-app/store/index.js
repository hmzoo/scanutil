export const state = () => ({
  SocketStatus: 'offline',
  Scans: [],
  Models: ["AT-FS980M 48P","AT-FS980M 24P","AT-x930"],
  currentScan:{mod:"",inv:"",ser:"",mac:""},
  currentScanFilled:false
})

export const mutations = {
  addScan (state, scan) {
    for(var n=0;n<state.Scans.length;n=n+1){
      if(state.Scans[n].inv===scan.inv){
        console.log("UPDATE SCAN LIST")
        state.Scans[n]=Object.assign(state.Scans[n],scan)
        return
      }
    }
    console.log("ADD SCAN TO LIST")
    state.Scans.push(scan)
    console.log(state.Scans)
  },
  updateCurrentScan (state, text) {
    console.log("UPDATE CURRENT SCAN")
    if(isInv(text)){
      state.currentScan=Object.assign(state.currentScan,{inv:text})
      return
    }
    if(isMac(text)){
      state.currentScan=Object.assign(state.currentScan,{mac:text})
      return
    }
    state.currentScan=Object.assign(state.currentScan,{ser:text})
  },
  updateCurrentModel (state, text) {
    console.log("UPDATE CURRENT MODEL")
    state.currentScan=Object.assign(state.currentScan,{mod:text})
    if(!state.Models.includes(text)){
    state.Models.push(text)
    }
  },
  clearCurrentScan(state) {
    console.log("CLEAR CURRENT SCAN")
      state.currentScan=Object.assign(state.currentScan,{inv:"",ser:"",mac:""})
  },
  checkCurrentScan(state) {
    console.log("CHECK CURRENT SCAN")
      state.currentScanFilled=(state.currentScan.inv !="" && state.currentScan.ser !="" && state.currentScan.mac !="" && state.currentScan.mod !="")
  },
  SOCKET_ONOPEN (state, payload) {
    state.SocketStatus = 'online'
  },
  SOCKET_ONCLOSE (state, payload) {
    state.SocketStatus = 'offline'
  },
  SOCKET_ONERROR (state, event) {
    console.error(state, event)
  },
  SOCKET_ONMESSAGE (state, msg) {
    console.log('MSG', msg)
    //const msg = JSON.parse(message)
    if (msg.item && msg.tp === 'ADD') {
      let scan = msg.item
      for(var n=0;n<state.Scans.length;n=n+1){
        if(state.Scans[n].inv===scan.inv){
          state.Scans[n]=Object.assign(state.Scans[n],scan)
          return
        }
      }
      state.Scans.push(scan)
    }
    if (msg.data && msg.tp === 'DATA') {
      Object.assign(state.Scans,[])
      for(var n=0;n<msg.data.length;n=n+1){
        state.Scans.push(msg.data[n])
      }

    }
    if (msg.ct && msg.tp === 'INF') {
      state.Infos = JSON.parse(msg.ct)
      console.log(msg.ct)
      console.log('INF ', state.Infos)
    }
  }

}


//Utils

function isInv(t) {
return (t.length == 6 ) && (!isNaN(parseInt(t)))
}

function isMac(t) {
var a = parseInt(t,16);
return (t.length == 12) && (a.toString(16) === t.toLowerCase())
}
