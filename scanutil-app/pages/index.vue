<template>
<v-container>
  <v-row>
    <v-col xs=12 sm=6>
      <ScanInput /><br/>
    </v-col>
    <v-col>
      <ScanCurrent />
    </v-col>
  </v-row>
  <v-row>
  <v-col>
    <ScanList />
  </v-col>
  </v-row>
</v-container>
</template>

<script>
import ScanCurrent from '~/components/ScanCurrent.vue'
import ScanInput from '~/components/ScanInput.vue'
import ScanList from '~/components/ScanList.vue'



export default {
  components: {
    ScanCurrent,
    ScanInput,
    ScanList
  },
  data: () => ({

  }),
  mounted() {

  },
  methods: {

  },
  computed: {

  }
}
//https://materialdesignicons.com/
</script>
