<script setup lang="ts">

</script>

<template>
    <div class="header">
        <a href="/">
            <h4>Home</h4>
        </a>
    </div>
    <div class="about-container">
        <div>
            <h1>US Roadside Attractions Photographs</h1>
            <p>Read more about this on the <a href="https://guides.loc.gov/roadside-america-photographs">Library of
                    Congress</a>
                website.</p>
            <p>Images are done by <a href="https://en.wikipedia.org/wiki/John_Margolies">John Margolies</a>.</p>
        </div>
        <div>
            <img src="/statue.jpg" class="image-container" />
            <p><i>My fav, the man himself.</i></p>
            <p>Code can be found at <a href="https://github.com/TimHi/Roadside-Gallery">GitHub</a>.</p>
        </div>
    </div>
</template>

<style>
.about-container {
    display: flex;
    flex-direction: column;
    align-items: center;
}
</style>