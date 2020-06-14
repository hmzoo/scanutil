<template>
  <v-card outlined >
  <v-card-title>Saisie</v-card-title>
  <v-card-text>
    <v-text-field
      ref="scanInput"
      hide-details
      dense
      v-model="scantext"
      required
      outlined
      prepend-inner-icon="mdi-barcode"
      @keyup.enter="submitScan" />
      <br/>
      <v-combobox
      :items="models"
      v-model="modeltext"
      outlined
      label="Sélection du Modèle"
      @keyup.enter="submitModel"
      @change="submitModel"
      ></v-combobox>
    </v-card-text>
    </v-card>

</template>

<script>
export default {
data: () => ({
  scantext: '',
  modeltext: '',
}),
mounted() {
  this.focus()
},
methods: {
  focus(){
    this.$refs.scanInput.focus();
  },
  submitScan(){
    if(this.scantext != ""){
      this.$store.commit('updateCurrentScan', this.scantext)
      this.scantext=''
      this.updateScanList()
    }
  },
  submitModel(){
    if(this.modeltext != ""){
      this.$store.commit('updateCurrentModel', this.modeltext)
      this.updateScanList()
    }

  },
  updateScanList() {
    this.$store.commit('checkCurrentScan')
    if(this.$store.state.currentScanFilled){
      let scan=this.$store.state.currentScan
      this.$socket.sendObj({
          tp:"ADD",
          item:{mod:scan.mod,inv:scan.inv,ser:scan.ser,mac:scan.mac}
        }
      )
      //this.$store.commit('updateScanList')
      this.$store.commit('clearCurrentScan')
      this.$store.commit('checkCurrentScan')
    }
    this.focus()
  }
},
computed: {
  models () {
    return this.$store.state.Models
  }
}
}
</script>
