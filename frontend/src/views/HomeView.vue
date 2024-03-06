<template>
  <div class="top-wrapper wrapper">
    <section class="top">
      <div class="main">
        <img src="../assets/images/MERCH STORE.png" alt="">
        <img src="../assets/images/logo-long.png" alt="">
        <router-link to="/shop">Shop now</router-link>
      </div>
      <img class="hero" src="../assets/images/pond.png" alt="">
    </section>
  </div>
  <div class="wrapper">
    <FeaturedProducts/>
  </div>
</template>

<script setup>
import FeaturedProducts from '../components/FeaturedProducts.vue'

</script>

<style scoped>

.top{
  height: 100vh;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
}
.wrapper{
  overflow-x: hidden;
}
.main{
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 25px;
}
.hero{
  position: absolute;
  bottom: 0;
  right: -300px;
}
.main>img{
  animation: shake2 10s infinite ease-in;
}
.main>img:nth-last-of-type(1){
  width: 100%;
  margin: 0 auto;
  display: block;
  animation-delay: 2s;
}
a{
  padding: 15px 25px;
  background-color: var(--black-color);
  border-radius: 50px;
  border: 5px solid white;
  text-shadow: 0 0 10px white;
  color: white;
  transition: all 0.3s;
  rotate: -5deg;
  animation: shake 10s infinite ease-in;
}
a:hover{
  background-color: white;
  color: var(--black-color);
  border: 5px solid var(--black-color);
}
@keyframes shake {
  0% { 
    transform: rotate(-3 deg) translate(0, 0); 
  }
  50% { 
    transform: rotate(0deg) translate(5px, 7px); 
  }
  100% { 
    transform: rotate(-3 deg) translate(0, 0); 
  }
}
@keyframes shake2 {
  0% { 
    transform: rotate(0deg) translate(0, 0); 
  }
  50% { 
    transform: rotate(-2deg) translate(-15px, -17px); 
  }
  100% { 
    transform: rotate(0deg) translate(0, 0); 
  }
}

</style>