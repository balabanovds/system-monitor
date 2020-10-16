<template>
  <div>
    <div v-for="(m, i) in data" :key="i">
      {{ m }}
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { GrpcMetricClient } from "../traits";

export default defineComponent({
  name: "Chart",
  props: {
    name: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    console.log(props);

    const client = GrpcMetricClient.getInstance();
    const { data } = client.getters();

    return {
      data,
    };
  },
});
</script>

<style></style>
