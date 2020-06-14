<template>
<div>
  <v-icon color="green" right v-if="connected" >mdi-check-network</v-icon>
  <v-icon color="red" right v-else >mdi-close-network</v-icon>
</div>
</template>
<script>

export default {

  data: () => ({
  }),
  methods: {
    sendcmd () {
    }
  },
  computed: {
    connected () {
      return this.$store.state.SocketStatus == "online"
    }
  }
}
</script>
