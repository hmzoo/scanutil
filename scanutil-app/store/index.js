export const state = () => ({
  SocketStatus: 'offline',
  Scans: [],
  Models: ["AT-FS980M 48P","AT-FS980M 24P","AT-x930"]
})

export const mutations = {
  addScan (state, scan) {
    for(var n=0;n<state.Scans.length;n=n+1){
      if(state.Scans[n].inv===scan.inv){
        console.log("UPDATED")

        state.Scans[n]=Object.assign(state.Scans[n],scan)
        return
      }
    }
    state.Scans.push(scan)
  },
  addModel (state, text) {
    if(!state.Models.includes(text)){
    state.Models.push(text)
    }
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
    // const msg = JSON.parse(message)
    if (msg.ct && msg.tp === 'MSG') {
      state.MsgList.push(msg.ct)
    }
    if (msg.ct && msg.tp === 'INF') {
      state.Infos = JSON.parse(msg.ct)
      console.log(msg.ct)
      console.log('INF ', state.Infos)
    }
  }

}
