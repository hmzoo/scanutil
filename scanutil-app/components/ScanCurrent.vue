<template>
  <v-card outlined >
    <v-card-title>Fiche produit <v-icon color="green" right v-if="filled" >mdi-check-bold</v-icon></v-card-title>
    <v-card-text>
      <v-simple-table dense>
        <tbody>
    <tr>
      <td>MODELE:</td>
      <td>{{ curscan.mod }}</td>
    </tr>
    <tr>
      <td>N°SERIE:</td>
      <td>{{ curscan.ser }}</td>
    </tr>
    <tr>
      <td>MAC:</td>
      <td>{{ curscan.mac }}</td>
    </tr>
    <tr>
      <td>N°INV:</td>
      <td>{{ curscan.inv }}</td>
    </tr>
  </tbody>
</v-simple-table>

</v-card-text>
</v-card>
</template>

<script>
export default {
  computed: {
    curscan () {
      return this.$store.state.currentScan
    },
    filled() {
      return this.$store.state.currentScanFilled
    }
}
}
</script>
