<template>
  <img src="./2_low0.5.svg" class="BackGroundImg">
  <body class="main-body">
  <h1>Note-Share-Platform</h1>
  <HelloWorld />
</body>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    HelloWorld
  },
}

</script>

<style>
.main-body{
  width : 100%;
  height: 100%;
  position: relative;
  /* background-color: red; */
}
.BackGroundImg{
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  object-fit: cover; /* 图片缩放以覆盖整个容器 */
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
