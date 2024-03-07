<template>
  <div class="wrapper">
    <section class="shop">
      <SideBar/>
      <div class="product-list">
        <ProductCard v-for="(product, i) in products" :key="i" :product="product"/>
      </div>
    </section>
  </div>
</template>
<script setup>
import SideBar from '@/components/SideBar.vue'
import ProductCard from '@/components/ProductCard.vue';

import {ref, onBeforeMount} from 'vue'
import axios from 'axios'

const products = ref()


onBeforeMount(async () => {
  const response = await axios.get('http://localhost:3001/products')
  products.value = response.data
})


</script>

<style scoped>

.wrapper{
  padding-top: 100px;
}
.shop{
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  gap: 1em;
}
.product-list{
  grid-column: 2/5;
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  gap: 1em;
}

</style>