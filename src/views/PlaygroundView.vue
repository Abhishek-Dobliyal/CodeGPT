<template>
  <div class="grid sm:grid-cols-2 grid-cols-1 h-screen">
    <div class="container relative">
      <div
        class="flex justify-start mt-4 ml-4 text-xl align-center space-x-2 items-center"
      >
        <i class="fa-solid fa-code fa-bounce"></i>
        <span class="font-semibold">CodeGPT</span>
      </div>
      <CodeArea></CodeArea>
    </div>
    <div class="text-center text-white bg-black">
      <button
        class="bg-black border hover:bg-white hover:text-black text-white font-semibold px-4 py-3 rounded-lg mb-5 mt-3"
        @click="this.$router.push('/playground')"
      >
        Get Started
      </button>
    </div>
  </div>
</template>

<script setup>
import CodeArea from "@/components/CodeArea.vue";
</script>

<style></style>
