<script setup lang="ts">
import { useRouter } from 'vue-router';
import { getRandomInt } from '../util/random'
const router = useRouter();

function getNewObject() {
  const id = getRandomInt(11711);
  router.push('/o/'+id);
}

</script>

<template>
  <div class="header">
    <a href="/about">
      <h4>About</h4>
    </a>
  </div>
  <div class="container">
    <div>
      <h1>Explore over 10.000 obscure roadside attractions</h1>
      <img src="/statue.jpg" class="image-container" />
            <p><i>My fav, the man himself.</i></p>
      <button @click="getNewObject" class="button-31" role="button">Start</button>
    </div>
  </div>
</template>


<style>
.header {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  margin: 6px;
}

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}


/* CSS */
.button-31 {
  margin-top: 12px;
  background-color: #222;
  border-radius: 4px;
  border-style: none;
  box-sizing: border-box;
  color: #fff;
  cursor: pointer;
  box-shadow: rgba(0, 0, 0, 0.35) 0px 8px 15px;
  display: inline-block;
  font-family: "Farfetch Basis", "Helvetica Neue", Arial, sans-serif;
  font-size: 16px;
  font-weight: 700;
  line-height: 1.5;
  max-width: none;
  min-height: 44px;
  min-width: 10px;
  outline: none;
  overflow: hidden;
  padding: 9px 20px 8px;
  position: relative;
  text-align: center;
  text-transform: none;
  user-select: none;
  -webkit-user-select: none;
  touch-action: manipulation;

}

.button-31:hover,
.button-31:focus {
  opacity: .75;
}
</style>