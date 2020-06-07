<template>
<v-container>
  <v-row>
    <v-col xs=12 sm=6>
      <v-card outlined >
      <v-card-title>Saisie</v-card-title>
      <v-card-text>
      <v-text-field
        ref="scanInput"
        hide-details
        dense
        v-model="cmdtext"
        required
        outlined
        prepend-inner-icon="mdi-barcode"
        @keyup.enter="sendScan" /><br/>
        <v-combobox
        :items="models"
        v-model="curMod"
        outlined
        label="Sélection du Modèle"
        @keyup.enter="submitScan"
        ></v-combobox>
        </v-card-text>
        </v-card>
    </v-col>

    <v-col>
      <v-card outlined >
        <v-card-title>Fiche produit</v-card-title>
        <v-card-text>
          <v-simple-table dense>
            <tbody>
        <tr>
          <td>MODELE:</td>
          <td>{{ curMod }}</td>
        </tr>
        <tr>
          <td>N°SERIE:</td>
          <td>{{ curSer }}</td>
        </tr>
        <tr>
          <td>MAC:</td>
          <td>{{ curMac }}</td>
        </tr>
        <tr>
          <td>N°INV:</td>
          <td>{{ curInv }}</td>
        </tr>
      </tbody>
    </v-simple-table>

    </v-card-text>
    </v-card>
    </v-col>
  </v-row>
  <v-col>
    <v-card outlined >
      <v-card-title>Inventaire</v-card-title>
      <v-card-text>
    <v-data-table
  :headers="headers"
  :items="scans"
  :items-per-page="20"
  class="elevation-1"
 ></v-data-table>
</v-card-text>
</v-card>
  </v-col>
  <v-row>
  </v-row>
</v-container>
</template>

<script>


export default {
  data: () => ({
    valid: true,
    cmdtext: '',
    curInv:'',
    curSer:'',
    curMac:'',
    curMod:'',
    headers: [
          {
            text: 'Modèle',
            align: 'start',
            sortable: true,
            value: 'mod',
          },
          { text: 'N° Série', value: 'ser' },
          { text: 'Adresse MAC', value: 'mac' },
          { text: 'N° Inventaire', value: 'inv' },

        ]
  }),
  mounted() {
    this.focusScanInput()
  },
  methods: {
    focusScanInput(){
      this.$refs.scanInput.focus();
    },
    sendScan() {
      //this.$socket.sendObj({ ct: this.cmdtext, tp: 'CMD' })


      if(this.cmdtext.length == 6 && !isNaN(parseInt(this.cmdtext))){
        this.curInv=this.cmdtext
        let test=false
        this.$store.state.Scans.forEach((scan) => {
           if(scan.inv==this.curInv){
             this.curMod=scan.mod
             this.curSer=scan.ser
             this.curMac=scan.mac

             test=true
             return
           }
        })
        if(test){
          this.cmdtext = ''
          this.focusScanInput()
          return
        }
      }else if (this.cmdtext.length == 12){
        this.curMac=this.cmdtext
      }else{
        this.curSer=this.cmdtext
      }
      this.submitScan()
      this.cmdtext = ''
      this.focusScanInput()
    },
    submitScan(){
      this.$store.commit('addModel', this.curMod)
      if(this.curInv!="" && this.curSer!="" && this.curMac!="" && this.curMod!=""){
        this.$store.commit('addScan', {cmd:"ADD",inv:this.curInv,ser:this.curSer,mac:this.curMac,mod:this.curMod})
        this.curInv=''
        this.curSer=''
        this.curMac=''
      }
    }
  },
  computed: {
    scans () {
      return this.$store.state.Scans
    },
    models () {
      return this.$store.state.Models
    }
  }
}
//https://materialdesignicons.com/
</script>
