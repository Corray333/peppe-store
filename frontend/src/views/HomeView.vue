<template>
  <div class="top-wrapper wrapper">
    <section class="top">
      <div class="main">
        <img src="../assets/images/MERCH STORE.png" alt="">
        <img src="../assets/images/logo-long.png" alt="">
        <router-link to="/products">Shop now</router-link>
      </div>
      <img class="hero" src="../assets/images/pond.png" alt="">
    </section>
  </div>
  <div class="wrapper">
    <FeaturedProducts/>
  </div>
  <div class="wrapper">
    <section class="about">
      <div class="about_picture">
        <img src="../assets/images/logo.png" alt="">
      </div>
      <div class="about_info_wrapper">
        <div class="about_info">
          <h2>About <green>CopeNothing</green></h2>
          <p>Are you ready for COPE? Our team has worked hard to create a unique merch for pepe & pndc holders. All our products are manufactured using the most modern technologies and high-quality materials. They are not only stylish and unique, but also comfortable to use. We strive to provide our customers with a high level of satisfaction and joy from shopping.</p>
          <button>Twitter</button>
        </div>
      </div>
    </section>
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
  text-transform: uppercase;
}
a:hover{
  background-color: white;
  color: var(--black-color);
  border: 5px solid var(--black-color);
}
@keyframes shake {
  0% { 
    transform: rotate(-3 deg) translate(-7, -5); 
  }
  50% { 
    transform: rotate(0deg) translate(5px, 7px); 
  }
  100% { 
    transform: rotate(-3 deg) translate(-7, -5); 
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

.about{
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  align-items: center;
  gap: 50px;
  padding: 25px;
  background-color: white;
  border-radius: 25px;
}
.about>img{
  width: 100%;
  border-radius: 25px;
}
.about_info_wrapper{
  grid-column: 2/4;
  display: flex;
  justify-content: center;
}
.about_picture{
  width: 100%;
  height: 100%;
  border-radius: 25px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-image: url("../assets/images/about_cope.png");
  background-size: cover;
  background-position: center;
}
.about_picture img{
  height: 7em;
  animation: shake 10s infinite ease-in;
}
.about_info{
  display: flex;
  flex-direction: column;
  gap: 15px;
  padding: 50px 0;
  width: min-content;
}
.about_info p{
  width: 100%;
}
.about_info h2{
  text-wrap: nowrap;
  font-size: 2.5em;
  font-weight: normal;
}
green{
  color: var(--accent-color);
}
.about_info button{
  width: fit-content;
}

</style>