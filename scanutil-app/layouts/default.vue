<template>
  <v-app dark>

    <v-app-bar
      :clipped-left="clipped"
      fixed
      app
    >
      <v-toolbar-title v-text="title" />
      <v-spacer />

    </v-app-bar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>


  </v-app>
</template>

<script>
export default {
  data () {
    return {
      title: 'ScanUtil'
    }
  }
}
</script>
