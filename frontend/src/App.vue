<template>
  <HeaderComp />

  <RouterView />
</template>

<script setup>
import { RouterView } from 'vue-router'
import { useStore } from 'vuex'
import HeaderComp from './components/HeaderComp.vue'

const store = useStore()
store.commit('loadCart')

</script>


<style scoped></style>
