<script setup lang="ts">
import type { RoadsideObjectModel } from '../model/RoadsideObjectModel';

interface Props {
    model: RoadsideObjectModel | undefined

}

const props = defineProps<Props>()
</script>

<template>
    <div v-if="props.model !== undefined">

        <h1>{{ props.model.title }} - {{ props.model.year }}</h1>
        <img v-bind:src="'data:image/jpeg;base64,' + props.model.image_data" class="image-container" />

    </div>
    <div v-else>
        <h1 class="error">Error fetching data.</h1>
    </div>
</template>

<style>
.image-container {
    margin-top: 12px;
    max-height: 70vh;
    width: auto;
    overflow: hidden;
    box-shadow: rgba(0, 0, 0, 0.35) 0px 8px 15px;
}

.image-container img {
    width: 100%;
    height: auto;
}
</style>