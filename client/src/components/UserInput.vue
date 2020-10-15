<template>
  <div>
    <input type="text" v-model="n" />
    <input type="text" v-model="m" />
    <button @click="onConnect">run</button>
    <button @click="onCancel">stop</button>
  </div>
</template>

<script lang="ts">
import { HOST, PORT } from "../main";
import { getStream, GrpcMetricClient } from "../traits";
import { defineComponent } from "vue";

export default defineComponent({
  name: "UserInput",
  setup() {
    const client = GrpcMetricClient.getInstance();
    // const { n, m, start, stop } = getStream(HOST, PORT);
    const { n, m } = client.getters();
    function onConnect() {
      client.startStream();
    }

    function onCancel() {
      client.stopStream();
    }

    return {
      n,
      m,
      onConnect,
      onCancel,
    };
  },
});
</script>

<style></style>
