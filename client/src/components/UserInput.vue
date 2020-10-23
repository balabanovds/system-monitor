<template>
  <div class="input-container">
    <div>
      <label for="upd-input">Update seconds</label>
      <input
          id="upd-input"
          type="text"
          v-model="n"
          :disabled="disabled"
          placeholder="tick period"
      />
    </div>
    <div>
      <label for="meas-input">Measurement period seconds</label>
      <input
          id="meas-input"
          type="text"
          v-model="m"
          :disabled="disabled"
          placeholder="measurement period"
      />
      <button @click="onConnect" :disabled="disabled" v-show="!active">run</button>
      <button @click="onCancel" :disabled="disabled" v-show="active">stop</button>
    </div>
  </div>

</template>

<script lang="ts">
import {computed, defineComponent} from 'vue';

import {GrpcMetricClient} from '@/traits';

export default defineComponent({
  name: 'UserInput',
  setup() {
    const client = GrpcMetricClient.getInstance();
    const {n, m, active} = client.streamGetters();
    const {error} = client.infoGetter()
    const disabled = computed(() => active.value && error.value !== '')

    function onConnect() {
      client.startStream();
    }

    function onCancel() {
      client.stopStream();
    }

    return {
      n,
      m,
      disabled,
      active,
      onConnect,
      onCancel,
    };
  },
});
</script>

<style scoped>
.input-container {
  margin: 0 auto;
  width: 300px;
  text-align: left;
  padding: 30px;
}

label, input {
  display: block;
  margin-bottom: 10px;
}

input {
  padding: 10px;
  width: 100%
}

button {
  width: 100%;
  padding: 10px;
}

</style>
