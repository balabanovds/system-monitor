<template>
  <div>
    <div v-for="(m, i) in data" :key="i">
      {{ m }}
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { getStream, GrpcMetricClient } from "../traits";
import { HOST, PORT } from "../main";

export default defineComponent({
  name: "Chart",
  setup() {
    // const { data } = getStream(HOST, PORT);
    const client = GrpcMetricClient.getInstance();
    const { data } = client.getters();

    return {
      data,
    };
  },
});
</script>

<style></style>
