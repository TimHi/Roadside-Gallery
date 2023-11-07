<script setup lang="ts">
import axios from 'axios';
import RoadsideObject from '../components/RoadsideObject.vue';
import type { RoadsideObjectModel } from '../model/RoadsideObjectModel'
import { ref } from 'vue';

const randomObjectUrl = 'http://0.0.0.0:9007/random-object'; // Replace with your API endpoint URL
const roadsideObject = ref<RoadsideObjectModel | undefined>(undefined)
fetchRoadsideObject().then((r) => roadsideObject.value = r)

function getNewObject() {
  fetchRoadsideObject().then((r) => roadsideObject.value = r)
}

async function fetchRoadsideObject() {
  try {
    // Make the GET request to the API endpoint
    const response = await axios.get(randomObjectUrl);

    // Check if the response status is 200 (OK)
    if (response.status === 200) {
      // Parse the JSON response into a RoadsideObject
      const roadsideObject: RoadsideObjectModel = response.data;
      return roadsideObject;
    }
  } catch (error) {
    console.error('Error fetching data:', error);
  }

  return undefined;
}
</script>

<template>
  <div class="header">
    <a href="/about">About</a>
  </div>
  <div class="container">
    <RoadsideObject :model="roadsideObject" />
    <button @click="getNewObject" class="button-31" role="button">Next</button>

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