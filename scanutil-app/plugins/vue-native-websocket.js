// @/plugins/vue-native-websocket.js
import Vue from 'vue'
import VueNativeSock from 'vue-native-websocket'

export default ({ store }, inject) => {
  Vue.use(VueNativeSock, 'ws://localhost:8080/ws', {
    format: 'json',
    reconnection: true,
    reconnectionAttempts: 5,
    reconnectionDelay: 3000,
    store,
    connectManually: false
  })
}
