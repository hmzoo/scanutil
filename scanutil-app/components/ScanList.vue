<template>
<v-card outlined >
  <v-card-title>Inventaire</v-card-title>
  <v-card-text>
<v-data-table
:headers="headers"
:items="scans"
:items-per-page="20"
class="elevation-1"
dense
>

<template  v-slot:item.inv="{ item }">
  <b>{{ item.inv }}</b>
</template>

</v-data-table>
</v-card-text>
</v-card>
</template>


<script>
export default {
  data: () => ({
    headers: [
      { text: 'N° Inventaire', value: 'inv' },
          {
            text: 'Modèle',
            align: 'start',
            sortable: true,
            value: 'mod',
          },
          { text: 'N° Série', value: 'ser' },
          { text: 'Adresse MAC', value: 'mac' },


        ]
  }),
  computed: {
    scans () {
      return this.$store.state.Scans
    }
  }
}
//https://materialdesignicons.com/
</script>
