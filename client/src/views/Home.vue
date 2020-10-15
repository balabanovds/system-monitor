<template>
  <div class="home">
    <user-input />
    <chart />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import UserInput from "../components/UserInput.vue";
import Chart from "../components/Chart.vue";

export default defineComponent({
  name: "Home",

  components: {
    UserInput,
    Chart,
  },
});
</script>
