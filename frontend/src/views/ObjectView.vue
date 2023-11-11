<script setup lang="ts">
import axios from 'axios';
import RoadsideObject from '../components/RoadsideObject.vue';
import type { RoadsideObjectModel } from '../model/RoadsideObjectModel'
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { getRandomInt } from '../util/random'
import router from '../router';

const randomObjectUrl = 'http://localhost:8090/object/'; //TODO_THL: .env
const isLoading = ref<boolean>(true);
const roadsideObject = ref<RoadsideObjectModel | undefined>(undefined);
const route = useRoute();
const id = route.params.id as string
fetchRoadsideObject(id).then((r) => roadsideObject.value = r);

async function fetchRoadsideObject(id: string) {
  isLoading.value = true;
  try {
    const response = await axios.get(randomObjectUrl + id);
    if (response.status === 200) {
      const roadsideObject: RoadsideObjectModel = response.data;
      isLoading.value = false;
      return roadsideObject;
    }
  } catch (error) {
    roadsideObject.value = undefined;
    console.error('Error fetching data:', error);
  }
  isLoading.value = false;
  return undefined;
}
</script>

<template>
  <div class="header">
    <a href="/about">
      <h4>About</h4>
    </a>
  </div>
  <div class="container">
    <div v-if="isLoading">
      <h1>Loading... 👀</h1>
    </div>
    <div v-else>
      <RoadsideObject :model="roadsideObject" />
      <button @click="() => {
        const id = getRandomInt(11711);
        $router.push('/o/'+id);
        fetchRoadsideObject(id.toString()).then((r) => {
          if(roadsideObject !== undefined && r !== undefined){
            roadsideObject = r;
          }else{
            router.replace('/');
          }
        });
      }" class="button-31" role="button">Next</button>
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