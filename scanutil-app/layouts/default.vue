<template>
  <v-app dark>

    <v-app-bar
      fixed
      app
    >
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <SocketInfo />

    </v-app-bar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>


  </v-app>
</template>

<script>
import SocketInfo from '~/components/SocketInfo.vue'
export default {
  components: {
    SocketInfo
  },
  data () {
    return {
      title: 'ScanUtil'
    }
  }
}
</script>
