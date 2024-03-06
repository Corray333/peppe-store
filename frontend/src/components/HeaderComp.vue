<template>
  <div class="header-wrapper">
    <div class="header">
      <img src="../assets/images/logo.png" alt="" class="logo">
      <NavBar/>
      <SearchCart/>
    </div>
  </div>
</template>

<script setup>
import NavBar from './NavBar.vue';
import SearchCart from './SearchCart.vue';
</script>

<style scoped>

.header-wrapper{
  position: fixed;
  background: linear-gradient(rgba(0,0,0,0.5) 0%, rgba(0, 0, 0, 0) 100%);
  z-index: 10;
  width: 100%;
}
.header{
  padding: 10px 0;
  width: 100%;
  max-width: 1040px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  align-items: center;
  justify-content: space-between;
}

</style>
